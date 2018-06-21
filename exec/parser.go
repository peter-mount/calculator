package exec

import (
  "fmt"
  "strconv"
  "strings"
  "text/scanner"
)

type Parser struct {
  calculator   *Calculator
  scanner       scanner.Scanner
  root         *Node
  funcs         FuncMap
  tokenType     rune
  token         string
}

func (c *Calculator) Parser() *Parser {
  p := &Parser{ calculator: c }
  p.funcs = make( FuncMap )
  p.AddFuncs( &logicFunctions )
  p.AddFuncs( &arithmeticFunctions )
  p.AddFuncs( &mathFunctions )
  return p
}

func (p *Parser) GetRoot() *Node {
  return p.root
}

func (p *Parser) Parse( rule string ) error {
  fmt.Printf( "parse:\"%s\"\n", rule )
  // Root node is special
  p.root = &Node{ token: "ROOT", handler: rootHandler }

  p.scanner.Init( strings.NewReader( rule ) )
  //p.scanner.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanStrings | scanner.ScanRawStrings | scanner.ScanComments | scanner.SkipComments
  p.scanner.Filename = "filter"

  // Parse the root but always return the root here
  n := p.root
  for p.tokenType != scanner.EOF {
    next, err := p.ParseToken( n )
    if p.tokenType != scanner.EOF {
      if err != nil {
        return err
      }
      n = next
    }
  }

  p.calculator.root = p.root
  return nil
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

// Just invokes the left hand side
func rootHandler( m *Context, n *Node ) error {
  return n.InvokeLhs(m)
}

func (p *Parser) UnknownToken() error {
  return fmt.Errorf( "Unknown token: \"%s\"", p.token )
}

func (p *Parser) Token() string {
  return p.token
}

func (p *Parser) Scan() error {
  p.tokenType = p.scanner.Scan()
  if p.tokenType == scanner.EOF {
    return fmt.Errorf( "EOF" )
  }

  p.token = p.scanner.TokenText()

  // Treat chars as an ident
  if p.tokenType > 32 && p.tokenType < 127 {
    p.tokenType = scanner.Ident
    for p.scanner.Peek() > 32 && p.scanner.Peek() < 127 {
      p.scanner.Scan()
      p.token = p.token + p.scanner.TokenText()
    }
  }

  fmt.Printf( "%2d:%s: %s\n", p.tokenType, p.scanner.Position, p.token )

  return nil
}

// Fail if the next token is not one thats expected
func (p *Parser) Expect( s string ) error {
  err := p.Scan()
  if err == nil && p.Token() != s {
    err = fmt.Errorf( "Unexpected token \"%s\" - expected \"%s\"", p.Token(), s )
  }
  return err
}

// An operation that takes no arguments
func ActionOp( p *Parser, n *Node, h NodeHandler ) (*Node,error) {
  bn := p.New( h )
  return bn, n.Append( bn )
}

// Parse a unary operation, e.g. NOT v
func UnaryOp( p *Parser, n *Node, h NodeHandler ) (*Node,error) {
  bn := p.New( h )

  err := n.Append( bn )
  if err != nil {
    return nil, err
  }

  a, err := p.ParseToken( bn )
  return a, err
}

// Parse a binary operation, e.g. n AND nextNode
func BinaryOp( p *Parser, n *Node, h NodeHandler ) (*Node,error) {
  bn := p.New( h )
  n.Replace( bn )

  a, err := p.ParseToken( bn )
  return a, err
}

// Create a new node for a handler
func (p *Parser) New( f NodeHandler ) *Node {
  return &Node{ token: p.Token(), handler: f }
}

func (p *Parser) NewConstant( v *Value ) *Node {
  return &Node{ token: p.Token(), value: v }
}

// This is the main parser function - in a separate file for maintainability
func (p *Parser) ParseToken( n *Node ) (*Node,error) {

  err := p.Scan()
  if err != nil {
    return nil, err
  }

  switch p.tokenType {
    case scanner.Ident:
      fme, ok := p.funcs[ p.Token() ]
      if ok {
        bn, err := fme.ParserDefinition( p, n, fme.NodeHandler )
        return bn, err
      }
    case scanner.Int:
      iv, err := strconv.ParseInt( p.Token(), 10, 64 )
      if err != nil {
        return nil, err
      }
      bn := p.NewConstant( IntValue( iv ) )
      return bn, n.Append( bn )

    case scanner.Float:
      fv, err := strconv.ParseFloat( p.Token(), 64 )
      if err != nil {
        return nil, err
      }
      bn := p.NewConstant( FloatValue( fv ) )
      err = n.Append( bn )
      return bn, err

    case scanner.String:
      bn := p.NewConstant( StringValue( p.Token() ) )
      err = n.Append( bn )
      return bn, err

    default:
      return nil, p.UnknownToken()
  }

  return nil, p.UnknownToken()
}
