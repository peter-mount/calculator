package calculator

import (
  "errors"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/lex"
  "github.com/peter-mount/calculator/parser"
)

type Calculator struct {
  root *context.Node
}

func (c *Calculator) Parse( s string ) error {

  lexer := &lex.Lexer{}
  lexer.Parse( s )

  parse := parser.NewParser( lexer )

  root, err := parse.ParseStatements()
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