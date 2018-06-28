package exec

import (
  "github.com/peter-mount/calculator/context"
)

func IfHandler( m *context.Context, n *context.Node ) error {
  err := n.InvokeLhs( m )
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err == nil {
    if a.Bool() {
      err = n.InvokeRhs( m )
    } else {
      err = InvokeScopeHandler( m, n )
    }
  }


  return err
}
