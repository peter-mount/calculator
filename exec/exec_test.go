package exec

import (
  //"fmt"
  "testing"
)

// Test basic math precedence
func TestExec( t *testing.T ) {

  calc := &Calculator{}

  err := calc.Parse( "1 + 1" )
  if err != nil {
    t.Error( err )
  }

  ctx := &Context{}
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
