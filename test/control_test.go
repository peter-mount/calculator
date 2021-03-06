package test

import (
  "github.com/peter-mount/calculator/calculator"
  "github.com/peter-mount/calculator/context"
  "testing"
)

func TestDoUntil( t *testing.T ) {

  calc := &calculator.Calculator{}

  e := "$a=0;do { $a = $a + 1 } until $a >= 10"

  err := calc.ParseScriptString( e )
  if err != nil {
    t.Error( err )
  }

  ctx := &context.Context{}
  err = calc.Execute( ctx )
  if err != nil {
    t.Error( err )
  }

  a := ctx.GetVar( "a" )
  if a == nil {
    t.Error( "No output" )
  } else if !a.IsNumeric() || a.Int() != 10 {
    t.Errorf( "Expected 10 got %v", a )
  }
}

func TestDoWhile( t *testing.T ) {

  calc := &calculator.Calculator{}

  e := "$a=0;do { $a = $a + 1 } while $a < 10"

  err := calc.ParseScriptString( e )
  if err != nil {
    t.Error( err )
  }

  ctx := &context.Context{}
  err = calc.Execute( ctx )
  if err != nil {
    t.Error( err )
  }

  a := ctx.GetVar( "a" )
  if a == nil {
    t.Error( "No output" )
  } else if !a.IsNumeric() || a.Int() != 10 {
    t.Errorf( "Expected 10 got %v", a )
  }
}

func TestFor( t *testing.T ) {

  calc := &calculator.Calculator{}

  e := "$a=-1;for($i=0; $i <= 10; $i = $i + 1) { $a = $i }"

  err := calc.ParseScriptString( e )
  if err != nil {
    t.Error( err )
  }

  ctx := &context.Context{}
  err = calc.Execute( ctx )
  if err != nil {
    t.Error( err )
  }

  a := ctx.GetVar( "a" )
  if a == nil {
    t.Error( "No output" )
  } else if !a.IsNumeric() || a.Int() != 10 {
    t.Errorf( "Expected 10 got %v", a )
  }
}

func TestIf( t *testing.T ) {

  calc := &calculator.Calculator{}

  e := "$a=-1;if $a == -1 { $a=1 }"

  err := calc.ParseScriptString( e )
  if err != nil {
    t.Error( err )
  }

  ctx := &context.Context{}
  err = calc.Execute( ctx )
  if err != nil {
    t.Error( err )
  }

  a := ctx.GetVar( "a" )
  if a == nil {
    t.Error( "No output" )
  } else if !a.IsNumeric() || a.Int() != 1 {
    t.Errorf( "Expected 1 got %v", a )
  }
}

func TestIfElse( t *testing.T ) {

  calc := &calculator.Calculator{}

  e := "$a=-1;if $a != -1 { $a=1 } else { $a=2 }"

  err := calc.ParseScriptString( e )
  if err != nil {
    t.Error( err )
  }

  ctx := &context.Context{}
  err = calc.Execute( ctx )
  if err != nil {
    t.Error( err )
  }

  a := ctx.GetVar( "a" )
  if a == nil {
    t.Error( "No output" )
  } else if !a.IsNumeric() || a.Int() != 2 {
    t.Errorf( "Expected 2 got %v", a )
  }
}

func TestUntil( t *testing.T ) {

  calc := &calculator.Calculator{}

  e := "$a=0;until( $a >= 10 ) { $a = $a + 1 }"

  err := calc.ParseScriptString( e )
  if err != nil {
    t.Error( err )
  }

  ctx := &context.Context{}
  err = calc.Execute( ctx )
  if err != nil {
    t.Error( err )
  }

  a := ctx.GetVar( "a" )
  if a == nil {
    t.Error( "No output" )
  } else if !a.IsNumeric() || a.Int() != 10 {
    t.Errorf( "Expected 10 got %v", a )
  }
}

func TestWhile( t *testing.T ) {

  calc := &calculator.Calculator{}

  e := "$a=0;while( $a < 10 ) { $a = $a + 1 }"

  err := calc.ParseScriptString( e )
  if err != nil {
    t.Error( err )
  }

  ctx := &context.Context{}
  err = calc.Execute( ctx )
  if err != nil {
    t.Error( err )
  }

  a := ctx.GetVar( "a" )
  if a == nil {
    t.Error( "No output" )
  } else if !a.IsNumeric() || a.Int() != 10 {
    t.Errorf( "Expected 10 got %v", a )
  }
}
