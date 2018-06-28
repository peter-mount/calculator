package exec

import (
  "github.com/peter-mount/calculator/context"
  "math"
)

// Abs returns the absolute value of x.
func Abs( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Abs( a.Float() ) ), nil
}

// Acos returns the arccosine, in radians, of x.
func Acos( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Acos( a.Float() ) ), nil
}

// Acosh returns the inverse hyperbolic cosine of x.
func Acosh( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Acosh( a.Float() ) ), nil
}

// Asin returns the arcsine, in radians, of x.
func Asin( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Asin( a.Float() ) ), nil
}

// Asinh returns the inverse hyperbolic sine of x.
func Asinh( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Asinh( a.Float() ) ), nil
}

// Atan returns the arctangent, in radians, of x.
func Atan( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Atan( a.Float() ) ), nil
}

// Atanh returns the inverse hyperbolic tangent of x.
func Atanh( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Atanh( a.Float() ) ), nil
}

// Cbrt returns the cube root of x.
func Cbrt( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Cbrt( a.Float() ) ), nil
}

// Ceil returns the least integer value greater than or equal to x.
func Ceil( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Ceil( a.Float() ) ), nil
}

// Cos returns the cosine of the radian argument x.
func Cos( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Cos( a.Float() ) ), nil
}

// Cosh returns the hyperbolic cosine of x.
func Cosh( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Cosh( a.Float() ) ), nil
}

// Erf returns the error function of x.
func Erf( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Erf( a.Float() ) ), nil
}

// Erfc returns the complementary error function of x.
func Erfc( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Erfc( a.Float() ) ), nil
}

// Erfcinv returns the inverse of Erfc(x).
func Erfcinv( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Erfcinv( a.Float() ) ), nil
}

// Erfinv returns the inverse error function of x.
func Erfinv( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Erfinv( a.Float() ) ), nil
}

// Exp returns e**x, the base-e exponential of x.
func Exp( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Exp( a.Float() ) ), nil
}

// Exp2 returns 2**x, the base-2 exponential of x.
func Exp2( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Exp2( a.Float() ) ), nil
}

// Expm1 returns e**x - 1, the base-e exponential of x minus 1. It is more accurate than Exp(x) - 1 when x is near zero.
func Expm1( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Expm1( a.Float() ) ), nil
}

// Floor returns the greatest integer value less than or equal to x.
func Floor( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Floor( a.Float() ) ), nil
}

// Gamma returns the Gamma function of x.
func Gamma( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Gamma( a.Float() ) ), nil
}

// Ilogb returns the binary exponent of x as an integer.
func Ilogb( a *context.Value ) (*context.Value,error) {
  return context.IntValue( int64( math.Ilogb( a.Float() ) ) ), nil
}

// IsNaN reports whether f is an IEEE 754 “not-a-number” value.
func IsNaN( a *context.Value ) (*context.Value,error) {
  return context.BoolValue( math.IsNaN( a.Float() ) ), nil
}

// J0 returns the order-zero Bessel function of the first kind.
func J0( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.J0( a.Float() ) ), nil
}

// J1 returns the order-one Bessel function of the first kind.
func J1( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.J1( a.Float() ) ), nil
}

// Log returns the natural logarithm of x.
func Log( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Log( a.Float() ) ), nil
}

// Log10 returns the decimal logarithm of x
func Log10( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Log10( a.Float() ) ), nil
}

// Log1p returns the natural logarithm of 1 plus its argument x.
// It is more accurate than Log(1 + x) when x is near zero.
func Log1p( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Log1p( a.Float() ) ), nil
}

// Log2 returns the binary logarithm of x.
func Log2( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Log2( a.Float() ) ), nil
}

// Logb returns the binary exponent of x.
func Logb( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Logb( a.Float() ) ), nil
}

// Pow10 returns 10**n, the base-10 exponential of n.
func Pow10( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Pow10( int(a.Int()) ) ), nil
}

// Round returns the nearest integer, rounding half away from zero.
func Round( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Round( a.Float() ) ), nil
}

// RoundToEven returns the nearest integer, rounding ties to even.
func Round2even( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.RoundToEven( a.Float() ) ), nil
}

// Sin returns the sine of the radian argument x.
func Sin( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Sin( a.Float() ) ), nil
}

// Sinh returns the hyperbolic sine of x.
func Sinh( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Sinh( a.Float() ) ), nil
}

// Sqrt returns the square root of x.
func Sqrt( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Sqrt( a.Float() ) ), nil
}

// Tan returns the tangent of the radian argument x.
func Tan( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Tan( a.Float() ) ), nil
}

// Tanh returns the hyperbolic tangent of x.
func Tanh( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Tanh( a.Float() ) ), nil
}

// Trunc returns the integer value of x.
func Trunc( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Trunc( a.Float() ) ), nil
}

// Y0 returns the order-zero Bessel function of the second kind.
func Y0( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Y0( a.Float() ) ), nil
}

// Y1 returns the order-one Bessel function of the second kind.
func Y1( a *context.Value ) (*context.Value,error) {
  return context.FloatValue( math.Y1( a.Float() ) ), nil
}
