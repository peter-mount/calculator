package exec

import (
  "errors"
  "github.com/peter-mount/calculator/context"
)

type Calculator struct {
  root *context.Node
}

func (c *Calculator) Parse( s string ) error {
  parser := c.Parser()
  err := parser.Parse( s )
  if err != nil {
    return err
  }
  c.root = parser.GetRoot()
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
