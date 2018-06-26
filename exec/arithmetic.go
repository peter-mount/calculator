package exec

import (
  "errors"
)

func (p *Parser) parse_additive() (*Node,error) {

  expr, err := p.parse_multiplicative()
  if err != nil {
    return nil, err
  }

  token := p.lexer.Peek()
  for token.text == "+" || token.text == "-" {
    token = p.lexer.Next()

    right, err := p.parse_multiplicative()
    if err != nil {
      return nil, err
    }

    switch token.text {
      case "+":
        expr = &Node{ token:token.text, left:expr, right: right, handler: addHandler }
      case "-":
        expr = &Node{ token:token.text, left:expr, right: right, handler: subHandler }
    }

    token = p.lexer.Peek()
  }

  return expr, err
}

func (p *Parser) parse_multiplicative() (*Node,error) {

  expr, err := p.parse_negative()
  if err != nil {
    return nil, err
  }

  token := p.lexer.Peek()
  for token.text == "*" || token.text == "/" {
    token = p.lexer.Next()

    right, err := p.parse_negative()
    if err != nil {
      return nil, err
    }

    switch token.text {
      case "*":
        expr = &Node{ token:token.text, left:expr, right: right, handler: multHandler }
      case "/":
        expr = &Node{ token:token.text, left:expr, right: right, handler: divHandler }
    }

    token = p.lexer.Peek()
  }

  return expr, err
}

// Handles the special case of -value. If value is a *Value then we negate it here
func (p *Parser) parse_negative() (*Node,error) {
  token := p.lexer.Peek()

  if token.text == "-" {
    token = p.lexer.Next()
    expr, err := p.parse_arithmetic()
    if err != nil {
      return nil, err
    }

    // If expr is a value then we can just negate it now
    if expr.value != nil {
      value := expr.value
      switch value.Type() {
        case VAR_BOOL:
          value = BoolValue( !value.Bool() )
        case VAR_INT:
          value = IntValue( -value.Int() )
        case VAR_FLOAT:
          value = FloatValue( -value.Float() )
        default:
          return nil, errors.New( "Unsupported type for neg" )
      }
      return &Node{ token:value.String(), value: value }, nil
    }

    return &Node{ token:token.text, left:expr, handler: negHandler }, nil
  }

  expr, err := p.parse_parens()
  return expr, err
}

func negHandler( m *Context, n *Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_BOOL:
      m.PushBool( !a.Bool() )
    case VAR_INT:
      m.PushInt( -a.Int() )
    case VAR_FLOAT:
      m.PushFloat( -a.Float() )
    default:
      return errors.New( "Unsupported type for neg" )
  }

  return nil
}

// OperationType returns the type of the suggested value when performing some
// operation like addition or multiplication to keep the precision of the result.
// For example, if a Value is an Integer but the passed value is Float then
// we should use float.
func (a *Value) OperationType( b *Value ) int {
  t  := a.Type();
  if a.Type() == VAR_FLOAT || b.Type() == VAR_FLOAT {
    t = VAR_FLOAT
  } else if a.Type() == VAR_INT || b.Type() == VAR_INT {
    t = VAR_INT
  }
  return t
}

func addHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.OperationType( b ) {
    case VAR_BOOL:
      m.PushInt( a.Int() + b.Int() )
    case VAR_INT:
      m.PushInt( a.Int() + b.Int() )
    case VAR_FLOAT:
      m.PushFloat( a.Float() + b.Float() )
    case VAR_STRING:
      m.PushString( a.String() + b.String() )
    default:
      return errors.New( "Unsupported type for add" )
  }

  return nil
}

func subHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.OperationType( b ) {
    case VAR_BOOL:
      m.PushInt( a.Int() - b.Int() )
    case VAR_INT:
      m.PushInt( a.Int() - b.Int() )
    case VAR_FLOAT:
      m.PushFloat( a.Float() - b.Float() )
    default:
      return errors.New( "Unsupported type for sub" )
  }

  return nil
}

func multHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.OperationType( b ) {
    case VAR_BOOL:
      m.PushInt( a.Int() * b.Int() )
    case VAR_INT:
      m.PushInt( a.Int() * b.Int() )
    case VAR_FLOAT:
      m.PushFloat( a.Float() * b.Float() )
    default:
      return errors.New( "Unsupported type for mult" )
  }

  return nil
}

func divHandler( m *Context, n *Node ) error {
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
    case VAR_BOOL:
      m.PushInt( a.Int() / b.Int() )
    case VAR_INT:
      m.PushInt( a.Int() / b.Int() )
    case VAR_FLOAT:
      m.PushFloat( a.Float() / b.Float() )
    default:
      return errors.New( "Unsupported type for div" )
  }

  return nil
}

func intHandler( m *Context, n *Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_BOOL:
      m.PushInt( a.Int() )
    case VAR_INT:
      m.PushInt( a.Int() )
    case VAR_FLOAT:
      m.PushInt( a.Int() )
    default:
      return errors.New( "Unsupported type for int" )
  }

  return nil
}
