package exec

import (
  "fmt"
  "strings"
  "text/scanner"
)

type Parser struct {
  calculator   *Calculator
  scanner       scanner.Scanner
  root         *Node
  funcs         FuncMap
}

func (c *Calculator) Parser() *Parser {
  p := &Parser{ calculator: c }
  p.funcs = make( FuncMap )
  return p
}

func (p *Parser) Parse( rule string ) error {
  // Root node is special
  p.root = &Node{ token: "ROOT", handler: rootHandler }

  p.scanner.Init( strings.NewReader( rule ) )
  p.scanner.Filename = "filter"

  // Parse the root but always return the root here
  err := p.ParseToken( p.root )
  if err != nil {
    return err
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
  return fmt.Errorf( "Unknown token: \"%s\"", p.scanner.TokenText() )
}

func (p *Parser) Token() string {
  return p.scanner.TokenText()
}

func (p *Parser) Scan() error {
  if p.scanner.Scan() == scanner.EOF {
    return fmt.Errorf( "EOF" )
  }

  fmt.Printf( "%s: %s\n", p.scanner.Position, p.scanner.TokenText() )

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
func ActionOp( p *Parser, n *Node, h NodeHandler ) error {
  return n.Append( p.New( h ) )
}

// Parse a unary operation, e.g. NOT v
func UnaryOp( p *Parser, n *Node, h NodeHandler ) error {
  bn := p.New( h )
  err := n.Append( bn )
  if err != nil {
    return err
  }
  return p.ParseToken( bn )
}

// Parse a binary operation, e.g. n AND nextNode
func BinaryOp( p *Parser, n *Node, h NodeHandler ) error {
  bn := p.New( h )
  n.Replace( bn )
  return p.ParseToken( bn )
}

// Parse a binary operation with separator, e.g. BETWEEN a AND b
// For an operator requiring 2 params use "" for s. e.g. ATAN2 a b
func BinaryOpSep( p *Parser, n *Node, h NodeHandler, s string ) error {
  bn := p.New( h )

  // this is just appended to
  err := n.Append( bn )
  if err != nil {
    return err
  }

  // lhs
  err = p.ParseToken( bn )
  if err != nil {
    return err
  }

  // Required separator, "" for not required
  if s != "" {
    err = p.Expect( s )
    if err != nil {
      return err
    }
  }

  // rhs
  return p.ParseToken( bn )
}

// Create a new node for a handler
func (p *Parser) New( f NodeHandler ) *Node {
  return &Node{ token: p.Token(), handler: f }
}

// This is the main parser function - in a separate file for maintainability
func (p *Parser) ParseToken( n *Node ) error {

  err := p.Scan()
  if err != nil {
    return err
  }

  if fme, ok := p.funcs[ p.Token() ]; ok {
    err = fme.ParserDefinition( p, n, fme.NodeHandler )
  } else {
    // Check for constant & if not then fail
    /*
    v, err := p.tokenInt()
    if err == nil {
      err = n.append( p.newConstant( v ) )
      if err != nil {
        return err
      }
    } else {
      return p.UnknownToken()
    }
    */
    return p.UnknownToken()
  }

  return err
}
