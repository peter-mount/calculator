package debug

import (
  "fmt"
  "github.com/peter-mount/calculator/context"
  "io"
)

// HtmlTree will write a simple html document to a writer showing the layout
// of a Node tree. This is useful in debugging the parser.
//
// For example:
//
// f, err := os.OpenFile( "/tmp/out.html", os.O_CREATE | os.O_TRUNC|os.O_WRONLY, 0666 )
// if err != nil { return nil }
// defer f.Close()
// e := "1 + 2 * 3 + 49"
// HtmlTreeStart( f )
// parser := calc.Parser()
// err = parser.Parse( e )
// if err != nil {
//   t.Error( err )
// } else {
//   HtmlTree( parser.GetRoot(), f, e )
// }
// HtmlTreeEnd( f )
//
func HtmlTree( r *context.Node, w io.Writer, title string ) {
  m := make( map[*context.Node]interface{} )
  c := &nodeCell{t:title}
  if r != nil {
    logTree( m, c, r )
  }
  c.print( w )
}

func HtmlTreeStart( w io.Writer ) {
  //body > table
  io.WriteString( w, "<html><head><style>" )
  io.WriteString( w, " table {border:1px solid grey;}" )
  io.WriteString( w, " body > table {display:inline-table; margin: 0.25em;}" )
  io.WriteString( w, " body > table table {width:100%;}" )
  io.WriteString( w, " td {vertical-align:top;width:50%;}" )
  io.WriteString( w, " th {width:50%;}" )
  io.WriteString( w, "</style><body>" )
}

func HtmlTreeEnd( w io.Writer ) {
  io.WriteString( w, "</body></html>" )
}

func logTree( m map[*context.Node]interface{}, p *nodeCell, r *context.Node ) {
  // Prevent infinite loops - should not happen except if a bug happens in the tree
  if _, visited := m[r]; visited {
    // We've already visited this which is an error
    p.append( "Looping:&nbsp;" + r.Token().Text() )
    return
  }
  m[r] = nil
  defer delete( m, r )

  c := p.append( r.Token().Text() )

  r.ForEachAll( func(n *context.Node) error {
    logTree( m, c, n )
    return nil
  })
}

type nodeCell struct {
  t string
  c []*nodeCell
}

func (n *nodeCell) width() int {
  w := len(n.t)
  for _, cc := range n.c {
    w = logMax( w, cc.width() )
  }
  return w
}

func (n *nodeCell) append( t string ) *nodeCell {
  nc := &nodeCell{t:t}
  n.c = append( n.c, nc )
  return nc
}

func (n *nodeCell) print( w io.Writer ) {
  io.WriteString( w, fmt.Sprintf( "<table><tr><th colspan=\"%d\">", logMax( 1, len(n.c) ) ) )
  io.WriteString( w, n.t )
  io.WriteString( w, "</th></tr>" )

  if len( n.c ) > 0 {
    io.WriteString( w, "<tr>" )
    for _, c := range n.c {
      io.WriteString( w, "<td align=\"center\">" )
      c.print( w )
      io.WriteString( w, "</td>" )
    }
    io.WriteString( w, "</tr>" )
  }

  io.WriteString( w, "</table>" )
}

func logMax( a, b int ) int {
  if a > b {
    return a
  }
  return b
}
