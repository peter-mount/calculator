package test

import (
  "github.com/peter-mount/calculator/lex"
  "testing"
)

// Test basic math precedence
func TestLexer( t *testing.T ) {

  lex := &lex.Lexer{}

  expected := []string{ "1","2","3","+","-","10","3.14159","(","2","*","2",")","+","(","3","*","3",")" }
  lex.Parse( "1 2 3 + - 10 3.14159 (2*2)+(3*3)")

  for i, s := range expected {
    token := lex.Peek()

    if token.Text() != s {
      t.Errorf( "Peek: Token %d expected \"%s\" got \"%s\"", i, s, token.Text() )
    }

    token = lex.Next()
    if token.Text() != s {
      t.Errorf( "Next: Token %d expected \"%s\" got \"%s\"", i, s, token.Text() )
    }
  }
}
