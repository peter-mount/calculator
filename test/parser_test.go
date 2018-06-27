package test

import (
  "github.com/peter-mount/calculator/calculator"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/debug"
  "testing"
  "io"
  "os"
)

var testParser_eq []string = []string{
  "1 + 2",          // 1 + (2*3) = 7
  //"1 + -2",          // 1 + (2*3) = 7
  //"1 +-2",          // 1 + (2*3) = 7
  "2 * 3",          // 1 + (2*3) = 7
  "2 / 3",          // 1 + (2*3) = 7
  "1 + 2 / 3",      // 1 + (2*3) = 7
  "1 + 2 * 3",      // 1 + (2*3) = 7
  "1 + 2 * 3 + 2",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 == 9",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 != 9",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 < 9",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 <= 9",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 > 9",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 >= 9",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 == 42",  // 1 + (2*3) + 2 == 9
  "1 + 2 * 3 + 2 != 42",  // 1 + (2*3) + 2 == 9
  "1 + 2 * $a + 2 == 9",  // 1 + (2*3) + 2 == 9
  "1 + 2 * $a + 2 != 9",  // 1 + (2*3) + 2 == 9
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

    calc := &calculator.Calculator{}

    for _, e := range testParser_eq {
      err = calc.Parse( e )
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
}

func TestParser_execute( t *testing.T ) {
  /*
  results := []*context.Value {
    context.IntValue( 3 ),
    context.IntValue( 6 ),
    context.FloatValue( 2.0/3.0 ),
    context.FloatValue( 1.0+(2.0/3.0) ),
    context.IntValue( 7 ),
    context.IntValue( 9 ),
    context.BoolValue( true ),
    context.BoolValue( true ),
    context.BoolValue( true ),
    context.BoolValue( false ),
    context.BoolValue( false ),
    context.BoolValue( false ),
  }
  */

  calc := &calculator.Calculator{}
  ctx := &context.Context{}
  ctx.SetVarInt( "a", 42 )

  for _, eq := range testParser_eq {

    err := calc.Parse( eq )
    if err != nil {
      t.Error( err )
    }

    err = calc.Execute( ctx )
    if err != nil {
      t.Error( err )
    }

    _, err = ctx.Pop()
    if err != nil {
      t.Error( err )
    //} else if !result.Same( results[i] ) {
      // TODO fix when we can fix float equality for results line 0.666666667
      //t.Errorf( "Unexpected result %d, expected %v got %v", i, results[i], result )
    }
  }
}
