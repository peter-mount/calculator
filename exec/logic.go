package exec

import (
  "github.com/peter-mount/calculator/context"
)

// Boolean true
func TrueHandler( m *context.Context, n *context.Node ) error {
  m.PushBool( true )
  return nil
}

// Boolean false
func FalseHandler( m *context.Context, n *context.Node ) error {
  m.PushBool( false )
  return nil
}

// Null/Nil value
func NullHandler( m *context.Context, n *context.Node ) error {
  m.PushNull()
  return nil
}

// Compare two values for equality
func EqualHandler( m *context.Context, n *context.Node ) error {
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

// Compare two values for inequality
func NotEqualHandler( m *context.Context, n *context.Node ) error {
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

// a < b
func LessThanHandler( m *context.Context, n *context.Node ) error {
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

// a <= b
func LessThanEqualHandler( m *context.Context, n *context.Node ) error {
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

// a >= b
func GreaterThanEqualHandler( m *context.Context, n *context.Node ) error {
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

// a > b
func GreaterThanHandler( m *context.Context, n *context.Node ) error {
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

// a between b and c
func BetweenHandler( m *context.Context, n *context.Node ) error {
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

// a and b
func AndHandler( m *context.Context, n *context.Node ) error {
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

// a or b
func OrHandler( m *context.Context, n *context.Node ) error {
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

// !a
func NotHandler( m *context.Context, n *context.Node ) error {
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
