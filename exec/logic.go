package exec

import (
  "github.com/peter-mount/calculator/context"
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
        expr = context.NewNode( token, equalHandler, expr, right )
      case "!=":
        expr = context.NewNode( token, notEqualHandler, expr, right )
    }

    token = p.lexer.Peek()
  }

  return expr, err
}

func trueHandler( m *context.Context, n *context.Node ) error {
  m.PushBool( true )
  return nil
}

func falseHandler( m *context.Context, n *context.Node ) error {
  m.PushBool( false )
  return nil
}

func nullHandler( m *context.Context, n *context.Node ) error {
  m.PushNull()
  return nil
}

func equalHandler( m *context.Context, n *context.Node ) error {
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

func notEqualHandler( m *context.Context, n *context.Node ) error {
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

func lessThanHandler( m *context.Context, n *context.Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
  case context.VAR_INT:
      m.PushBool( a.Int() < b.Int() )
    case context.VAR_FLOAT:
      m.PushBool( a.Float() < b.Float() )
    case context.VAR_STRING:
      m.PushBool( a.String() < b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func lessThanEqualHandler( m *context.Context, n *context.Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
    case context.VAR_INT:
      m.PushBool( a.Int() <= b.Int() )
    case context.VAR_FLOAT:
      m.PushBool( a.Float() <= b.Float() )
    case context.VAR_STRING:
      m.PushBool( a.String() <= b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func greaterThanEqualHandler( m *context.Context, n *context.Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
    case context.VAR_INT:
      m.PushBool( a.Int() >= b.Int() )
    case context.VAR_FLOAT:
      m.PushBool( a.Float() >= b.Float() )
    case context.VAR_STRING:
      m.PushBool( a.String() >= b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func greaterThanHandler( m *context.Context, n *context.Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
    case context.VAR_INT:
      m.PushBool( a.Int() > b.Int() )
    case context.VAR_FLOAT:
      m.PushBool( a.Float() > b.Float() )
    case context.VAR_STRING:
      m.PushBool( a.String() > b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func betweenHandler( m *context.Context, n *context.Node ) error {
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
    case context.VAR_INT:
      ci := c.Int()
      m.PushBool( ci >= a.Int() && ci <= b.Int() )
    case context.VAR_FLOAT:
      cf := c.Float()
      m.PushBool( cf >= a.Float() && cf <= b.Float() )
    default:
      m.PushBool( false )
  }
  return nil
}

func andHandler( m *context.Context, n *context.Node ) error {
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

func orHandler( m *context.Context, n *context.Node ) error {
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

func notHandler( m *context.Context, n *context.Node ) error {
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
