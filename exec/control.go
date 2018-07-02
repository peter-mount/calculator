package exec

import (
  "errors"
  "github.com/peter-mount/calculator/context"
)

var (
  noConditionError = errors.New( "No condition" )
)

func loopTrue(c bool) bool { return c }
func loopFalse(c bool) bool { return !c }

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

// common code for While/Until
func whileLoop( m *context.Context, n *context.Node, f func(bool) bool ) error {
  cond, err := evalCondition( m, n.Left() )
  if err != nil {
    return err
  }

  return doLoop( m, n, cond, f )
}

// Common code for DoWhile & DoUntil
// cond initial value, f function to check cond if valud
func doLoop( m *context.Context, n *context.Node, cond bool, f func(bool) bool ) error {
  for f(cond) {
    err := n.InvokeRhs( m )
    if err != nil {
      return err
    }

    cond, err = evalCondition( m, n.Left() )
    if err != nil {
      return err
    }
  }

  return nil
}

func DoWhileHandler( m *context.Context, n *context.Node ) error {
  return doLoop( m, n, true, loopTrue )
}

func DoUntilHandler( m *context.Context, n *context.Node ) error {
  return doLoop( m, n, false, loopFalse )
}

func ForHandler( m *context.Context, n *context.Node ) error {
  m.StartScope()
  defer m.EndScope()

  err := n.InvokeLhs( m )
  if err != nil {
    return err
  }

  cond, err := evalCondition( m, n.Center() )
  if err != nil {
    return err
  }

  for cond {
    err = InvokeAllHandler( m, n )
    if err != nil {
      return err
    }

    err = n.InvokeRhs( m )
    if err != nil {
      return err
    }

    cond, err = evalCondition( m, n.Center() )
    if err != nil {
      return err
    }
  }

  return nil
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
  return whileLoop( m, n, loopTrue )
}

func UntilHandler( m *context.Context, n *context.Node ) error {
  return whileLoop( m, n, loopFalse )
}
