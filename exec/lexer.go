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

const (
  TOKEN_VARIABLE = -(iota+100)
)

// Returns true if a rune should be matched as an ident
func IsIdent( r rune ) bool {
  return r=='=' ||
    r=='^' || r=='&' ||
    r=='%' || r=='!' ||
    r==':' || r==';'
}

// IsVariableStart true if the rune is valid for the first char of a variable name
func IsVariableStart( r rune ) bool {
  return (r>='a' && r<='z') || (r>='A' && r<='Z')
}
// IsVariableSuccessor true if the rune is valid for the successive chars in a variable name
func IsVariableSuccessor( r rune ) bool {
  return IsVariableStart(r) || r=='_' || (r>='0' && r<='9')
}

func IsPlusMinus( r rune ) bool {
  return r=='+' || r=='-'
}

func IsDigit( r rune ) bool {
  return r>='0' && r<='9'
}

func (l *Lexer) scan() *Token {
  return &Token{ token: l.scanner.Scan() }
}

func (l *Lexer) scanNext() string {
  l.scanner.Scan()
  return l.scanner.TokenText()
}

func (l *Lexer) scanWhile(f func(rune)bool) string {
  var s string
  for f( l.scanner.Peek() ) {
    s = s + l.scanNext()
  }
  return s
}

func (l *Lexer) Parse( rule string ) {
  l.scanner.Init( strings.NewReader( rule ) )

  var token *Token
  for token == nil || token.token != scanner.EOF {
    token = l.scan()

    if token.token != scanner.EOF {
      token.text = l.scanner.TokenText()

      // Treat chars as an ident
      if IsIdent( token.token ) {
        token.token = scanner.Ident
        token.text = token.text + l.scanWhile( IsIdent )
      } else if token.token == '$' {
        // $AAA as a variable name so parse it without $ in the text
        token.token = TOKEN_VARIABLE
        token.text = l.scanWhile( IsVariableStart ) + l.scanWhile( IsVariableStart )
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
