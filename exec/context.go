package exec

import (
  "errors"
)

// The runtime context used when performing processing
type Context struct {
  stack   []*Value
  vars    []map[string]*Value
}

// Push a value onto the stack
func (c *Context) Push( v *Value ) {
  c.stack = append( c.stack, v )
}

func (c *Context) PushBool( b bool ) {
  c.Push( BoolValue( b ) )
}

func (c *Context) PushInt( i int64 ) {
  c.Push( IntValue( i ) )
}

func (c *Context) PushFloat( f float64 ) {
  c.Push( FloatValue( f ) )
}

func (c *Context) PushString( s string ) {
  c.Push( StringValue( s ) )
}

func (c *Context) Peek() (*Value, error) {
  l := len(c.stack)
  if l == 0 {
    return nil, errors.New( "Stack underflow" )
  }

  return c.stack[l-1], nil
}

// Pop a value from the stack.
// Returns error if the stack is empty
func (c *Context) Pop() (*Value, error) {
  l := len(c.stack)
  if l == 0 {
    return nil, errors.New( "Stack underflow" )
  }

  v := c.stack[l-1]
  c.stack = c.stack[:l-1]
  return v, nil
}

// Pops 2 values from the stack.
// Returns error if the stack is empty
func (c *Context) Pop2() (*Value, *Value, error) {
  // B is first as it's the top value
  b, err := c.Pop()
  if err != nil {
    return nil, nil, err
  }

  // Now a
  a, err := c.Pop()
  if err != nil {
    return nil, nil, err
  }

  // a, b in the order you would expect
  return a, b, nil
}

// Swap swaps the top 2 values on the stack.
// Returns error if the stack doesn't have 2 items to swap.
func (c *Context) Swap() error {
  l := len(c.stack)
  if l < 2 {
    return errors.New( "Stack underflow" )
  }

  a := c.stack[l-2]
  c.stack[l-2] = c.stack[l-1]
  c.stack[l-1] = a
  return nil
}

func (c *Context) GetVar( n string ) *Value {
  for _, v := range c.vars {
    if val, exists := v[n]; exists {
      return val
    }
  }
  return nil
}

func (c *Context) SetVar( n string, val *Value ) {
  // Force a scope to start if we don't have none
  if len(c.vars) == 0 {
    c.StartScope()
  }

  // Find existing entry
  for _, v := range c.vars {
    if _, exists := v[n]; exists {
      v[n] = val
      return
    }
  }

  // Set in current scope
  c.vars[0][n] = val
}

func (c *Context) SetVarBool( n string, val bool ) {
  c.SetVar( n, BoolValue( val ) )
}

func (c *Context) SetVarInt( n string, val int64 ) {
  c.SetVar( n, IntValue( val ) )
}

func (c *Context) SetVarFloat( n string, val float64 ) {
  c.SetVar( n, FloatValue( val ) )
}

func (c *Context) SetVarString( n string, val string ) {
  c.SetVar( n, StringValue( val ) )
}

// Starts a variable scope
func (c *Context) StartScope() {
  c.vars = append( []map[string]*Value{make(map[string]*Value)}, c.vars... )
}

func (c *Context) ResetScope() {
  if len(c.vars) == 0 {
    c.StartScope()
  } else {
    c.vars = []map[string]*Value{c.vars[0]}
  }
}

// Ends the current variable scope. Note this will not remove the first scope
// if present
func (c *Context) EndScope() {
  if len(c.vars) > 1 {
    c.vars = c.vars[1:]
  }
}
