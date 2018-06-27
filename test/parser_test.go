package test

import (
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/debug"
  "github.com/peter-mount/calculator/exec"
  "testing"
  "io"
  "os"
)

var testParser_eq []string = []string{
  "1 + 2",          // 1 + (2*3) = 7
  //"1 + -2",          // 1 + (2*3) = 7
  //"1 +-2",          // 1 + (2*3) = 7
  "2 * 3",          // 1 + (2*3) = 7
  "1 + 2 * 3",      // 1 + (2*3) = 7
  "1 + 2 * 3 + 2",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 == 9",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 == 3 * 3",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 == (3 * 2)+3",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 == (3 * 2 - 1)+3",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 == (3 * 2 - 1)/(3*21)",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 == (3 * 2 - 1)/3*21",  // 1 + (2*3) + 2 == 9
}

func TestParser_parse( t *testing.T ) {

  f, err := os.OpenFile( "/tmp/parser.html", os.O_CREATE | os.O_TRUNC|os.O_WRONLY, 0666 )
  if err != nil {
    t.Error( err )
  } else {
    defer f.Close()

    debug.HtmlTreeStart( f )

    calc := &exec.Calculator{}

    for _, e := range testParser_eq {
      parser := calc.Parser()
      err = parser.Parse( e )
      if err != nil {
        t.Error( err )
      } else if parser.GetRoot() == nil {
        io.WriteString( f, "*** nil root ***" )
      } else {
        debug.HtmlTree( parser.GetRoot(), f, e )
      }
    }

    debug.HtmlTreeEnd( f )
  }
}

func TestParser_execute( t *testing.T ) {
  results := []*context.Value {
    context.IntValue( 3 ),
    context.IntValue( 6 ),
    context.IntValue( 7 ),
    context.IntValue( 9 ),
    context.BoolValue( true ),
    context.BoolValue( true ),
    context.BoolValue( true ),
    context.BoolValue( false ),
    context.BoolValue( false ),
    context.BoolValue( false ),
  }

  calc := &exec.Calculator{}
  ctx := &context.Context{}

  for i, eq := range testParser_eq {

    err := calc.Parse( eq )
    if err != nil {
      t.Error( err )
    }

    err = calc.Execute( ctx )
    if err != nil {
      t.Error( err )
    }

    result, err := ctx.Pop()
    if err != nil {
      t.Error( err )
    } else if !result.Same( results[i] ) {
      t.Errorf( "Unexpected result %d, expected %v got %v", i, results[i], result )
    }
  }
}
