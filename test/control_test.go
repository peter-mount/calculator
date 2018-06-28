package test

import (
  "github.com/peter-mount/calculator/calculator"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/debug"
  "os"
  "testing"
)

// Test basic math precedence
func TestIf( t *testing.T ) {

  f, err := os.OpenFile( "/tmp/if.html", os.O_CREATE | os.O_TRUNC|os.O_WRONLY, 0666 )
  if err != nil {
    t.Error( err )
    return
  }
  defer f.Close()

  debug.HtmlTreeStart( f )

  calc := &calculator.Calculator{}

  e := "$a=-1;if 1 { $a=1 ; 42}"

  err = calc.ParseScriptString( e )
  if err != nil {
    t.Error( err )
  }

  debug.HtmlTree( calc.GetRoot(), f, e )

  ctx := &context.Context{}
  err = calc.Execute( ctx )
  if err != nil {
    t.Error( err )
  }
  f.WriteString( "<p><strong>" )
  f.WriteString( e )
  f.WriteString( "</strong> = ")
  debug.StackDump( f, ctx )
  debug.VarDump( f, ctx )
  f.WriteString( "</p> ")

  debug.HtmlTreeEnd( f )
}
