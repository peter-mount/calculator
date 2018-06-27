package parser

import (
  "errors"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/exec"
  "github.com/peter-mount/calculator/lex"
)

// parse_math1_handler creates a node for the current token but expects ( ) &
// some arithmetic providing 1 response
func (p *Parser) parse_math1_handler( token *lex.Token, handler context.NodeHandler ) (*context.Node,error) {
  token = p.lexer.Next()

  nextToken := p.lexer.Peek()
  if nextToken.Text() != "(" {
    return nil, errors.New( "Expected " + token.Text() + "(arg)" )
  }

  left, err := p.parse_parens()
  if err != nil {
    return nil, err
  }

  expr := context.NewNode( token, handler, left, nil )
  return expr, nil
}

// parse_math parses built in single argument function
func (p *Parser) parse_math() (*context.Node,error) {
  var expr *context.Node
  var err error

  token := p.lexer.Peek()
  switch token.Text() {
    case "abs":
      expr, err = p.parse_math1_handler( token, exec.AbsHandler )
    case "acos":
      expr, err = p.parse_math1_handler( token, exec.AcosHandler )
    case "acosh":
      expr, err = p.parse_math1_handler( token, exec.AcoshHandler )
    case "asin":
      expr, err = p.parse_math1_handler( token, exec.AsinHandler )
    case "asinh":
      expr, err = p.parse_math1_handler( token, exec.AsinhHandler )
    case "atan":
      expr, err = p.parse_math1_handler( token, exec.AtanHandler )
    case "atanh":
      expr, err = p.parse_math1_handler( token, exec.AtanhHandler )
    case "cbrt":
      expr, err = p.parse_math1_handler( token, exec.CbrtHandler )
    case "ceil":
      expr, err = p.parse_math1_handler( token, exec.CeilHandler )
    case "cos":
      expr, err = p.parse_math1_handler( token, exec.CosHandler )
    case "cosh":
      expr, err = p.parse_math1_handler( token, exec.CoshHandler )

    case "erf":
      expr, err = p.parse_math1_handler( token, exec.ErfHandler )
    case "erfc":
      expr, err = p.parse_math1_handler( token, exec.ErfcHandler )
    case "erfcinv":
      expr, err = p.parse_math1_handler( token, exec.ErfcinvHandler )
    case "erfinv":
      expr, err = p.parse_math1_handler( token, exec.ErfinvHandler )
    case "exp":
      expr, err = p.parse_math1_handler( token, exec.ExpHandler )
    case "exp2":
      expr, err = p.parse_math1_handler( token, exec.Exp2Handler )
    case "expm1":
      expr, err = p.parse_math1_handler( token, exec.Expm1Handler )
    case "floor":
      expr, err = p.parse_math1_handler( token, exec.FloorHandler )

    case "gamma":
      expr, err = p.parse_math1_handler( token, exec.GammaHandler )

    case "ilogb":
      expr, err = p.parse_math1_handler( token, exec.IlogbHandler )
    case "isnan":
      expr, err = p.parse_math1_handler( token, exec.IsNaNHandler )

    case "j0":
      expr, err = p.parse_math1_handler( token, exec.J0Handler )
    case "j1":
      expr, err = p.parse_math1_handler( token, exec.J1Handler )

    case "log":
      expr, err = p.parse_math1_handler( token, exec.LogHandler )
    case "log10":
      expr, err = p.parse_math1_handler( token, exec.Log10Handler )
    case "log1p":
      expr, err = p.parse_math1_handler( token, exec.Log1pHandler )
    case "log2":
      expr, err = p.parse_math1_handler( token, exec.Log2Handler )
    case "logb":
      expr, err = p.parse_math1_handler( token, exec.LogbHandler )

    case "pow10":
      expr, err = p.parse_math1_handler( token, exec.Pow10Handler )

    case "round":
      expr, err = p.parse_math1_handler( token, exec.RoundHandler )

    case "roundtoeven":
      expr, err = p.parse_math1_handler( token, exec.Round2evenHandler )

    case "sin":
      expr, err = p.parse_math1_handler( token, exec.SinHandler )
    case "sinh":
      expr, err = p.parse_math1_handler( token, exec.SinhHandler )
    case "sqrt":
      expr, err = p.parse_math1_handler( token, exec.SqrtHandler )

    case "tan":
      expr, err = p.parse_math1_handler( token, exec.TanHandler )
    case "tanh":
      expr, err = p.parse_math1_handler( token, exec.TanhHandler )
    case "trunc":
      expr, err = p.parse_math1_handler( token, exec.TruncHandler )

    case "y0":
      expr, err = p.parse_math1_handler( token, exec.Y0Handler )
    case "y1":
      expr, err = p.parse_math1_handler( token, exec.Y1Handler )

    default:
      expr, err = p.parse_constants()
  }

  return expr, err
}
