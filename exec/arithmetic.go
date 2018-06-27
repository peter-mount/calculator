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
func Add( a *context.Value, b *context.Value ) (*context.Value,error) {
  if a != nil && b != nil {
    switch a.OperationType( b ) {
      case context.VAR_BOOL:
        return context.IntValue( a.Int() + b.Int() ), nil
      case context.VAR_INT:
        return context.IntValue( a.Int() + b.Int() ), nil
      case context.VAR_FLOAT:
        return context.FloatValue( a.Float() + b.Float() ), nil
      case context.VAR_STRING:
        return context.StringValue( a.String() + b.String() ), nil
    }
  }

  return nil, errors.New( "Unsupported type for add" )
}

// Subtraction of two values
// If one value is a float then the result will be a float.
// If an Integer then the result will be an integer.
// For booleans then it's integer equivalent is used.
func Sub( a *context.Value, b *context.Value ) (*context.Value,error) {
  if a != nil && b != nil {
    switch a.OperationType( b ) {
      case context.VAR_BOOL:
        return context.IntValue( a.Int() - b.Int() ), nil
      case context.VAR_INT:
        return context.IntValue( a.Int() - b.Int() ), nil
      case context.VAR_FLOAT:
        return context.FloatValue( a.Float() - b.Float() ), nil
      default:
    }
  }

  return nil, errors.New( "Unsupported type for sub" )
}

// Multiplication of two values.
// If one value is a float then the result will be a float.
// If an Integer then the result will be an integer.
// For booleans then it's integer equivalent is used.
func Mult( a *context.Value, b *context.Value ) (*context.Value,error) {
  if a != nil && b != nil {
    switch a.OperationType( b ) {
      case context.VAR_BOOL:
        return context.IntValue( a.Int() * b.Int() ), nil
      case context.VAR_INT:
        return context.IntValue( a.Int() * b.Int() ), nil
      case context.VAR_FLOAT:
        return context.FloatValue( a.Float() * b.Float() ), nil
      }
  }
  return nil, errors.New( "Unsupported type for mult" )
}

// Division of two values.
// If one value is a float then the result will be a float.
// If an Integer then the result will be an integer.
// For booleans then it's integer equivalent is used.
// An error will be returned is a divide-by-zero would occur.
func Div( a *context.Value, b *context.Value ) (*context.Value,error) {
  if a != nil && b != nil {
    if b.IsZero() {
      return nil, errors.New( "Division by zero")
    }

    switch a.OperationType( b ) {
      case context.VAR_BOOL:
        return context.IntValue( a.Int() / b.Int() ), nil
      case context.VAR_INT:
        return context.FloatValue( a.Float() / b.Float() ), nil
      case context.VAR_FLOAT:
        return context.FloatValue( a.Float() / b.Float() ), nil
    }
  }

  return nil, errors.New( "Unsupported type for div" )
}
