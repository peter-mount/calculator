package exec

import (
  //"fmt"
  "testing"
)

// Test basic math precedence
func TestLexer( t *testing.T ) {

  lex := &Lexer{}

  expected := []string{ "1","2","3","+","-","10","3.14159","(","2","*","2",")","+","(","3","*","3",")" }
  lex.Parse( "1 2 3 + - 10 3.14159 (2*2)+(3*3)")

  //for i, t := range lex.tokens { fmt.Printf( "%d:%s\n", i, t.text ) }

  // len( lex.tokens ) is 6 as we also have eof
  if lex.last != 18 {
    t.Errorf( "Not enough tokens, expected 18 got %d", lex.last )
  }

  for i, s := range expected {
    token := lex.Peek()

    if token.text != s {
      t.Errorf( "Peek: Token %d expected \"%s\" got \"%s\"", i, s, token.text )
    }

    token = lex.Next()
    if token.text != s {
      t.Errorf( "Next: Token %d expected \"%s\" got \"%s\"", i, s, token.text )
    }
  }
}
