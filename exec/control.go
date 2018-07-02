package exec

import (
  "errors"
  "github.com/peter-mount/calculator/context"
)

var (
  noConditionError = errors.New( "No condition" )
)

func evalCondition( m *context.Context, n *context.Node ) (bool,error) {
  if n == nil {
    return false, noConditionError
  }

  err := n.Invoke( m )
  if err != nil {
    return false, err
  }

  a, err := m.Pop()
  if err != nil {
    return false, err
  }

  return a.Bool(), nil
}

func IfHandler( m *context.Context, n *context.Node ) error {
  cond, err := evalCondition( m, n.Left() )
  if err != nil {
    return err
  }

  if cond {
    err = n.InvokeRhs( m )
  } else {
    err = InvokeScopeHandler( m, n )
  }

  return err
}

func WhileHandler( m *context.Context, n *context.Node ) error {
  cond, err := evalCondition( m, n.Left() )
  if err != nil {
    return err
  }

  for cond {
    err = n.InvokeRhs( m )
    if err != nil {
      return err
    }

    cond, err = evalCondition( m, n.Left() )
    if err != nil {
      return err
    }
  }

  return err
}
