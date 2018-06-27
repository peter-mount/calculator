package parser

import (
  "github.com/peter-mount/calculator/context"
  "math"
)

var constants = map[string]*context.Value{
  // Constants defined in math package
  "e":        context.FloatValue( math.E ),
  "pi":       context.FloatValue( math.Pi ),
  "phi":      context.FloatValue( math.Phi ),
  "sqrt2":    context.FloatValue( math.Sqrt2 ),
  "sqrte":    context.FloatValue( math.SqrtE ),
  "sqrtpi":   context.FloatValue( math.SqrtPi ),
  "sqrtphi":  context.FloatValue( math.SqrtPhi ),
  "ln2":      context.FloatValue( math.Ln2 ),
  "log2e":    context.FloatValue( math.Log2E ),
  "ln10":     context.FloatValue( math.Ln10 ),
  "log10e":   context.FloatValue( math.Log10E ),
  "nan":      context.FloatValue( math.NaN() ),
}

// parse_constants Handles mathematical constants
func (p *Parser) parse_constants() (*context.Node,error) {
  token := p.lexer.Peek()

  if value, exists := constants[token.Text()]; exists {
    token = p.lexer.Next()
    return context.NewConstant( token, value ), nil
  }

  expr, err := p.parse_unary()
  return expr, err
}
