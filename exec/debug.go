package exec

import (
  "fmt"
  "io"
)

// HtmlTree will write a simple html document to a writer showing the layout
// of a Node tree. This is useful in debugging the parser.
func (r *Node) HtmlTree( w io.Writer ) {
  m := make( map[*Node]interface{} )

  io.WriteString( w, "<html><head><style>body > table table {border:1px solid grey;} td {vertical-align:top;}</style><body>" )

  c := &nodeCell{}
  r.logTree( m, c )
  c.print( w )

  io.WriteString( w, "</body></html>" )
}

func (r *Node) logTree( m map[*Node]interface{}, p *nodeCell ) {
  // Prevent infinite loops - should not happen except if a bug happens in the tree
  if _, visited := m[r]; visited {
    // We've already visited this which is an error
    c := p.append( "Looping:&nbsp;" + r.token )
    return
  }
  m[r] = nil
  defer delete( m, r )

  c := p.append( r.token )

  if r.lhs != nil {
    r.lhs.logTree( m, c )
  }
  if r.rhs != nil {
    r.rhs.logTree( m, c )
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
