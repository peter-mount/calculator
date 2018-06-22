package main

import (
  "flag"
  "fmt"
  "github.com/peter-mount/calculator/exec"
  "os"
)

func main() {

  flag.Parse()

  calc := &exec.Calculator{}

  for _, arg := range flag.Args() {
    errorExit( calc.Parse( arg ) )
    ctx := &exec.Context{}
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
