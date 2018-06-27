package parser

import (
  "errors"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/exec"
  "github.com/peter-mount/calculator/lex"
)

// OptimizeOperation will if both left and right nodes are constants return
// a constant node with the result of some function.
// If either are not constant then a new node will be created with the supplied handler
// attached.
func OptimizeOperation( token *lex.Token, left *context.Node, right *context.Node, f func(*context.Value,*context.Value)(*context.Value,error) ) (*context.Node,error) {
  if left.IsConstant() && right.IsConstant() {
    c, err := f( left.Value(), right.Value() )
    if err != nil {
      return nil, err
    }
    return context.NewConstant( token, c ), nil
  } else {
    return context.NewNode(
      token,
      func( m *context.Context, n *context.Node ) error {
        err := n.Invoke2(m)
        if err != nil {
          return err
        }

        a, b, err := m.Pop2()
        if err != nil {
          return err
        }

        c, err := f( a, b )
        if err == nil {
          m.Push( c )
        }
        return err
      },
      left,
      right ), nil
  }
}

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
        expr, err = OptimizeOperation( token, expr, right, exec.Add )
        if err != nil {
          return nil, err
        }
      case "-":
        expr, err = OptimizeOperation( token, expr, right, exec.Sub )
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
        expr, err = OptimizeOperation( token, expr, right, exec.Mult )
        if err != nil {
          return nil, err
        }

      case "/":
        expr, err = OptimizeOperation( token, expr, right, exec.Div )
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

    expr, err := p.parse_arithmetic()
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
