package exec

import (
  "errors"
  "github.com/peter-mount/calculator/context"
  "math"
)

// Handles math functions that take 1 parameter
func mathInvoke1( m *context.Context, n *context.Node, f func(float64) float64 ) error {
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
  }
  return errors.New( "Unsupported type" )
}

// Handles math functions that take 2 parameters
func mathInvoke2( m *context.Context, n *context.Node, f func(float64,float64) float64 ) error {
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

func AbsHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Abs )
}

func AcosHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Acos )
}

func AcoshHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Acosh )
}

func AsinHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Asin )
}

func AsinhHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Asinh )
}

func AtanHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Atan )
}

func Atan2Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke2( m, n, math.Atan2 )
}

func AtanhHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Atanh )
}

func CbrtHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Cbrt )
}

func CeilHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Ceil )
}

func CosHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Cos )
}

func CoshHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Cosh )
}

func ErfHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Erf )
}

func ErfcHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Erfc )
}

func ErfinvHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Erfinv )
}

func ExpHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Exp )
}

func Exp2Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Exp2 )
}

func Expm1Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Expm1 )
}

func FloorHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Floor )
}

func IlogbHandler( m *context.Context, n *context.Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  if a.IsNumeric() {
    m.PushInt( int64( math.Ilogb( a.Float() ) ) )
    return nil
  }

  return errors.New( "Unsupported type" )
}

func IsNaNHandler( m *context.Context, n *context.Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  if a.IsNumeric() {
    m.PushBool( math.IsNaN( a.Float() ) )
    return nil
  }

  return errors.New( "Unsupported type" )
}

func J0Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.J0 )
}

func J1Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.J1 )
}

func LogHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Log )
}

func Log10Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Log10 )
}

func Log1pHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Log1p )
}

func Log2Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Log2 )
}

func LogbHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Logb )
}

func Pow10Handler( m *context.Context, n *context.Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  if a.IsNumeric() {
    m.PushFloat( math.Pow10( int(a.Int()) ) )
    return nil
  }

  return errors.New( "Unsupported type" )
}

func RoundHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Round )
}

func Round2evenHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.RoundToEven )
}

func SinHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Sin )
}

func SincosHandler( m *context.Context, n *context.Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  if a.IsNumeric() {
    s, c := math.Sincos( a.Float() )
    m.PushFloat( s )
    m.PushFloat( c )
    return nil
  }

  return errors.New( "Unsupported type" )
}

func SinhHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Sinh )
}

func SqrtHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Sqrt )
}

func TanHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Tan )
}

func TanhHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Tanh )
}

func TruncHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Trunc )
}

func Y0Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Y0 )
}

func Y1Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Y1 )
}
