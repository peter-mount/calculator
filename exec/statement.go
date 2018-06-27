package exec

import (
  "errors"
  "fmt"
  "github.com/peter-mount/calculator/lex"
  "text/scanner"
)

func (p *Parser) parse_statements() (*Node,error) {
  block := &Node{token: "{ }", handler: invokeAllHandler}
  token := p.lexer.Peek()
  for token.Token() != scanner.EOF {
    expr, err := p.parse_statement_block()
    if err != nil {
      return nil, err
    }
    if expr != nil {
      block.list = append( block.list, expr )
    }
    token = p.lexer.Peek()
  }
  return block, nil
}

func (p *Parser) parse_statement_block() (*Node,error) {

  token := p.lexer.Peek()
  if token.Text() == "{" {
    p.lexer.Next()

    block := &Node{token: "{ }", handler: invokeScopeHandler}
    token := p.lexer.Peek()
    for token.Token() != scanner.EOF && token.Text() != "}" {
      expr, err := p.parse_statement()
      if err != nil {
        return nil, err
      }
      if expr != nil {
        block.list = append( block.list, expr )
      }
      token = p.lexer.Peek()
    }
    if token.Text() != "}" {
      return nil, fmt.Errorf( "Expecting }" )
    }
    return block, nil
  }

  expr, err := p.parse_statement()
  return expr, err
}

func (p *Parser) parse_statement() (*Node,error) {

  // Skip ; as optional terminators
  token := p.lexer.Peek()
  if token.Text() == "}" {
    token = p.lexer.Next()
    return nil, nil
  }
  for token.Text() == ";" {
    token = p.lexer.Skip()
  }

  expr, err := p.parse_arithmetic()
  if err != nil {
    return nil, err
  }

  token = p.lexer.Peek()
  for token.Text() == "=" && expr != nil && expr.tokenRune == lex.TOKEN_VARIABLE {
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

  m.Push( val )
  return nil
}

// invokeAllHandler Invokes all nodes within the supplied list
func invokeAllHandler( m *Context, n *Node ) error {
  for _, s := range n.list {
    err := s.Invoke(m)
    if err != nil {
      return err
    }
  }
  return nil
}

// invokeScopeHandler calls invokeAllHandler with a variable scope that lasts
// for the duration of the call
func invokeScopeHandler( m *Context, n *Node ) error {
  m.StartScope()
  defer m.EndScope()
  return invokeAllHandler( m, n )
}
