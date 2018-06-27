package exec

import (
  "errors"
  "github.com/peter-mount/calculator/context"
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
        expr = context.NewNode( token, addHandler, expr, right )
      case "-":
        expr = context.NewNode( token, subHandler, expr, right )
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
        expr = context.NewNode( token, multHandler, expr, right )
      case "/":
        expr = context.NewNode( token, divHandler, expr, right )
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

    return context.NewNode( token, negHandler, expr, nil ), nil
  }

  // Special case a + here means positive so skip the token as it's a nop
  if token.Text() == "+" {
    token = p.lexer.Next()
    // Future: If we want "++" operator put test here
  }

  expr, err := p.parse_parens()
  return expr, err
}

func negHandler( m *context.Context, n *context.Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  switch a.Type() {
    case context.VAR_BOOL:
      m.PushBool( !a.Bool() )
    case context.VAR_INT:
      m.PushInt( -a.Int() )
    case context.VAR_FLOAT:
      m.PushFloat( -a.Float() )
    default:
      return errors.New( "Unsupported type for neg" )
  }

  return nil
}

func addHandler( m *context.Context, n *context.Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.OperationType( b ) {
    case context.VAR_BOOL:
      m.PushInt( a.Int() + b.Int() )
    case context.VAR_INT:
      m.PushInt( a.Int() + b.Int() )
    case context.VAR_FLOAT:
      m.PushFloat( a.Float() + b.Float() )
    case context.VAR_STRING:
      m.PushString( a.String() + b.String() )
    default:
      return errors.New( "Unsupported type for add" )
  }

  return nil
}

func subHandler( m *context.Context, n *context.Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.OperationType( b ) {
    case context.VAR_BOOL:
      m.PushInt( a.Int() - b.Int() )
    case context.VAR_INT:
      m.PushInt( a.Int() - b.Int() )
    case context.VAR_FLOAT:
      m.PushFloat( a.Float() - b.Float() )
    default:
      return errors.New( "Unsupported type for sub" )
  }

  return nil
}

func multHandler( m *context.Context, n *context.Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.OperationType( b ) {
    case context.VAR_BOOL:
      m.PushInt( a.Int() * b.Int() )
    case context.VAR_INT:
      m.PushInt( a.Int() * b.Int() )
    case context.VAR_FLOAT:
      m.PushFloat( a.Float() * b.Float() )
    default:
      return errors.New( "Unsupported type for mult" )
  }

  return nil
}

func divHandler( m *context.Context, n *context.Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  if b.IsZero() {
    return errors.New( "Division by zero")
  }

  switch a.OperationType( b ) {
    case context.VAR_BOOL:
      m.PushInt( a.Int() / b.Int() )
    case context.VAR_INT:
      m.PushInt( a.Int() / b.Int() )
    case context.VAR_FLOAT:
      m.PushFloat( a.Float() / b.Float() )
    default:
      return errors.New( "Unsupported type for div" )
  }

  return nil
}

func intHandler( m *context.Context, n *context.Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  switch a.Type() {
    case context.VAR_BOOL:
      m.PushInt( a.Int() )
    case context.VAR_INT:
      m.PushInt( a.Int() )
    case context.VAR_FLOAT:
      m.PushInt( a.Int() )
    default:
      return errors.New( "Unsupported type for int" )
  }

  return nil
}
