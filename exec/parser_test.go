package exec

import (
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

  f, err := os.OpenFile( "/tmp/out.html", os.O_CREATE | os.O_TRUNC|os.O_WRONLY, 0666 )
  if err != nil {
    t.Error( err )
  } else {
    defer f.Close()

    HtmlTreeStart( f )

    calc := &Calculator{}

    for _, e := range testParser_eq {
      parser := calc.Parser()
      // Uncomment to see debugging
      //parser.Debug = true

      err = parser.Parse( e )
      if err != nil {
        t.Error( err )
      } else if parser.GetRoot() == nil {
        io.WriteString( f, "*** nil root ***" )
      } else {
        HtmlTree( parser.GetRoot(), f, e )
      }
    }

    HtmlTreeEnd( f )
  }
}

func TestParser_execute( t *testing.T ) {
  results := []*Value {
    IntValue( 3 ),
    IntValue( 6 ),
    IntValue( 7 ),
    IntValue( 9 ),
    BoolValue( true ),
    BoolValue( true ),
    BoolValue( true ),
    BoolValue( false ),
    BoolValue( false ),
    BoolValue( false ),
  }

  calc := &Calculator{}
  ctx := &Context{}

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
