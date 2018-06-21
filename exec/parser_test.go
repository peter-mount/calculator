package exec

import (
  "testing"
)

// Test push/pop works & we detect an empty stack
func TestParser_1( t *testing.T ) {

  calc := &Calculator{}

  parser := calc.Parser()

  //err := parser.Parse( "true" )
  //err := parser.Parse( "1 && 2 == 3" )
  err := parser.Parse( "1 + 2" )
  if err != nil {
    t.Error( err )
  }

  parser.GetRoot().LogTree()

  //ctx := &Context{}
}
