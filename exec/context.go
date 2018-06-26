package exec

import (
  "errors"
)

// The runtime context used when performing processing
type Context struct {
  stack   []*Value
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
