package exec

import (
  "testing"
)

// Test push/pop works & we detect an empty stack
func TestContext_Pop( t *testing.T ) {
  ctx := &Context{}

  // Push a value & expect just 1 back
  ctx.Push( NullValue() )
  if v, err := ctx.Pop(); err != nil {
    t.Error( "Stack underflow" )
  } else if !v.IsNull() {
    t.Error( "Not null" )
  }

  // This should fail
  if _, err := ctx.Pop(); err == nil {
    t.Error( "Stack underflow expected" )
  }
}

// Test pop2 works
func TestContext_Pop2( t *testing.T ) {
  ctx := &Context{}

  // Push a value & expect just 1 back
  ctx.Push( IntValue(1) )
  ctx.Push( IntValue(2) )

  // This should pull them back off, a=1 & b=2
  a, b, err := ctx.Pop2()
  if err != nil {
    t.Error( "Stack underflow" )
  }
  if a.Int() != 1 {
    t.Errorf( "Expected a=1 got %d", a.Int() )
  }
  if b.Int() != 2 {
    t.Errorf( "Expected b=2 got %d", b.Int() )
  }

  // This should fail
  if _, err := ctx.Pop(); err == nil {
    t.Error( "Stack underflow expected" )
  }
}

func TestContext_PushBool( t *testing.T ) {
  ctx := &Context{}

  // Push a value & expect just 1 back
  ctx.PushBool( true )
  v, err := ctx.Pop()
  if err != nil {
    t.Error( "Stack underflow" )
  }
  if v.Type() != VAR_BOOL {
    t.Error( "Not bool" )
  }
  if !v.Bool() {
    t.Error( "Not true" )
  }
}

func TestContext_PushInt( t *testing.T ) {
  ctx := &Context{}

  // Push a value & expect just 1 back
  ctx.PushInt( 42 )
  v, err := ctx.Pop()
  if err != nil {
    t.Error( "Stack underflow" )
  }
  if v.Type() != VAR_INT {
    t.Error( "Not int" )
  }
  if v.Int() != 42 {
    t.Error( "Not 42" )
  }
}

func TestContext_PushFloat( t *testing.T ) {
  ctx := &Context{}

  // Push a value & expect just 1 back
  ctx.PushFloat( 3.14159 )
  v, err := ctx.Pop()
  if err != nil {
    t.Error( "Stack underflow" )
  }
  if v.Type() != VAR_FLOAT {
    t.Error( "Not float" )
  }
  if v.Float() != 3.14159 {
    t.Error( "Not pi" )
  }
}

func TestContext_PushString( t *testing.T ) {
  ctx := &Context{}

  s := "Arthur Dent"

  // Push a value & expect just 1 back
  ctx.PushString( s )
  v, err := ctx.Pop()
  if err != nil {
    t.Error( "Stack underflow" )
  }
  if v.Type() != VAR_STRING {
    t.Error( "Not string" )
  }
  if v.String() != s {
    t.Errorf( "Not %s got %s", s, v.String() )
  }
}

func TestContext_Swap( t *testing.T ) {
  ctx := &Context{}

  // Same as Pop2 test but here we push them as 2,1 then swap which should
  // give us 1,2 which Pop2 returns correctly
  ctx.Push( IntValue(2) )
  ctx.Push( IntValue(1) )

  // Test swap
  ctx.Swap()

  // This should pull them back off, a=1 & b=2
  a, b, err := ctx.Pop2()
  if err != nil {
    t.Error( "Stack underflow" )
  }
  if a.Int() != 1 {
    t.Errorf( "Expected a=1 got %d", a.Int() )
  }
  if b.Int() != 2 {
    t.Errorf( "Expected b=2 got %d", b.Int() )
  }

  // This should fail
  if _, err := ctx.Pop(); err == nil {
    t.Error( "Stack underflow expected" )
  }
}
