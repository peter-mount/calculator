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
  lhs        *Node
  // right hand side
  rhs        *Node
  // handler for this node
  handler     NodeHandler
  // The value of this node
  value      *Value
  test  bool
}

// set lhs or rhs if lhs is occupied
func (n *Node) Append( next *Node ) error {
  if n.lhs == nil {
    n.lhs = next
  } else if n.rhs == nil {
    n.rhs = next
  } else {
    return errors.New( "Node full" )
  }
  next.parent = n
  return nil
}

// Replace this node in the tree with a new node and make this one the new node's
// lhs. Used when parsing a AND b where this is a.
func (n *Node) Replace( next *Node ) error {
  if n.parent == nil {
    return errors.New( "No lhs for " + next.token )
  }

  p := n.parent
  if p.lhs == n {
    p.lhs = next
  } else {
    p.rhs = next
  }

  next.parent = p

  next.Append( n )

  return nil
}

// Invoke the handler of this node
func (n *Node) Invoke( m *Context ) error {
  if n.handler != nil {
    return n.handler( m, n )
  }
  return nil
}

// Invokes the left hand side node or returns false if none
func (n *Node) InvokeLhs( m *Context ) error {
  if n.lhs != nil {
    return n.lhs.Invoke(m)
  }
  return nil
}

// Invokes the right hand side node or returns false if none
func (n *Node) InvokeRhs( m *Context ) error {
  if n.rhs != nil {
    return n.rhs.Invoke(m)
  }
  return nil
}

// Invoke2 is for use by handlers. It will invoke both lhs & rhs in one go
func (n *Node) Invoke2( m *Context ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  return n.InvokeRhs(m)
}
