package exec

import (
  "fmt"
  "github.com/peter-mount/calculator/lex"
  "strconv"
  "text/scanner"
)

type Parser struct {
  calculator   *Calculator
  lexer         lex.Lexer
  root         *Node
  funcs         FuncMap
  tokenType     rune
  token         string
  precedence    int
}

func (c *Calculator) Parser() *Parser {
  p := &Parser{ calculator: c }
  p.funcs = make( FuncMap )
  return p
}

func (p *Parser) GetRoot() *Node {
  return p.root
}

func (p *Parser) Parse( rule string ) error {
  p.lexer.Parse( rule )
  root, err := p.parse()
  p.root = root
  return err
}

func (p *Parser) AddFuncs( m *FuncMap ) error {
  for k, f := range *m {
    if _, exists := p.funcs[k]; exists {
      return fmt.Errorf( "Token \"%s\" already has a mapping", k )
    }
    p.funcs[k] = f
  }
  return nil
}

func (p *Parser) parse() (*Node,error) {
  n1, err := p.parse_statements()
  return n1, err
}

// Top level for normal arithmetic
func (p *Parser) parse_arithmetic() (*Node,error) {
  n1, err := p.parse_logic()
  return n1, err
}

func (p *Parser) parse_parens() (*Node,error) {

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

func (p *Parser) parse_unary() (*Node,error) {
  var expr *Node
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
      expr = &Node{ token:token.Text(), value: IntValue( iv ) }

    case scanner.Float:
      token = p.lexer.Next()
      fv, err := strconv.ParseFloat( token.Text(), 64 )
      if err != nil {
        return nil, err
      }
      expr = &Node{ token:token.Text(), value: FloatValue( fv )  }

    case scanner.String:
      token = p.lexer.Next()
      expr = &Node{ token:token.Text(), value: StringValue( token.Text()[1:len(token.Text())-1] )  }

    case lex.TOKEN_VARIABLE:
      token = p.lexer.Next()
      expr = &Node{ tokenRune: lex.TOKEN_VARIABLE, token: token.Text(), handler: getVarHandler }

    default:
      err = fmt.Errorf( "Unknown token: \"%s\"", token.Text() )
  }

  return expr, err
}
