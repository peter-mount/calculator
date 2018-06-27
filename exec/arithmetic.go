package exec

import (
  "errors"
  "github.com/peter-mount/calculator/context"
)

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

func IntHandler( m *context.Context, n *context.Node ) error {
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
