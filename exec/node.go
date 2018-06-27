package exec

type NodeHandler func( *Context, *Node ) error

type ParserDefinition func( *Parser, *Node, NodeHandler ) (*Node,error)

type FuncMapEntry struct {
  NodeHandler
  ParserDefinition
}

type FuncMap map[string]NodeHandler

// A node in the filter tree
type Node struct {
  tokenRune   rune
  token       string
  // left hand side
  left       *Node
  // right hand side
  right      *Node
  // handler for this node
  handler     NodeHandler
  // The value of this node
  value      *Value
  // slice of nodes for when we have multiple args or statements
  list     []*Node
}

func NewNode( t string, f NodeHandler ) *Node {
  return &Node{ token: t, handler: f }
}

// Invoke the handler of this node
func (n *Node) Invoke( m *Context ) error {
  if n.handler != nil {
    return n.handler( m, n )
  } else if n.value != nil {
    m.Push( n.value )
  }
  return nil
}

// Invokes the left hand side node or returns false if none
func (n *Node) InvokeLhs( m *Context ) error {
  if n.left != nil {
    return n.left.Invoke(m)
  }
  return nil
}

// Invokes the right hand side node or returns false if none
func (n *Node) InvokeRhs( m *Context ) error {
  if n.right != nil {
    return n.right.Invoke(m)
  }
  return nil
}

// Invoke2 is for use by handlers. It will invoke both left & right in one go
func (n *Node) Invoke2( m *Context ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  return n.InvokeRhs(m)
}

// Token returns this nodes token text
func (n *Node) Token() string {
  return n.token
}

// Rune returns this nodes rune (optional)
func (n *Node) Rune() rune {
  return n.tokenRune
}

// Value returns this nodes value or nil
func (n *Node) Value() *Value {
  return n.value
}

// Left returns the left hand node or nil
func (n *Node) Left() *Node {
  return n.left
}

// Right returns the right hand node or nil
func (n *Node) Right() *Node {
  return n.left
}

// ForEach invokes a function for each node in this nodes list
func (n *Node) ForEach( f func(*Node) error ) error {
  for _, n1 := range n.list {
    err := f(n1)
    if err != nil {
      return err
    }
  }
  return nil
}

// ForEachAll invokes a function for the left & right hand nodes (if present)
// and for any node within this nodes list
func (n *Node) ForEachAll( f func(*Node) error ) error {
  var err error
  if n.left != nil {
    err = f(n.left)
  }
  if err == nil && n.right != nil {
    err = f(n.right)
  }
  if err == nil {
    err = n.ForEach( f )
  }
  return err
}
