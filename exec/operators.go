package exec
/*
// An operation that takes no arguments
func ActionOp( p *Parser, n *Node, h NodeHandler ) (*Node,error) {
  n1 := p.New( h )
  return n1, n.Append( n1 )
}

// Parse a unary operation, e.g. NOT v
func UnaryOp( p *Parser, n *Node, h NodeHandler ) (*Node,error) {
  n1 := p.New( h )

  err := n.Append( n1 )
  if err != nil {
    return nil, err
  }

  a, err := p.ParseToken( n1 )
  return a, err
}

// Parse a binary operation, e.g. a AND b
func BinaryOp( p *Parser, n *Node, h NodeHandler ) (*Node,error) {
  n1 := p.New( h )
  n.Replace( n1 )

  a, err := p.ParseToken( n1 )
  return a, err
}
*/
