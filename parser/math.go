package parser

import (
  "errors"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/exec"
  "github.com/peter-mount/calculator/lex"
)

// parse_math1_handler creates a node for the current token but expects ( ) &
// some arithmetic providing 1 response
func (p *Parser) parse_math1_handler( token *lex.Token, f UnaryFunction ) (*context.Node,error) {
  token = p.lexer.Next()

  nextToken := p.lexer.Peek()
  if nextToken.Text() != "(" {
    return nil, errors.New( "Expected " + token.Text() + "(arg)" )
  }

  left, err := p.parse_parens()
  if err != nil {
    return nil, err
  }

  expr, err := OptimizeUnaryFunction( token, left, f )
  return expr, nil
}

// parse_math parses built in single argument function
func (p *Parser) parse_math() (*context.Node,error) {
  var expr *context.Node
  var err error

  token := p.lexer.Peek()
  switch token.Text() {
    case "abs":
      expr, err = p.parse_math1_handler( token, exec.Abs )
    case "acos":
      expr, err = p.parse_math1_handler( token, exec.Acos )
    case "acosh":
      expr, err = p.parse_math1_handler( token, exec.Acosh )
    case "asin":
      expr, err = p.parse_math1_handler( token, exec.Asin )
    case "asinh":
      expr, err = p.parse_math1_handler( token, exec.Asinh )
    case "atan":
      expr, err = p.parse_math1_handler( token, exec.Atan )
    case "atanh":
      expr, err = p.parse_math1_handler( token, exec.Atanh )
    case "cbrt":
      expr, err = p.parse_math1_handler( token, exec.Cbrt )
    case "ceil":
      expr, err = p.parse_math1_handler( token, exec.Ceil )
    case "cos":
      expr, err = p.parse_math1_handler( token, exec.Cos )
    case "cosh":
      expr, err = p.parse_math1_handler( token, exec.Cosh )

    case "erf":
      expr, err = p.parse_math1_handler( token, exec.Erf )
    case "erfc":
      expr, err = p.parse_math1_handler( token, exec.Erfc )
    case "erfcinv":
      expr, err = p.parse_math1_handler( token, exec.Erfcinv )
    case "erfinv":
      expr, err = p.parse_math1_handler( token, exec.Erfinv )
    case "exp":
      expr, err = p.parse_math1_handler( token, exec.Exp )
    case "exp2":
      expr, err = p.parse_math1_handler( token, exec.Exp2 )
    case "expm1":
      expr, err = p.parse_math1_handler( token, exec.Expm1 )
    case "floor":
      expr, err = p.parse_math1_handler( token, exec.Floor )

    case "gamma":
      expr, err = p.parse_math1_handler( token, exec.Gamma )

    case "ilogb":
      expr, err = p.parse_math1_handler( token, exec.Ilogb )
    case "isnan":
      expr, err = p.parse_math1_handler( token, exec.IsNaN )

    case "j0":
      expr, err = p.parse_math1_handler( token, exec.J0 )
    case "j1":
      expr, err = p.parse_math1_handler( token, exec.J1 )

    case "log":
      expr, err = p.parse_math1_handler( token, exec.Log )
    case "log10":
      expr, err = p.parse_math1_handler( token, exec.Log10 )
    case "log1p":
      expr, err = p.parse_math1_handler( token, exec.Log1p )
    case "log2":
      expr, err = p.parse_math1_handler( token, exec.Log2 )
    case "logb":
      expr, err = p.parse_math1_handler( token, exec.Logb )

    case "pow10":
      expr, err = p.parse_math1_handler( token, exec.Pow10 )

    case "round":
      expr, err = p.parse_math1_handler( token, exec.Round )

    case "roundtoeven":
      expr, err = p.parse_math1_handler( token, exec.Round2even )

    case "sin":
      expr, err = p.parse_math1_handler( token, exec.Sin )
    case "sinh":
      expr, err = p.parse_math1_handler( token, exec.Sinh )
    case "sqrt":
      expr, err = p.parse_math1_handler( token, exec.Sqrt )

    case "tan":
      expr, err = p.parse_math1_handler( token, exec.Tan )
    case "tanh":
      expr, err = p.parse_math1_handler( token, exec.Tanh )
    case "trunc":
      expr, err = p.parse_math1_handler( token, exec.Trunc )

    case "y0":
      expr, err = p.parse_math1_handler( token, exec.Y0 )
    case "y1":
      expr, err = p.parse_math1_handler( token, exec.Y1 )

    default:
      expr, err = p.parse_constants()
  }

  return expr, err
}
