package exec

import (
  "testing"
  "io"
  "os"
)

func TestParser_math( t *testing.T ) {

  testdata := []string{
    "1 + abs(1)",
    "1 + abs(-1)",
    "pi",
    "2 * pi",
    "-phi",
    "1+-0.5",
    "1 + \"0.5\"",
    "$a = 6*7",
    "$a = 6*7 $a/2",
    "$a = 6*7;$a/2;$a*2",
    "{$a = 6*7;$a/2;$a*2}",
  }

  f, err := os.OpenFile( "/tmp/math.html", os.O_CREATE | os.O_TRUNC|os.O_WRONLY, 0666 )
  if err != nil {
    t.Error( err )
    return
  }
  defer f.Close()

  HtmlTreeStart( f )

  calc := &Calculator{}

  for _, e := range testdata {
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

  for _, eq := range testdata {
    err := calc.Parse( eq )
    if err != nil {
      t.Error( err )
    } else {
      ctx := &Context{}
      //ctx.SetVarInt( "a", 42 )

      err = calc.Execute( ctx )
      if err != nil {
        t.Error( err )
      }
      f.WriteString( "<p><strong>" )
      f.WriteString( eq )
      f.WriteString( "</strong> = ")
      ctx.StackDump( f )
      ctx.VarDump( f )
      f.WriteString( "</p> ")
    }
  }
}
