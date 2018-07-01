package test

import (
  "fmt"
  "github.com/peter-mount/calculator/context"
  "testing"
)

// test v is a null
func TestValue_NullValue( t *testing.T ) {
  v := context.NullValue()
  if v.Type() != context.VAR_NULL { t.Error( "Type" ) }
  if !v.IsZero()          { t.Error( "IsZero" ) }
  if v.IsNumeric()        { t.Error( "IsNumeric" ) }
  if v.Bool()             { t.Error( "Bool" ) }
  if v.Int() != 0         { t.Error( "Int" ) }
  if v.Float() != 0.0     { t.Error( "Float" ) }
  if v.String() != ""     { t.Errorf( "String \"%s\"", v.String() ) }
}

// Test false
func TestValue_Bool_false( t *testing.T ) {
  v := context.BoolValue( false )

  if v.Type() != context.VAR_BOOL { t.Error( "Type" ) }
  if !v.IsZero()          { t.Error( "IsZero" ) }
  if v.IsNumeric()        { t.Error( "IsNumeric" ) }
  if v.Bool()             { t.Error( "Bool" ) }
  if v.Int() != 0         { t.Error( "Int" ) }
  if v.Float() != 0.0     { t.Error( "Float" ) }
  if v.String() != "false"    { t.Errorf( "String \"%s\"", v.String() ) }
}

// Test true
func TestValue_Bool_true( t *testing.T ) {
  v := context.BoolValue( true )

  if v.Type() != context.VAR_BOOL { t.Error( "Type" ) }
  if v.IsZero()           { t.Error( "IsZero" ) }
  if v.IsNumeric()        { t.Error( "IsNumeric" ) }
  if !v.Bool()            { t.Error( "Bool" ) }
  if v.Int() != 1         { t.Error( "Int" ) }
  if v.Float() != 1.0     { t.Error( "Float" ) }
  if v.String() != "true" { t.Errorf( "String \"%s\"", v.String() ) }
}

// Test int 0
func TestValue_Int_0( t *testing.T ) {
  v := context.IntValue( 0 )

  if v.Type() != context.VAR_INT  { t.Error( "Type" ) }
  if !v.IsZero()          { t.Error( "IsZero" ) }
  if !v.IsNumeric()       { t.Error( "IsNumeric" ) }
  if v.Bool()             { t.Error( "Bool" ) }
  if v.Int() != 0         { t.Error( "Int" ) }
  if v.Float() != 0.0     { t.Error( "Float" ) }
  if v.String() != "0"    { t.Errorf( "String \"%s\"", v.String() ) }
}

// Test int 1
func TestValue_Int_1( t *testing.T ) {
  v := context.IntValue( 1 )

  if v.Type() != context.VAR_INT  { t.Error( "Type" ) }
  if v.IsZero()          { t.Error( "IsZero" ) }
  if !v.IsNumeric()       { t.Error( "IsNumeric" ) }
  if !v.Bool()             { t.Error( "Bool" ) }
  if v.Int() != 1         { t.Error( "Int" ) }
  if v.Float() != 1.0     { t.Error( "Float" ) }
  if v.String() != "1"    { t.Errorf( "String \"%s\"", v.String() ) }
}

// Test integers -100 to 100
func TestValue_Int_pm100( t *testing.T ) {
  for i := -100; i <= 100; i++ {
    v := context.IntValue( int64(i) )

    if ( i==0 && !v.IsZero() ) || ( i!=0 && v.IsZero() ) { t.Errorf( "IsZero %d", i ) }
    if ( i==0 && v.Bool() ) || ( i!=0 && !v.Bool() ) { t.Errorf( "Bool %d", i ) }
    if v.Int() != int64(i)         { t.Errorf( "Int %d %d", i, v.Int() ) }
    if v.Float() != float64(i)     { t.Errorf( "Float %d %f", i, v.Float() ) }
    if v.String() != fmt.Sprintf("%d",i) { t.Errorf( "String %d \"%s\"", i, v.String() ) }
  }
}

// Test int 0
func TestValue_Float_0( t *testing.T ) {
  v := context.FloatValue( 0 )

  if v.Type() != context.VAR_FLOAT  { t.Error( "Type" ) }
  if !v.IsZero()            { t.Error( "IsZero" ) }
  if !v.IsNumeric()         { t.Error( "IsNumeric" ) }
  if v.Bool()               { t.Error( "Bool" ) }
  if v.Int() != 0           { t.Error( "Int" ) }
  if v.Float() != 0.0       { t.Error( "Float" ) }
  if v.String() != "0.0000000000"      { t.Errorf( "String \"%s\"", v.String() ) }
}

// Test int 1
func TestValue_Float_1( t *testing.T ) {
  v := context.FloatValue( 1 )

  if v.Type() != context.VAR_FLOAT  { t.Error( "Type" ) }
  if v.IsZero()          { t.Error( "IsZero" ) }
  if !v.IsNumeric()       { t.Error( "IsNumeric" ) }
  if !v.Bool()             { t.Error( "Bool" ) }
  if v.Int() != 1         { t.Error( "Int" ) }
  if v.Float() != 1.0     { t.Error( "Float" ) }
  if v.String() != "1.0000000000"    { t.Errorf( "String \"%s\"", v.String() ) }
}

// Test integers -100 to 100
func TestValue_Float_pm100( t *testing.T ) {
  for i := -100.0; i <= 100.0; i+=0.1 {
    v := context.FloatValue( float64(i) )

    if ( i==0 && !v.IsZero() ) || ( i!=0 && v.IsZero() ) { t.Errorf( "IsZero %f", i ) }
    if ( i==0 && v.Bool() ) || ( i!=0 && !v.Bool() ) { t.Errorf( "Bool %f", i ) }
    if v.Int() != int64(i)         { t.Errorf( "Int %f %d", i, v.Int() ) }
    if v.Float() != float64(i)     { t.Errorf( "Float %f %f", i, v.Float() ) }
    if v.String() != fmt.Sprintf("%.10f",i) { t.Errorf( "String %f \"%s\"", i, v.String() ) }
  }
}

func Test_Value_Complex( t *testing.T ) {
  c := complex( 1, 2 )
  vc := context.ComplexValue( c )
  if vc.Type() != context.VAR_COMPLEX {
    t.Error( "Wrong type")
  }
  if vc.IsNumeric() {
    t.Error( "Is numeric")
  }
  if !vc.IsComplex() {
    t.Error( "Not complex")
  }
  if vc.Complex() != c {
    t.Error( "Wrong complex value")
  }
  if vc.Real() != 1 {
    t.Error( "Wrong real value")
  }
  if vc.Imaginary() != 2 {
    t.Error( "Wrong imaginary value")
  }
}

func Test_Value_Complex_0( t *testing.T ) {
  c := complex( 0, 0 )
  vc := context.ComplexValue( c )
  if !vc.IsZero() {
    t.Error( "Not zero")
  }
}

func Test_Value_Complex_0_1i( t *testing.T ) {
  c := complex( 0, 1 )
  vc := context.ComplexValue( c )
  if vc.IsZero() {
    t.Error( "Is zero")
  }
}

func Test_Value_Complex_1_0i( t *testing.T ) {
  c := complex( 1, 0 )
  vc := context.ComplexValue( c )
  if vc.IsZero() {
    t.Error( "Is zero")
  }
}

func Test_Value_Complex_1_1i( t *testing.T ) {
  c := complex( 1, 1 )
  vc := context.ComplexValue( c )
  if vc.IsZero() {
    t.Error( "Is zero")
  }
}
