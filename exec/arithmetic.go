package exec

import (
  "errors"
  "github.com/peter-mount/calculator/context"
)

// NegHandler negates a value
func NegHandler( m *context.Context, n *context.Node ) error {
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

// Addition of 2 values.
// If one value is a float then the result will be a float.
// If an Integer then the result will be an integer.
// For booleans then it's integer equivalent is used.
// Strings will be concatenated.
func AddHandler( m *context.Context, n *context.Node ) error {
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

// Subtraction of two values
// If one value is a float then the result will be a float.
// If an Integer then the result will be an integer.
// For booleans then it's integer equivalent is used.
func SubHandler( m *context.Context, n *context.Node ) error {
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

// Multiplication of two values.
// If one value is a float then the result will be a float.
// If an Integer then the result will be an integer.
// For booleans then it's integer equivalent is used.
func MultHandler( m *context.Context, n *context.Node ) error {
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

// Division of two values.
// If one value is a float then the result will be a float.
// If an Integer then the result will be an integer.
// For booleans then it's integer equivalent is used.
// An error will be returned is a divide-by-zero would occur.
func DivHandler( m *context.Context, n *context.Node ) error {
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
