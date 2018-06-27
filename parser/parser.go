package parser

import (
  "fmt"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/exec"
  "github.com/peter-mount/calculator/lex"
  "strconv"
  "text/scanner"
)

type Parser struct {
  lexer        *lex.Lexer
  //funcs         FuncMap
  tokenType     rune
  token         string
  precedence    int
}

func NewParser( l *lex.Lexer ) *Parser {
  p := &Parser{ lexer: l }
  //p.funcs = make( FuncMap )
  return p
}

func (p *Parser) ParseStatements() (*context.Node,error) {
  n1, err := p.parse_statements()
  return n1, err
}

// Top level for normal arithmetic
func (p *Parser) parse_arithmetic() (*context.Node,error) {
  n1, err := p.parse_logic()
  return n1, err
}

func (p *Parser) parse_parens() (*context.Node,error) {

  token := p.lexer.Peek()
  if token.Text() == "(" {
    p.lexer.Next()

    expr, err := p.parse_arithmetic()
    if err != nil {
      return nil, err
    }

    token = p.lexer.Next()
    if token.Text() != ")" {
      return nil, fmt.Errorf( "Expecting )" )
    }
    return expr, nil
  }

  expr, err := p.parse_math()
  return expr, err
}

func (p *Parser) parse_unary() (*context.Node,error) {
  var expr *context.Node
  var err error

  token := p.lexer.Peek()

  switch token.Token() {
    case scanner.Ident:
      token = p.lexer.Next()
      /*

      fme, ok := p.funcs[ token ]
      if ok {
        n1, err := fme.ParserDefinition( p, n, fme.NodeHandler )
        return n1, err
      }
      */
      err = fmt.Errorf( "XXX Unknown token: \"%s\"", token.Text() )
    case scanner.Int:
      token = p.lexer.Next()
      iv, err := strconv.ParseInt( token.Text(), 10, 64 )
      if err != nil {
        return nil, err
      }
      expr = context.NewConstant( token, context.IntValue( iv ) )

    case scanner.Float:
      token = p.lexer.Next()
      fv, err := strconv.ParseFloat( token.Text(), 64 )
      if err != nil {
        return nil, err
      }
      expr = context.NewConstant( token, context.FloatValue( fv ) )

    case scanner.String:
      token = p.lexer.Next()
      expr = context.NewConstant( token, context.StringValue( token.Text()[1:len(token.Text())-1] ) )

    case lex.TOKEN_VARIABLE:
      token = p.lexer.Next()
      expr = context.NewNode( token, exec.GetVarHandler, nil, nil )

    default:
      err = fmt.Errorf( "Unknown token: \"%s\"", token.Text() )
  }

  return expr, err
}
