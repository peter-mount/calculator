package exec

import (
  "errors"
  "github.com/peter-mount/calculator/context"
)

// Set a variable
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

// Get a variable. If the variable does not exist then an error is returned.
func GetVarHandler( m *context.Context, n *context.Node ) error {

  val := m.GetVar( n.Token().Text() )
  if val == nil {
    return errors.New( "Unknown variable " + n.Token().Text() )
  }

  m.Push( val )
  return nil
}

// InvokeAllHandler Invokes all nodes within the supplied list
func InvokeAllHandler( m *context.Context, n *context.Node ) error {
  return n.ForEach( func(n1 *context.Node) error {
    return n1.Invoke(m)
  } )
}

// InvokeScopeHandler calls invokeAllHandler with a variable scope that lasts
// for the duration of the call
func InvokeScopeHandler( m *context.Context, n *context.Node ) error {
  m.StartScope()
  defer m.EndScope()
  return InvokeAllHandler( m, n )
}
