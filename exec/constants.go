package exec

import (
  "math"
)

var constants = map[string]*Value{
  // Constants defined in math package
  "e": FloatValue( math.E ),
  "pi": FloatValue( math.Pi ),
  "phi": FloatValue( math.Phi ),
  "sqrt2": FloatValue( math.Sqrt2 ),
  "sqrte": FloatValue( math.SqrtE ),
  "sqrtpi": FloatValue( math.SqrtPi ),
  "sqrtphi": FloatValue( math.SqrtPhi ),
  "ln2": FloatValue( math.Ln2 ),
  "log2e": FloatValue( math.Log2E ),
  "ln10": FloatValue( math.Ln10 ),
  "log10e": FloatValue( math.Log10E ),
  "nan": FloatValue( math.NaN() ),
}

// parse_constants Handles mathematical constants
func (p *Parser) parse_constants() (*Node,error) {
  token := p.lexer.Peek()

  if value, exists := constants[token.Text()]; exists {
    token = p.lexer.Next()
    return &Node{ token:token.Text(), value: value }, nil
  }

  expr, err := p.parse_unary()
  return expr, err
}
