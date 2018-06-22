package exec

import (
  "fmt"
  "strconv"
  "text/scanner"
)

type Parser struct {
  calculator   *Calculator
  lexer        Lexer
  root         *Node
  funcs         FuncMap
  tokenType     rune
  token         string
  Debug         bool
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
  n1, err := p.parse_additive()
  return n1, err
}

func (p *Parser) parse_additive() (*Node,error) {

  expr, err := p.parse_multiplicative()
  if err != nil {
    return nil, err
  }

  token := p.lexer.Peek()
  for token.text == "+" || token.text == "-" {
    token = p.lexer.Next()

    right, err := p.parse_multiplicative()
    if err != nil {
      return nil, err
    }

    switch token.text {
      case "+":
        expr = &Node{ token:token.text, left:expr, right: right, handler: addHandler }
      case "-":
        expr = &Node{ token:token.text, left:expr, right: right, handler: subHandler }
    }

    token = p.lexer.Peek()
  }

  return expr, err
}

func (p *Parser) parse_multiplicative() (*Node,error) {

  expr, err := p.parse_unary()
  if err != nil {
    return nil, err
  }

  token := p.lexer.Peek()
  for token.text == "*" || token.text == "/" {
    token = p.lexer.Next()

    right, err := p.parse_unary()
    if err != nil {
      return nil, err
    }

    switch token.text {
      case "*":
        expr = &Node{ token:token.text, left:expr, right: right, handler: multHandler }
      case "/":
        expr = &Node{ token:token.text, left:expr, right: right, handler: divHandler }
    }

    token = p.lexer.Peek()
  }

  return expr, err
}

func (p *Parser) parse_unary() (*Node,error) {
  var expr *Node
  var err error

  token := p.lexer.Peek()

  switch token.token {
    /*
    case scanner.Ident:
      fme, ok := p.funcs[ p.Token() ]
      if ok {
        n1, err := fme.ParserDefinition( p, n, fme.NodeHandler )
        return n1, err
      }
      */
    case scanner.Int:
      token = p.lexer.Next()
      iv, err := strconv.ParseInt( token.text, 10, 64 )
      if err != nil {
        return nil, err
      }
      expr = &Node{ token:token.text, value: IntValue( iv ) }

    case scanner.Float:
      token = p.lexer.Next()
      fv, err := strconv.ParseFloat( token.text, 64 )
      if err != nil {
        return nil, err
      }
      expr = &Node{ token:token.text, value: FloatValue( fv )  }

    case scanner.String:
      token = p.lexer.Next()
      expr = &Node{ token:token.text, value: StringValue( token.text )  }

    default:
      err = fmt.Errorf( "Unknown token: \"%s\"", token.text )
  }

  return expr, err
}
