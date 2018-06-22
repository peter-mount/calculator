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

  expr, err := p.parse_parens()
  if err != nil {
    return nil, err
  }

  token := p.lexer.Peek()
  for token.text == "*" || token.text == "/" {
    token = p.lexer.Next()

    right, err := p.parse_parens()
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

func addHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
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

  switch a.Type() {
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

  switch a.Type() {
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

  switch a.Type() {
    case VAR_BOOL:
      m.PushInt( a.Int() / b.Int() )
    case VAR_INT:
      m.PushInt( a.Int() / b.Int() )
    case VAR_FLOAT:
      m.PushFloat( a.Float() / b.Float() )
    default:
      return errors.New( "Unsupported type for sub" )
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
