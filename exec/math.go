package exec

import (
  "errors"
  "math"
)

var mathFunctions = FuncMap{
  "abs":      FuncMapEntry{ absHandler,     UnaryOp  },
  "acos":     FuncMapEntry{ acosHandler,    UnaryOp  },
  "acosh":    FuncMapEntry{ acoshHandler,   UnaryOp  },
  "asin":     FuncMapEntry{ asinHandler,    UnaryOp  },
  "asinh":    FuncMapEntry{ asinhHandler,   UnaryOp  },
  "atan":     FuncMapEntry{ atanHandler,    UnaryOp  },
  //"atan2":      FuncMapEntry{ atan2Handler,             BinaryOp  },
  "atanh":    FuncMapEntry{ atanhHandler,   UnaryOp  },
}

// Handles math functions that take 1 parameter
func mathInvoke1( m *Context, n *Node, f func(float64) float64 ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  if a.IsNumeric() {
    m.PushFloat( f( a.Float() ) )
    return nil
  } else {
    return errors.New( "Unsupported type" )
  }
}

// Handles math functions that take 2 parameters
func mathInvoke2( m *Context, n *Node, f func(float64,float64) float64 ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  if a.IsNumeric() && b.IsNumeric() {
    m.PushFloat( f( a.Float(), b.Float() ) )
    return nil
  } else {
    return errors.New( "Unsupported type" )
  }
}

func absHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Abs )
}

func acosHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Acos )
}

func acoshHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Acosh )
}

func asinHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Asin )
}

func asinhHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Asinh )
}

func atanHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Atan )
}

func atan2Handler( m *Context, n *Node ) error {
  return mathInvoke2( m, n, math.Atan2 )
}

func atanhHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Atanh )
}
