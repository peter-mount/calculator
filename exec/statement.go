package exec

import (
  "errors"
  "github.com/peter-mount/calculator/context"
)

func SetVarHandler( m *context.Context, n *context.Node ) error {
  err := n.InvokeLhs( m )
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  m.SetVar( n.Token().Text(), a )

  return nil
}

func GetVarHandler( m *context.Context, n *context.Node ) error {

  val := m.GetVar( n.Token().Text() )
  if val == nil {
    return errors.New( "Unknown variable " + n.Token().Text() )
  }

  m.Push( val )
  return nil
}

// invokeAllHandler Invokes all nodes within the supplied list
func InvokeAllHandler( m *context.Context, n *context.Node ) error {
  return n.ForEach( func(n1 *context.Node) error {
    return n1.Invoke(m)
  } )
}

// invokeScopeHandler calls invokeAllHandler with a variable scope that lasts
// for the duration of the call
func InvokeScopeHandler( m *context.Context, n *context.Node ) error {
  m.StartScope()
  defer m.EndScope()
  return InvokeAllHandler( m, n )
}
