package exec

import (
  "testing"
  "io"
  "os"
)

// Test basic math precedence
func TestParser_1( t *testing.T ) {

  calc := &Calculator{}

  eq := []string{
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

  f, err := os.OpenFile( "/tmp/out.html", os.O_CREATE | os.O_TRUNC|os.O_WRONLY, 0666 )
  if err != nil {
    t.Error( err )
  } else {
    defer f.Close()
    HtmlTreeStart( f )

    for _, e := range eq {
      parser := calc.Parser()
      parser.Debug = true

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
