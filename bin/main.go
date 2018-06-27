package main

import (
  "flag"
  "fmt"
  "github.com/peter-mount/calculator/calculator"
  "github.com/peter-mount/calculator/context"
  "os"
)

func main() {

  flag.Parse()

  calc := &calculator.Calculator{}

  for _, arg := range flag.Args() {
    errorExit( calc.Parse( arg ) )

    ctx := &context.Context{}
    errorExit( calc.Execute( ctx ) )

    result, err := ctx.Pop()
    errorExit(err)

    fmt.Printf("%v\n",result)
  }

}

func errorExit( err error ) {
  if err != nil {
    fmt.Fprintf( os.Stderr, "error: %v\n", err )
    os.Exit(1)
  }
}
