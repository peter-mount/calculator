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
  token       string
  // left hand side
  left       *Node
  // right hand side
  right      *Node
  // handler for this node
  handler     NodeHandler
  // The value of this node
  value      *Value
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
