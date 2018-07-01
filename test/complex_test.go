package test

import (
  "github.com/peter-mount/calculator/calculator"
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
  "(1 + $a i) + (2 + 1i)",
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
