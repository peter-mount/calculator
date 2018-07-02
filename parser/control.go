package parser

import (
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

// while( expression ) { statement }
func (p *Parser) parse_while() (*context.Node,error) {
  token := p.lexer.Next()

  // the condition
  left, err := p.parse_parens()
  if err != nil {
    return nil, err
  }

  // The statements block
  right, err := p.parse_statement_block()
  if err != nil {
    return nil, err
  }

  expr := context.NewNode( token, exec.WhileHandler, left, right )

  return expr, nil
}
