package exec

import (
  "errors"
  "fmt"
)

func (p *Parser) parse_statement() (*Node,error) {

  expr, err := p.parse_arithmetic()
  if err != nil {
    return nil, err
  }

  token := p.lexer.Peek()
  for token.text == "=" && expr != nil && expr.tokenRune == TOKEN_VARIABLE {
    token = p.lexer.Next()

    left, err := p.parse_arithmetic()
    if err != nil {
      return nil, err
    }

    expr = &Node{ token:expr.token, left:left, handler: setVarHandler }

    token = p.lexer.Peek()
  }

  return expr, err
}

func setVarHandler( m *Context, n *Node ) error {
  err := n.InvokeLhs( m )
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  m.SetVar( n.token, a )

  return nil
}

func getVarHandler( m *Context, n *Node ) error {

  val := m.GetVar( n.token )
  if val == nil {
    return errors.New( "Unknown variable " + n.token )
  }

  fmt.Printf( "Get \"%s\" = %v\n", n.token, val )

  m.Push( val )
  return nil
}
