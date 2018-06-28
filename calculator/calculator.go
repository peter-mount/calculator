package calculator

import (
  "errors"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/lex"
  "github.com/peter-mount/calculator/parser"
  "io"
  "strings"
)

type Calculator struct {
  root *context.Node
}

func (c *Calculator) ParseExpressionString( s string ) error {
  return c.ParseExpression( strings.NewReader( s ) )
}

func (c *Calculator) ParseExpression( r io.Reader ) error {
  return c.parse( r, c.parseExpression )
}

func (c *Calculator) parseExpression(lexer *lex.Lexer) (*context.Node,error) {
  root, err := parser.NewParser( lexer ).ParseExpression()
  return root, err
}

func (c *Calculator) ParseScriptString( s string ) error {
  return c.ParseScript( strings.NewReader( s ) )
}

func (c *Calculator) ParseScript( r io.Reader ) error {
  return c.parse( r, c.parseStatements )
}

func (c *Calculator) parseStatements(lexer *lex.Lexer) (*context.Node,error) {
  root, err := parser.NewParser( lexer ).ParseStatements()
  return root, err
}

func (c *Calculator) parse( r io.Reader, f func(*lex.Lexer)(*context.Node,error) ) error {
  lexer := &lex.Lexer{}
  lexer.Parse( r )

  root, err := f( lexer )
  if err != nil {
    return err
  }
  c.root = root

  if c.root == nil {
    return errors.New( "Nothing generated from parser" )
  }

  return nil
}

func (c *Calculator) Execute( ctx *context.Context ) error {
  if c.root == nil {
    return errors.New( "Parse() required first" )
  }
  if ctx == nil {
    return errors.New( "No Context" )
  }

  ctx.ResetScope()

  return c.root.Invoke( ctx )
}

func (c *Calculator) GetRoot() *context.Node {
  return c.root
}
