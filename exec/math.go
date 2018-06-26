package exec

import (
  "errors"
  "math"
)

// parse_math1_handler creates a node for the current token but expects ( ) &
// some arithmetic providing 1 response
func (p *Parser) parse_math1_handler( token *Token, handler NodeHandler ) (*Node,error) {
  token = p.lexer.Next()

  nextToken := p.lexer.Peek()
  if nextToken.text != "(" {
    return nil, errors.New( "Expected " + token.text + "(arg)" )
  }

  left, err := p.parse_parens()
  if err != nil {
    return nil, err
  }

  expr := &Node{ token:token.text, left:left, handler: handler }
  return expr, nil
}

// parse_math parses the internal math functions
func (p *Parser) parse_math() (*Node,error) {
  var expr *Node
  var err error

  token := p.lexer.Peek()
  switch token.text {
    case "abs":
      expr, err = p.parse_math1_handler( token, absHandler )
    case "acos":
      expr, err = p.parse_math1_handler( token, acosHandler )
    case "acosh":
        expr, err = p.parse_math1_handler( token, acoshHandler )
    case "asin":
        expr, err = p.parse_math1_handler( token, asinHandler )
    case "asinh":
        expr, err = p.parse_math1_handler( token, asinhHandler )
    case "atan":
        expr, err = p.parse_math1_handler( token, atanHandler )
    //"atan2":      FuncMapEntry{ atan2Handler,             BinaryOp  },
    case "atanh":
        expr, err = p.parse_math1_handler( token, atanhHandler )
    case "cbrt":
        expr, err = p.parse_math1_handler( token, cbrtHandler )
    case "ceil":
        expr, err = p.parse_math1_handler( token, ceilHandler )
    case "cos":
        expr, err = p.parse_math1_handler( token, cosHandler )
    case "cosh":
        expr, err = p.parse_math1_handler( token, coshHandler )
    default:
      expr, err = p.parse_constants()
  }

  return expr, err
}

// Handles math functions that take 1 parameter
func mathInvoke1( m *Context, n *Node, f func(float64) float64 ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  if a.IsNumeric() {
    m.PushFloat( f( a.Float() ) )
    return nil
  } else {
    return errors.New( "Unsupported type" )
  }
}

// Handles math functions that take 2 parameters
func mathInvoke2( m *Context, n *Node, f func(float64,float64) float64 ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  if a.IsNumeric() && b.IsNumeric() {
    m.PushFloat( f( a.Float(), b.Float() ) )
    return nil
  } else {
    return errors.New( "Unsupported type" )
  }
}

func absHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Abs )
}

func acosHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Acos )
}

func acoshHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Acosh )
}

func asinHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Asin )
}

func asinhHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Asinh )
}

func atanHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Atan )
}

func atan2Handler( m *Context, n *Node ) error {
  return mathInvoke2( m, n, math.Atan2 )
}

func atanhHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Atanh )
}

func cbrtHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Cbrt )
}

func ceilHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Ceil )
}

func cosHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Cos )
}

func coshHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Cosh )
}
