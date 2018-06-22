package exec

import (
  "errors"
)

type NodeHandler func( *Context, *Node ) error

type ParserDefinition func( *Parser, *Node, NodeHandler ) (*Node,error)

type FuncMapEntry struct {
  NodeHandler
  ParserDefinition
}

type FuncMap map[string]FuncMapEntry

// A node in the filter tree
type Node struct {
  token       string
  // parent node
  parent     *Node
  // left hand side
  left       *Node
  // right hand side
  right      *Node
  // handler for this node
  handler     NodeHandler
  // The value of this node
  value      *Value
  // precedence
  precedence  int
}

// set left or right if left is occupied
func (n *Node) Append( next *Node ) error {
  if n.left == nil {
    n.left = next
  } else if n.right == nil {
    n.right = next
  } else {
    return errors.New( "Node full" )
  }
  next.parent = n
  return nil
}

func NewNode( t string, f NodeHandler ) *Node {
  return &Node{ token: t, handler: f }
}

func (n *Node) AppendHandler( p *Parser, h NodeHandler ) (*Node,error) {
  n1 := NewNode( p.token, h )
  return n1, n.Append( n1 )
}

func (n *Node) AppendValue( p *Parser, v *Value ) (*Node,error) {
  n1 := &Node{ token: v.String(), value: v, precedence: p.precedence }
  return n1, n.Append( n1 )
}

// Replace this node in the tree with a new node and make this one the new node's
// left. Used when parsing a AND b where this is a.
func (n *Node) Replace( next *Node ) error {
  if n.parent == nil {
    return errors.New( "No left for " + next.token )
  }

  p := n.parent
  if p.left == n {
    p.left = next
  } else {
    p.right = next
  }

  next.parent = p

  next.Append( n )

  return nil
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
