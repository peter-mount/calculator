package parser

import (
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/lex"
)

// Function that performs an operation on two Values
type BinaryFunction func(*context.Value,*context.Value)(*context.Value,error)

// OptimizeOperation will if both left and right nodes are constants return
// a constant node with the result of some function.
// If either are not constant then a new node will be created with the supplied handler
// attached.
func OptimizeOperation( token *lex.Token, left *context.Node, right *context.Node, f BinaryFunction ) (*context.Node,error) {
  if left != nil && right != nil && left.IsConstant() && right.IsConstant() {
    c, err := f( left.Value(), right.Value() )
    if err != nil {
      return nil, err
    }
    return context.NewConstant( token, c ), nil
  } else {
    return context.NewNode(
      token,
      func( m *context.Context, n *context.Node ) error {
        err := n.Invoke2(m)
        if err != nil {
          return err
        }

        a, b, err := m.Pop2()
        if err != nil {
          return err
        }

        c, err := f( a, b )
        if err == nil {
          m.Push( c )
        }
        return err
      },
      left,
      right ), nil
  }
}
