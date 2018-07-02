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

  c, err := Neg( a )
  if err != nil {
    return err
  }

  m.Push( c )
  return nil
}

func Neg( a *context.Value ) (*context.Value,error) {
  switch a.Type() {
    case context.VAR_BOOL:
      return context.BoolValue( !a.Bool() ), nil
    case context.VAR_INT:
      return context.IntValue( -a.Int() ), nil
    case context.VAR_FLOAT:
      return context.FloatValue( -a.Float() ), nil
    case context.VAR_COMPLEX:
      return context.ComplexValue( -a.Complex() ), nil
    default:
      return nil, errors.New( "Unsupported type for neg" )
  }
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
      case context.VAR_COMPLEX:
        return context.ComplexValue( a.Complex() + b.Complex() ), nil
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
      case context.VAR_COMPLEX:
        return context.ComplexValue( a.Complex() - b.Complex() ), nil
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
      case context.VAR_COMPLEX:
        return context.ComplexValue( a.Complex() * b.Complex() ), nil
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
      case context.VAR_COMPLEX:
        return context.ComplexValue( a.Complex() / b.Complex() ), nil
    }
  }

  return nil, errors.New( "Unsupported type for div" )
}
