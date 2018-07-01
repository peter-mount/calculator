package exec

import (
  "errors"
  "github.com/peter-mount/calculator/context"
)

// ToImaginary converts a value into a complex number using the value
// as the imaginary component.
func ToImaginary( a *context.Value ) (*context.Value,error) {
  if a != nil {

    if a.IsComplex() {
      return a, nil
    }

    if a.IsNumeric() {
      return context.ComplexValue( complex(0, a.Float()) ), nil
    }

  }

  return nil, errors.New( "Unsupported type for add" )
}
