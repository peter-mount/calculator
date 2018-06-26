package exec

import (
  "fmt"
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
func HtmlTree( r *Node, w io.Writer, title string ) {
  m := make( map[*Node]interface{} )
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

func logTree( m map[*Node]interface{}, p *nodeCell, r *Node ) {
  // Prevent infinite loops - should not happen except if a bug happens in the tree
  if _, visited := m[r]; visited {
    // We've already visited this which is an error
    p.append( "Looping:&nbsp;" + r.token )
    return
  }
  m[r] = nil
  defer delete( m, r )

  c := p.append( r.token )

  if r.left != nil {
    logTree( m, c, r.left )
  }
  if r.right != nil {
    logTree( m, c, r.right )
  }
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

// StackDump writes the current state of the stack to a Writer
func (c *Context) StackDump( w io.Writer) {
  io.WriteString( w, "[")
  for i, v := range c.stack {
    if i>0 {
      io.WriteString( w, ", ")
    }
    s := v.Type() == VAR_STRING
    if s {
      io.WriteString( w, "\"")
    }
    io.WriteString( w, v.String() )
    if s {
      io.WriteString( w, "\"")
    }
  }
  io.WriteString( w, "]\n")
}
