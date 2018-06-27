package parser

import (
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/exec"
)

func (p *Parser) parse_logic() (*context.Node,error) {

  expr, err := p.parse_additive()
  if err != nil {
    return nil, err
  }

  token := p.lexer.Peek()
  for token.Text() == "==" || token.Text() == "!=" {
    token = p.lexer.Next()

    right, err := p.parse_additive()
    if err != nil {
      return nil, err
    }

    switch token.Text() {
      case "==":
        expr = context.NewNode( token, exec.EqualHandler, expr, right )
      case "!=":
        expr = context.NewNode( token, exec.NotEqualHandler, expr, right )
    }

    token = p.lexer.Peek()
  }

  return expr, err
}
