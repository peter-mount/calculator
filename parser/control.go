package parser

import (
  "fmt"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/exec"
)

// if expression { true statement }
// if expression { true statement } else { false statement }
func (p *Parser) parse_if() (*context.Node,error) {
  token := p.lexer.Next()

  // the condition
  left, err := p.ParseExpression()
  if err != nil {
    return nil, err
  }

  // The true block
  right, err := p.parse_statement_block()
  if err != nil {
    return nil, err
  }

  expr := context.NewNode( token, exec.IfHandler, left, right )

  // optional else
  token = p.lexer.Peek()
  if token.Text() == "else" {
    token = p.lexer.Next()

    elseExpr, err := p.parse_statement_block()
    if err != nil {
      return nil, err
    }
    expr.Append( elseExpr )

    token = p.lexer.Peek()
  }

  return expr, nil
}

// do { statement } while expression
// do { statement } until expression
func (p *Parser) parse_do() (*context.Node,error) {
  token := p.lexer.Next()

  // The statements block
  right, err := p.parse_statement_block()
  if err != nil {
    return nil, err
  }

  token = p.lexer.Peek()
  switch token.Text() {
    case "while":
      token = p.lexer.Next()
      left, err := p.ParseExpression()
      if err != nil {
        return nil, err
      }
      return context.NewNode( token, exec.DoWhileHandler, left, right ), nil
    case "until":
      token = p.lexer.Next()
      left, err := p.ParseExpression()
      if err != nil {
        return nil, err
      }
      return context.NewNode( token, exec.DoUntilHandler, left, right ), nil
    default:
      return nil, fmt.Errorf( "Expected while or until, got: \"%s\"", token.Text() )
  }
}

// while( expression ) { statement }
// until( expression ) { statement }
// h is the actual NodeHandler to implement this
func (p *Parser) parse_condLoop( h context.NodeHandler ) (*context.Node,error) {
  token := p.lexer.Next()

  // the condition
  left, err := p.ParseExpression()
  if err != nil {
    return nil, err
  }

  // The statements block
  right, err := p.parse_statement_block()
  if err != nil {
    return nil, err
  }

  expr := context.NewNode( token, h, left, right )

  return expr, nil
}
