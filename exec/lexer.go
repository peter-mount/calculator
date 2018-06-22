package exec

import (
  "strings"
  "text/scanner"
)

type Lexer struct {
  tokens     []*Token
  pos           int
  last          int
  scanner       scanner.Scanner
}

type Token struct {
  token    rune
  text    string
}

func (l *Lexer) Parse( rule string ) {
  l.scanner.Init( strings.NewReader( rule ) )

  var token *Token
  for token == nil || token.token != scanner.EOF {
    token = &Token{ token: l.scanner.Scan() }

    if token.token != scanner.EOF {
      token.text = l.scanner.TokenText()

      // Treat chars as an ident
      if token.token > 32 && token.token < 127 {
        token.token = scanner.Ident
        for l.scanner.Peek() > 32 && l.scanner.Peek() < 127 {
          l.scanner.Scan()
          token.text = token.text + l.scanner.TokenText()
        }
      }

    } else {
      token.text = "EOF"
    }

    l.tokens = append( l.tokens, token )
  }

  l.last = len(l.tokens) -1
}

// Get the current token and move forward one place
func (l *Lexer) Next() *Token {
  token := l.tokens[l.pos]
  if l.pos <= l.last {
    l.pos++;
  }
  return token
}

// Get the current token but do not move
func (l *Lexer) Peek() *Token {
  return l.tokens[l.pos]
}
