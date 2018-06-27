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

// Abs returns the absolute value of x.
func AbsHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Abs )
}

// Acos returns the arccosine, in radians, of x.
func AcosHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Acos )
}

// Acosh returns the inverse hyperbolic cosine of x.
func AcoshHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Acosh )
}

// Asin returns the arcsine, in radians, of x.
func AsinHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Asin )
}

// Asinh returns the inverse hyperbolic sine of x.
func AsinhHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Asinh )
}

// Atan returns the arctangent, in radians, of x.
func AtanHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Atan )
}

// Atan2 returns the arc tangent of y/x, using the signs of the two to determine the quadrant of the return value.
func Atan2Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke2( m, n, math.Atan2 )
}

// Atanh returns the inverse hyperbolic tangent of x.
func AtanhHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Atanh )
}

// Cbrt returns the cube root of x.
func CbrtHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Cbrt )
}

// Ceil returns the least integer value greater than or equal to x.
func CeilHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Ceil )
}

// Cos returns the cosine of the radian argument x.
func CosHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Cos )
}

// Cosh returns the hyperbolic cosine of x.
func CoshHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Cosh )
}

// Erf returns the error function of x.
func ErfHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Erf )
}

// Erfc returns the complementary error function of x.
func ErfcHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Erfc )
}

// Erfcinv returns the inverse of Erfc(x).
func ErfcinvHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Erfcinv )
}

// Erfinv returns the inverse error function of x.
func ErfinvHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Erfinv )
}

// Exp returns e**x, the base-e exponential of x.
func ExpHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Exp )
}

// Exp2 returns 2**x, the base-2 exponential of x.
func Exp2Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Exp2 )
}

// Expm1 returns e**x - 1, the base-e exponential of x minus 1. It is more accurate than Exp(x) - 1 when x is near zero.
func Expm1Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Expm1 )
}

// Floor returns the greatest integer value less than or equal to x.
func FloorHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Floor )
}

// Gamma returns the Gamma function of x.
func GammaHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Gamma )
}

// Ilogb returns the binary exponent of x as an integer.
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

// IsNaN reports whether f is an IEEE 754 “not-a-number” value.
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

// J0 returns the order-zero Bessel function of the first kind.
func J0Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.J0 )
}

// J1 returns the order-one Bessel function of the first kind.
func J1Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.J1 )
}

// Log returns the natural logarithm of x.
func LogHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Log )
}

// Log10 returns the decimal logarithm of x
func Log10Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Log10 )
}

// Log1p returns the natural logarithm of 1 plus its argument x.
// It is more accurate than Log(1 + x) when x is near zero.
func Log1pHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Log1p )
}

// Log2 returns the binary logarithm of x.
func Log2Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Log2 )
}

// Logb returns the binary exponent of x.
func LogbHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Logb )
}

// Pow10 returns 10**n, the base-10 exponential of n.
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

// Round returns the nearest integer, rounding half away from zero.
func RoundHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Round )
}

// RoundToEven returns the nearest integer, rounding ties to even.
func Round2evenHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.RoundToEven )
}

// Sin returns the sine of the radian argument x.
func SinHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Sin )
}

// Sinh returns the hyperbolic sine of x.
func SinhHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Sinh )
}

// Sqrt returns the square root of x.
func SqrtHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Sqrt )
}

// Tan returns the tangent of the radian argument x.
func TanHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Tan )
}

// Tanh returns the hyperbolic tangent of x.
func TanhHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Tanh )
}

// Trunc returns the integer value of x.
func TruncHandler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Trunc )
}

// Y0 returns the order-zero Bessel function of the second kind.
func Y0Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Y0 )
}

// Y1 returns the order-one Bessel function of the second kind.
func Y1Handler( m *context.Context, n *context.Node ) error {
  return mathInvoke1( m, n, math.Y1 )
}
