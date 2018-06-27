package exec

import (
  "errors"
  "fmt"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/lex"
  "text/scanner"
)

func (p *Parser) parse_statements() (*context.Node,error) {
  block := context.NewBlock( invokeAllHandler )
  token := p.lexer.Peek()
  for token.Token() != scanner.EOF {
    expr, err := p.parse_statement_block()
    if err != nil {
      return nil, err
    }
    if expr != nil {
      block.Append( expr )
    }
    token = p.lexer.Peek()
  }
  return block, nil
}

func (p *Parser) parse_statement_block() (*context.Node,error) {

  token := p.lexer.Peek()
  if token.Text() == "{" {
    p.lexer.Next()

    block := context.NewBlock( invokeScopeHandler )
    token := p.lexer.Peek()
    for token.Token() != scanner.EOF && token.Text() != "}" {
      expr, err := p.parse_statement()
      if err != nil {
        return nil, err
      }
      if expr != nil {
        block.Append( expr )
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

func (p *Parser) parse_statement() (*context.Node,error) {

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
  for token.Text() == "=" && expr != nil && expr.Token().Token() == lex.TOKEN_VARIABLE {
    token = p.lexer.Next()

    left, err := p.parse_arithmetic()
    if err != nil {
      return nil, err
    }

    expr = context.NewNode( expr.Token(), setVarHandler, left, nil )

    token = p.lexer.Peek()
  }

  return expr, err
}

func setVarHandler( m *context.Context, n *context.Node ) error {
  err := n.InvokeLhs( m )
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  m.SetVar( n.Token().Text(), a )

  return nil
}

func getVarHandler( m *context.Context, n *context.Node ) error {

  val := m.GetVar( n.Token().Text() )
  if val == nil {
    return errors.New( "Unknown variable " + n.Token().Text() )
  }

  m.Push( val )
  return nil
}

// invokeAllHandler Invokes all nodes within the supplied list
func invokeAllHandler( m *context.Context, n *context.Node ) error {
  return n.ForEach( func(n1 *context.Node) error {
    return n1.Invoke(m)
  } )
}

// invokeScopeHandler calls invokeAllHandler with a variable scope that lasts
// for the duration of the call
func invokeScopeHandler( m *context.Context, n *context.Node ) error {
  m.StartScope()
  defer m.EndScope()
  return invokeAllHandler( m, n )
}
