package parser

import (
  "fmt"
  "github.com/peter-mount/calculator/context"
  "github.com/peter-mount/calculator/exec"
  "github.com/peter-mount/calculator/lex"
  "text/scanner"
)

func (p *Parser) parse_statements() (*context.Node,error) {
  block := context.NewBlock( exec.InvokeAllHandler )
  token := p.lexer.Peek()
  for token.Token() != scanner.EOF {
    expr, err := p.parse_statement_block()
    if err != nil {
      return nil, err
    }
    if expr != nil {
      block.Append( expr )
    }
    token = p.lexer.Peek()
  }
  return block, nil
}

func (p *Parser) parse_statement_block() (*context.Node,error) {

  token := p.lexer.Peek()
  if token.Text() == "{" {
    p.lexer.Next()

    block := context.NewBlock( exec.InvokeScopeHandler )
    token := p.lexer.Peek()
    for token.Token() != scanner.EOF && token.Text() != "}" {
      expr, err := p.parse_statement()
      if err != nil {
        return nil, err
      }
      if expr != nil {
        block.Append( expr )
      }
      token = p.lexer.Peek()
    }

    if token.Text() != "}" {
      return nil, fmt.Errorf( "Expecting }" )
    }

    // Skip the }
    p.lexer.Next()
    return block, nil
  }

  expr, err := p.parse_statement()
  return expr, err
}

func (p *Parser) parse_statement() (*context.Node,error) {

  // Skip ; as optional terminators
  token := p.lexer.Peek()
  if token.Text() == "}" {
    token = p.lexer.Next()
    return nil, nil
  }

  for token.Text() == ";" {
    token = p.lexer.Skip()
  }

  var expr *context.Node
  var err error
  switch token.Text() {
    case "if":
      expr, err = p.parse_if()

    case "while":
      expr, err = p.parse_while()

    default:
      // Pass to setvar which will pass on to ParseExpression
      expr, err = p.parse_setvar()
  }

  if err != nil {
    return nil, err
  }
  return expr, nil
}

func (p *Parser) parse_setvar() (*context.Node,error) {

  expr, err := p.ParseExpression()
  if err != nil {
    return nil, err
  }

  token := p.lexer.Peek()
  for token.Text() == "=" && expr != nil && expr.Token().Token() == lex.TOKEN_VARIABLE {
    token = p.lexer.Next()

    left, err := p.ParseExpression()
    if err != nil {
      return nil, err
    }

    expr = context.NewNode( expr.Token(), exec.SetVarHandler, left, nil )

    token = p.lexer.Peek()
  }

  return expr, err
}
