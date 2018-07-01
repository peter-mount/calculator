package parser

import (
  //"errors"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/exec"
)

func (p *Parser) parse_complex() (*context.Node,error) {
  expr, err := p.parse_parens()
  if err != nil {
    return nil, err
  }

  token := p.lexer.Peek()
  if token.Text() == "i" {
    token = p.lexer.Next()
    expr, err = OptimizeUnaryFunction( token, expr, exec.ToImaginary )
    if err != nil {
      return nil, err
    }
  }

  return expr, err
}
