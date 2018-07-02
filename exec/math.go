package exec

import (
  "errors"
  "github.com/peter-mount/calculator/context"
  "math"
  "math/cmplx"
)

var (
  unsupportedMathType = errors.New( "Unsupported number type" )
)

// RealComplex1 invokes a function based on the value type.
// If a is numeric then rf will be used.
// If a is complex then cf will be used.
// An error will be returned if the value is not numeric or complex
func RealComplex1( rf func(float64) float64, cf func(complex128) complex128, a *context.Value ) (*context.Value,error) {
  if a != nil {
    if rf != nil && a.IsNumeric() {
      return context.FloatValue( rf( a.Float() ) ), nil
    }

    if cf != nil && a.IsComplex() {
      return context.ComplexValue( cf( a.Complex() ) ), nil
    }
  }

  return nil, unsupportedMathType
}

// Abs returns the absolute value of x.
func Abs( a *context.Value ) (*context.Value,error) {
  if a != nil {
    if a.IsNumeric() {
      return context.FloatValue( math.Abs( a.Float() ) ), nil
    }

    if a.IsComplex() {
      return context.FloatValue( cmplx.Abs( a.Complex() ) ), nil
    }
  }

  return nil, unsupportedMathType
}

// Acos returns the arccosine, in radians, of x.
func Acos( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Acos, cmplx.Acos, a )
  return r, err
}

// Acosh returns the inverse hyperbolic cosine of x.
func Acosh( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Acosh, cmplx.Acosh, a )
  return r, err
}

// Asin returns the arcsine, in radians, of x.
func Asin( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Asin, cmplx.Asin, a )
  return r, err
}

// Asinh returns the inverse hyperbolic sine of x.
func Asinh( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Asinh, cmplx.Asinh, a )
  return r, err
}

// Atan returns the arctangent, in radians, of x.
func Atan( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Atan, cmplx.Atan, a )
  return r, err
}

// Atanh returns the inverse hyperbolic tangent of x.
func Atanh( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Atanh, cmplx.Atanh, a )
  return r, err
}

// Cbrt returns the cube root of x.
func Cbrt( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Cbrt, nil, a )
  return r, err
}

// Ceil returns the least integer value greater than or equal to x.
func Ceil( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Ceil, nil, a )
  return r, err
}

// Conj returns the complex conjugate of x.
func Conj( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( nil, cmplx.Conj, a )
  return r, err
}

// Cos returns the cosine of the radian argument x.
func Cos( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Cos, cmplx.Cos, a )
  return r, err
}

// Cosh returns the hyperbolic cosine of x.
func Cosh( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Cosh, cmplx.Cosh, a )
  return r, err
}

// Cot returns the cotangent of x.
func Cot( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( nil, cmplx.Cot, a )
  return r, err
}

// Erf returns the error function of x.
func Erf( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Erf, nil, a )
  return r, err
}

// Erfc returns the complementary error function of x.
func Erfc( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Erfc, nil, a )
  return r, err
}

// Erfcinv returns the inverse of Erfc(x).
func Erfcinv( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Erfcinv, nil, a )
  return r, err
}

// Erfinv returns the inverse error function of x.
func Erfinv( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Erfinv, nil, a )
  return r, err
}

// Exp returns e**x, the base-e exponential of x.
func Exp( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Exp, nil, a )
  return r, err
}

// Exp2 returns 2**x, the base-2 exponential of x.
func Exp2( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Exp2, nil, a )
  return r, err
}

// Expm1 returns e**x - 1, the base-e exponential of x minus 1. It is more accurate than Exp(x) - 1 when x is near zero.
func Expm1( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Expm1, nil, a )
  return r, err
}

// Floor returns the greatest integer value less than or equal to x.
func Floor( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Floor, nil, a )
  return r, err
}

// Gamma returns the Gamma function of x.
func Gamma( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Gamma, nil, a )
  return r, err
}

// Ilogb returns the binary exponent of x as an integer.
func Ilogb( a *context.Value ) (*context.Value,error) {
  return context.IntValue( int64( math.Ilogb( a.Float() ) ) ), nil
}

// Imag returns the imaginary component of a complex number
func Imag( a *context.Value ) (*context.Value,error) {
  if a != nil {
    return context.FloatValue( a.Imaginary() ), nil
  }

  return nil, unsupportedMathType
}

// IsNaN reports whether f is an IEEE 754 “not-a-number” value.
func IsNaN( a *context.Value ) (*context.Value,error) {
  if a != nil {
    if a.IsNumeric() {
      return context.BoolValue( math.IsNaN( a.Float() ) ), nil
    }

    if a.IsComplex() {
      return context.BoolValue( cmplx.IsNaN( a.Complex() ) ), nil
    }
  }
  return nil, unsupportedMathType
}

// J0 returns the order-zero Bessel function of the first kind.
func J0( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.J0, nil, a )
  return r, err
}

// J1 returns the order-one Bessel function of the first kind.
func J1( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.J1, nil, a )
  return r, err
}

// Log returns the natural logarithm of x.
func Log( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Log, cmplx.Log, a )
  return r, err
}

// Log10 returns the decimal logarithm of x
func Log10( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Log10, cmplx.Log10, a )
  return r, err
}

// Log1p returns the natural logarithm of 1 plus its argument x.
// It is more accurate than Log(1 + x) when x is near zero.
func Log1p( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Log1p, nil, a )
  return r, err
}

// Log2 returns the binary logarithm of x.
func Log2( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Log2, nil, a )
  return r, err
}

// Logb returns the binary exponent of x.
func Logb( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Logb, nil, a )
  return r, err
}

// Pow10 returns 10**n, the base-10 exponential of n.
func Phase( a *context.Value ) (*context.Value,error) {
  if a != nil && a.IsComplex() {
    return context.FloatValue( cmplx.Phase( a.Complex() ) ), nil
  }
  return nil, unsupportedMathType
}

// Pow10 returns 10**n, the base-10 exponential of n.
func Pow10( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Pow10( int(a.Int()) ) ), nil
}

// Real returns the real component of a complex number
func Real( a *context.Value ) (*context.Value,error) {
  if a != nil {
    return context.FloatValue( a.Real() ), nil
  }

  return nil, unsupportedMathType
}

// Round returns the nearest integer, rounding half away from zero.
func Round( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Round, nil, a )
  return r, err
}

// RoundToEven returns the nearest integer, rounding ties to even.
func Round2even( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.RoundToEven, nil, a )
  return r, err
}

// Sin returns the sine of the radian argument x.
func Sin( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Sin, cmplx.Sin, a )
  return r, err
}

// Sinh returns the hyperbolic sine of x.
func Sinh( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Sinh, cmplx.Sinh, a )
  return r, err
}

// Sqrt returns the square root of x.
func Sqrt( a *context.Value ) (*context.Value,error) {
  // Special case: sqrt of a real negative value then convert to complex
  // so sqrt(-1) returns (0+1i)
  if a != nil && a.IsNegative() {
    return context.ComplexValue( cmplx.Sqrt( complex( a.Float(), 0 ) ) ), nil
  }
  r, err := RealComplex1( math.Sqrt, cmplx.Sqrt, a )
  return r, err
}

// Tan returns the tangent of the radian argument x.
func Tan( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Tan, cmplx.Tan, a )
  return r, err
}

// Tanh returns the hyperbolic tangent of x.
func Tanh( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Tanh, cmplx.Tanh, a )
  return r, err
}

// Trunc returns the integer value of x.
func Trunc( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Trunc, nil, a )
  return r, err
}

// Y0 returns the order-zero Bessel function of the second kind.
func Y0( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Y0, nil, a )
  return r, err
}

// Y1 returns the order-one Bessel function of the second kind.
func Y1( a *context.Value ) (*context.Value,error) {
  r, err := RealComplex1( math.Y1, nil, a )
  return r, err
}
