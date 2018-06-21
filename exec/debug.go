package exec

import (
  "log"
  "strings"
)

func (r *Node) LogTree() {
  m := make( map[*Node]interface{} )
  r.logTree( m, 0 )
}

func (r *Node) logTree( m map[*Node]interface{}, depth int ) {
  s := strings.Repeat( " ", depth ) + "+ %s"

  // Prevent infinite loops - should not happen except if a bug happens in the tree
  if _, visited := m[r]; visited {
    log.Printf( s + " INFINITE LOOP?", r.token )
    return
  }
  m[r] = nil
  defer delete( m, r )

  if r.value != nil {
    log.Printf( s + " %d \"%s\"", r.token, r.value.varType, r.value.String() )
  } else {
    log.Printf( s, r.token )
  }

  if r.lhs != nil {
    r.lhs.logTree( m, depth+1 )
  }
  if r.rhs != nil {
    r.rhs.logTree( m, depth+1 )
  }
}
