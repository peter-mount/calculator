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
  "cbrt":     FuncMapEntry{ cbrtHandler,    UnaryOp  },
  "ceil":     FuncMapEntry{ ceilHandler,    UnaryOp  },
  "cos":      FuncMapEntry{ cosHandler,     UnaryOp  },
  "cosh":     FuncMapEntry{ coshHandler,   UnaryOp  },
  // Constants
  "e":        FuncMapEntry{ constE,         ActionOp  },
  "pi":       FuncMapEntry{ constPI,        ActionOp  },
  "phi":      FuncMapEntry{ constPHI,       ActionOp  },
  "sqrt2":    FuncMapEntry{ constSQRT2,     ActionOp  },
  "sqrte":    FuncMapEntry{ constSQRTE,     ActionOp  },
  "sqrtpi":   FuncMapEntry{ constSQRTPI,    ActionOp  },
  "sqrtphi":  FuncMapEntry{ constSQRTPHI,   ActionOp  },
  "ln2":      FuncMapEntry{ constLN2,       ActionOp  },
  "log2e":    FuncMapEntry{ constLOG2E,     ActionOp  },
  "ln10":     FuncMapEntry{ constLN10,      ActionOp  },
  "log10e":   FuncMapEntry{ constLOG10E,    ActionOp  },
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

func cbrtHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Cbrt )
}

func ceilHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Ceil )
}

func cosHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Cos )
}

func coshHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Cosh )
}

func constE( m *Context, n *Node ) error {
  m.PushFloat( math.E )
  return nil
}

func constPI( m *Context, n *Node ) error {
  m.PushFloat( math.Pi )
  return nil
}

func constPHI( m *Context, n *Node ) error {
  m.PushFloat( math.Phi )
  return nil
}

func constSQRT2( m *Context, n *Node ) error {
  m.PushFloat( math.Sqrt2 )
  return nil
}

func constSQRTE( m *Context, n *Node ) error {
  m.PushFloat( math.SqrtE )
  return nil
}

func constSQRTPI( m *Context, n *Node ) error {
  m.PushFloat( math.SqrtPi )
  return nil
}

func constSQRTPHI( m *Context, n *Node ) error {
  m.PushFloat( math.SqrtPhi )
  return nil
}

func constLN2( m *Context, n *Node ) error {
  m.PushFloat( math.Ln2 )
  return nil
}

func constLOG2E( m *Context, n *Node ) error {
  m.PushFloat( math.Log2E )
  return nil
}

func constLN10( m *Context, n *Node ) error {
  m.PushFloat( math.Ln10 )
  return nil
}

func constLOG10E( m *Context, n *Node ) error {
  m.PushFloat( math.Log10E )
  return nil
}
