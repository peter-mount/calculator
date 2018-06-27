package exec

func (p *Parser) parse_logic() (*Node,error) {

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
        expr = &Node{ token:token.Text(), left:expr, right: right, handler: equalHandler }
      case "!=":
        expr = &Node{ token:token.Text(), left:expr, right: right, handler: notEqualHandler }
    }

    token = p.lexer.Peek()
  }

  return expr, err
}

func trueHandler( m *Context, n *Node ) error {
  m.Push( &trueValue )
  return nil
}

func falseHandler( m *Context, n *Node ) error {
  m.Push( &falseValue )
  return nil
}

func nullHandler( m *Context, n *Node ) error {
  m.Push( &nullValue )
  return nil
}

func equalHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  if a==nil {
    m.PushBool( b == nil )
  } else {
    m.PushBool( a.Equal( b ) )
  }

  return nil
}

func notEqualHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  if a==nil {
    m.PushBool( b != nil )
  } else {
    m.PushBool( !a.Equal( b ) )
  }

  return nil
}

func lessThanHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_INT:
      m.PushBool( a.Int() < b.Int() )
    case VAR_FLOAT:
      m.PushBool( a.Float() < b.Float() )
    case VAR_STRING:
      m.PushBool( a.String() < b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func lessThanEqualHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_INT:
      m.PushBool( a.Int() <= b.Int() )
    case VAR_FLOAT:
      m.PushBool( a.Float() <= b.Float() )
    case VAR_STRING:
      m.PushBool( a.String() <= b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func greaterThanEqualHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_INT:
      m.PushBool( a.Int() >= b.Int() )
    case VAR_FLOAT:
      m.PushBool( a.Float() >= b.Float() )
    case VAR_STRING:
      m.PushBool( a.String() >= b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func greaterThanHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_INT:
      m.PushBool( a.Int() > b.Int() )
    case VAR_FLOAT:
      m.PushBool( a.Float() > b.Float() )
    case VAR_STRING:
      m.PushBool( a.String() > b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func betweenHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  // The values to be between
  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  // The value to test
  c, err := m.Pop()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_INT:
      ci := c.Int()
      m.PushBool( ci >= a.Int() && ci <= b.Int() )
    case VAR_FLOAT:
      cf := c.Float()
      m.PushBool( cf >= a.Float() && cf <= b.Float() )
    default:
      m.PushBool( false )
  }
  return nil
}

func andHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  m.PushBool( a.Bool() && b.Bool() )
  return nil
}

func orHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  m.PushBool( a.Bool() || b.Bool() )
  return nil
}

func notHandler( m *Context, n *Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  m.PushBool( !a.Bool() )
  return nil
}
