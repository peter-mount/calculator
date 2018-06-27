package exec

import (
  "errors"
  "github.com/peter-mount/calculator/lex"
  "math"
)

// parse_math1_handler creates a node for the current token but expects ( ) &
// some arithmetic providing 1 response
func (p *Parser) parse_math1_handler( token *lex.Token, handler NodeHandler ) (*Node,error) {
  token = p.lexer.Next()

  nextToken := p.lexer.Peek()
  if nextToken.Text() != "(" {
    return nil, errors.New( "Expected " + token.Text() + "(arg)" )
  }

  left, err := p.parse_parens()
  if err != nil {
    return nil, err
  }

  expr := &Node{ token:token.Text(), left:left, handler: handler }
  return expr, nil
}

// parse_math parses built in single argument function
func (p *Parser) parse_math() (*Node,error) {
  var expr *Node
  var err error

  token := p.lexer.Peek()
  switch token.Text() {
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

    case "erf":
      expr, err = p.parse_math1_handler( token, erfHandler )
    case "erfc":
      expr, err = p.parse_math1_handler( token, erfcHandler )
    case "erfinv":
      expr, err = p.parse_math1_handler( token, erfinvHandler )
    case "exp":
      expr, err = p.parse_math1_handler( token, expHandler )
    case "exp2":
      expr, err = p.parse_math1_handler( token, exp2Handler )
    case "expm1":
      expr, err = p.parse_math1_handler( token, expm1Handler )
    case "floor":
      expr, err = p.parse_math1_handler( token, floorHandler )

    case "ilogb":
      expr, err = p.parse_math1_handler( token, ilogbHandler )
    case "isnan":
      expr, err = p.parse_math1_handler( token, isNaNHandler )

    case "j0":
      expr, err = p.parse_math1_handler( token, j0Handler )
    case "j1":
      expr, err = p.parse_math1_handler( token, j1Handler )

    case "log":
      expr, err = p.parse_math1_handler( token, logHandler )
    case "log10":
      expr, err = p.parse_math1_handler( token, log10Handler )
    case "log1p":
      expr, err = p.parse_math1_handler( token, log1pHandler )
    case "log2":
      expr, err = p.parse_math1_handler( token, log2Handler )
    case "logb":
      expr, err = p.parse_math1_handler( token, logbHandler )

    case "pow10":
      expr, err = p.parse_math1_handler( token, pow10Handler )

    case "round":
      expr, err = p.parse_math1_handler( token, roundHandler )

    case "roundtoeven":
      expr, err = p.parse_math1_handler( token, round2evenHandler )

    case "sin":
      expr, err = p.parse_math1_handler( token, sinHandler )
    case "sincos":
      expr, err = p.parse_math1_handler( token, sincosHandler )
    case "sinh":
      expr, err = p.parse_math1_handler( token, sinhHandler )
    case "sqrt":
      expr, err = p.parse_math1_handler( token, sqrtHandler )

    case "tan":
      expr, err = p.parse_math1_handler( token, tanHandler )
    case "tanh":
      expr, err = p.parse_math1_handler( token, tanhHandler )
    case "trunc":
      expr, err = p.parse_math1_handler( token, truncHandler )

    case "y0":
      expr, err = p.parse_math1_handler( token, y0Handler )
    case "y1":
      expr, err = p.parse_math1_handler( token, y1Handler )

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
  }
  return errors.New( "Unsupported type" )
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

func erfHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Erf )
}

func erfcHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Erfc )
}

func erfinvHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Erfinv )
}

func expHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Exp )
}

func exp2Handler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Exp2 )
}

func expm1Handler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Expm1 )
}

func floorHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Floor )
}

func ilogbHandler( m *Context, n *Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  if a.IsNumeric() {
    m.PushInt( int64( math.Ilogb( a.Float() ) ) )
    return nil
  }

  return errors.New( "Unsupported type" )
}

func isNaNHandler( m *Context, n *Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  if a.IsNumeric() {
    m.PushBool( math.IsNaN( a.Float() ) )
    return nil
  }

  return errors.New( "Unsupported type" )
}

func j0Handler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.J0 )
}

func j1Handler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.J1 )
}

func logHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Log )
}

func log10Handler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Log10 )
}

func log1pHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Log1p )
}

func log2Handler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Log2 )
}

func logbHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Logb )
}

func pow10Handler( m *Context, n *Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  if a.IsNumeric() {
    m.PushFloat( math.Pow10( int(a.Int()) ) )
    return nil
  }

  return errors.New( "Unsupported type" )
}

func roundHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Round )
}

func round2evenHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.RoundToEven )
}

func sinHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Sin )
}

func sincosHandler( m *Context, n *Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  if a.IsNumeric() {
    s, c := math.Sincos( a.Float() )
    m.PushFloat( s )
    m.PushFloat( c )
    return nil
  }

  return errors.New( "Unsupported type" )
}

func sinhHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Sinh )
}

func sqrtHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Sqrt )
}

func tanHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Tan )
}

func tanhHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Tanh )
}

func truncHandler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Trunc )
}

func y0Handler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Y0 )
}

func y1Handler( m *Context, n *Node ) error {
  return mathInvoke1( m, n, math.Y1 )
}
