package parser

import (
  "errors"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/exec"
)

func (p *Parser) parse_additive() (*context.Node,error) {

  expr, err := p.parse_multiplicative()
  if err != nil {
    return nil, err
  }

  token := p.lexer.Peek()
  for token.Text() == "+" || token.Text() == "-" {
    token = p.lexer.Next()

    right, err := p.parse_multiplicative()
    if err != nil {
      return nil, err
    }

    switch token.Text() {
      case "+":
        expr, err = OptimizeBinaryFunction( token, expr, right, exec.Add )
        if err != nil {
          return nil, err
        }
      case "-":
        expr, err = OptimizeBinaryFunction( token, expr, right, exec.Sub )
        if err != nil {
          return nil, err
        }
    }

    token = p.lexer.Peek()
  }

  return expr, err
}

func (p *Parser) parse_multiplicative() (*context.Node,error) {

  expr, err := p.parse_negative()
  if err != nil {
    return nil, err
  }

  token := p.lexer.Peek()
  for token.Text() == "*" || token.Text() == "/" {
    token = p.lexer.Next()

    right, err := p.parse_negative()
    if err != nil {
      return nil, err
    }

    switch token.Text() {
      case "*":
        expr, err = OptimizeBinaryFunction( token, expr, right, exec.Mult )
        if err != nil {
          return nil, err
        }

      case "/":
        expr, err = OptimizeBinaryFunction( token, expr, right, exec.Div )
    }

    token = p.lexer.Peek()
  }

  return expr, err
}

// Handles the special case of -value. If value is a *Value then we negate it here
func (p *Parser) parse_negative() (*context.Node,error) {
  token := p.lexer.Peek()

  if token.Text() == "-" {
    token = p.lexer.Next()

    // Future: If we want "--" operator put test here

    expr, err := p.ParseExpression()
    if err != nil {
      return nil, err
    }

    // If expr is a value then we can just negate it now
    if expr.Value() != nil {
      value := expr.Value()
      switch value.Type() {
        case context.VAR_BOOL:
          value = context.BoolValue( !value.Bool() )
        case context.VAR_INT:
          value = context.IntValue( -value.Int() )
        case context.VAR_FLOAT:
          value = context.FloatValue( -value.Float() )
        default:
          return nil, errors.New( "Unsupported type for neg" )
      }
      return context.NewConstant( token, value), nil
    }

    return context.NewNode( token, exec.NegHandler, expr, nil ), nil
  }

  // Special case a + here means positive so skip the token as it's a nop
  if token.Text() == "+" {
    token = p.lexer.Next()
    // Future: If we want "++" operator put test here
  }

  expr, err := p.parse_parens()
  return expr, err
}
