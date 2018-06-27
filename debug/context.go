package debug

import (
  "github.com/peter-mount/calculator/context"
  "io"
)

func write( w io.Writer, v *context.Value ) {
  s := v.Type() == context.VAR_STRING
  if s {
    io.WriteString( w, "\"")
  }
  io.WriteString( w, v.String() )
  if s {
    io.WriteString( w, "\"")
  }
}

// StackDump writes the current state of the stack to a Writer
func StackDump( w io.Writer, c *context.Context) {
  io.WriteString( w, "[")
  for i, v := range c.Stack() {
    if i>0 {
      io.WriteString( w, ", ")
    }
    write( w, v )
  }
  io.WriteString( w, "]\n")
}

func VarDump( w io.Writer, c *context.Context ) {
  for i, m := range c.Vars() {
    if i>0 {
      io.WriteString( w, ", ")
    }
    io.WriteString( w, "{")
    for k, v := range m {
      io.WriteString( w, "\"")
      io.WriteString( w, k )
      io.WriteString( w, "\"=")
      write( w, v )
    }
    io.WriteString( w, "}")
  }
  io.WriteString( w, "\n")
}
