package test

import (
  "fmt"
  "github.com/peter-mount/calculator/calculator"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/debug"
  "io"
  "os"
  "testing"
)

var testParser_complex []string = []string{
  "1 + 1i",
  "(1 + 1i)",
  "(1 + 1i) + (1 + 1i)",
  "(1 + 1i) - (1 + 1i)",
  "(1 + 1i) * (1 + 1i)",
  "(1 + 1i) / (1 + 1i)",
  "(1 + 1i) == (1 + 1i)",
}

// Test basic math precedence
func TestComplex( t *testing.T ) {


  f, err := os.OpenFile( "/tmp/complex.html", os.O_CREATE | os.O_TRUNC|os.O_WRONLY, 0666 )
  if err != nil {
    t.Error( err )
    return
  }
  defer f.Close()

  debug.HtmlTreeStart( f )

  c := complex( 1, 2 )
  vc := context.ComplexValue( c )
  f.WriteString( fmt.Sprintf( "<p>Complex %v %v %s</p>", c, vc.Complex(), vc.String() ) )

  calc := &calculator.Calculator{}

  for _, e := range testParser_complex {
    err = calc.ParseScriptString( e )
    if err != nil {
      t.Error( err )
    } else if calc.GetRoot() == nil {
      io.WriteString( f, "*** nil root ***" )
    } else {
      debug.HtmlTree( calc.GetRoot(), f, e )
    }
  }

  debug.HtmlTreeEnd( f )

}
