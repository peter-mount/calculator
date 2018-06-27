package test

import (
  "github.com/peter-mount/calculator/calculator"
  "github.com/peter-mount/calculator/context"
  "testing"
)

// Test basic math precedence
func TestExec( t *testing.T ) {

  calc := &calculator.Calculator{}

  err := calc.Parse( "1 + 1" )
  if err != nil {
    t.Error( err )
  }

  ctx := &context.Context{}
  err = calc.Execute( ctx )
  if err != nil {
    t.Error( err )
  }

  result, err := ctx.Pop()
  if err != nil {
    t.Error( err )
  } else if result.Int() != 2 {
    t.Errorf( "Unexpected result, expected 2 got %v", result )
  }
}
