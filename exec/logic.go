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
func Equal( a *context.Value, b *context.Value ) (*context.Value,error) {
  if a==nil {
    return context.BoolValue( b == nil ), nil
  }

  return context.BoolValue( a.Equal( b ) ), nil
}

// Compare two values for inequality
func NotEqual( a *context.Value, b *context.Value ) (*context.Value,error) {
  if a==nil {
    return context.BoolValue( b != nil ), nil
  }

  return context.BoolValue( !a.Equal( b ) ), nil
}

// a < b
func LessThan( a *context.Value, b *context.Value ) (*context.Value,error) {
  if a != nil && b != nil {
    switch a.Type() {
      case context.VAR_INT:
        return context.BoolValue( a.Int() < b.Int() ), nil
      case context.VAR_FLOAT:
        return context.BoolValue( a.Float() < b.Float() ), nil
      case context.VAR_STRING:
        return context.BoolValue( a.String() < b.String() ), nil
    }
  }
  return context.BoolValue( false ), nil
}

// a <= b
func LessThanEqual( a *context.Value, b *context.Value ) (*context.Value,error) {
  if a != nil && b != nil {
    switch a.Type() {
      case context.VAR_INT:
        return context.BoolValue( a.Int() <= b.Int() ), nil
      case context.VAR_FLOAT:
        return context.BoolValue( a.Float() <= b.Float() ), nil
      case context.VAR_STRING:
        return context.BoolValue( a.String() <= b.String() ), nil
    }
  }
  return context.BoolValue( false ), nil
}

// a >= b
func GreaterThanEqual( a *context.Value, b *context.Value ) (*context.Value,error) {
  if a != nil && b != nil {
    switch a.Type() {
      case context.VAR_INT:
        return context.BoolValue( a.Int() >= b.Int() ), nil
      case context.VAR_FLOAT:
        return context.BoolValue( a.Float() >= b.Float() ), nil
      case context.VAR_STRING:
        return context.BoolValue( a.String() >= b.String() ), nil
    }
  }
  return context.BoolValue( false ), nil
}

// a > b
func GreaterThan( a *context.Value, b *context.Value ) (*context.Value,error) {
  if a != nil && b != nil {
    switch a.Type() {
      case context.VAR_INT:
        return context.BoolValue( a.Int() > b.Int() ), nil
      case context.VAR_FLOAT:
        return context.BoolValue( a.Float() > b.Float() ), nil
      case context.VAR_STRING:
        return context.BoolValue( a.String() > b.String() ), nil
    }
  }
  return context.BoolValue( false ), nil
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
