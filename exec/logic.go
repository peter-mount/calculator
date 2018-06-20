package exec

var logicFunctions = FuncMap{
  "=":        FuncMapEntry{ equalHandler,             BinaryOp },
  "!=":       FuncMapEntry{ notEqualHandler,          BinaryOp },
  "<":        FuncMapEntry{ lessThanHandler,          BinaryOp },
  "<=":       FuncMapEntry{ lessThanEqualHandler,     BinaryOp },
  ">=":       FuncMapEntry{ greaterThanEqualHandler,  BinaryOp },
  ">":        FuncMapEntry{ greaterThanHandler,       BinaryOp },
//"between":  FuncMapEntry{ betweenHandler, },
  "and":      FuncMapEntry{ andHandler,               BinaryOp },
  "&&":       FuncMapEntry{ andHandler,               BinaryOp },
  "or":       FuncMapEntry{ orHandler,                BinaryOp },
  "||":       FuncMapEntry{ orHandler,                BinaryOp },
  "not":      FuncMapEntry{ notHandler,               BinaryOp },
}

func equalHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_BOOL:
      m.PushBool( a.Bool() == b.Bool() )
    case VAR_INT:
      m.PushBool( a.Int() == b.Int() )
    case VAR_FLOAT:
      m.PushBool( a.Float() == b.Float() )
    case VAR_STRING:
      m.PushBool( a.String() == b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func notEqualHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_BOOL:
      m.PushBool( a.Bool() == b.Bool() )
    case VAR_INT:
      m.PushBool( a.Int() == b.Int() )
    case VAR_FLOAT:
      m.PushBool( a.Float() == b.Float() )
    case VAR_STRING:
      m.PushBool( a.String() == b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func lessThanHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_INT:
      m.PushBool( a.Int() < b.Int() )
    case VAR_FLOAT:
      m.PushBool( a.Float() < b.Float() )
    case VAR_STRING:
      m.PushBool( a.String() < b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func lessThanEqualHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_INT:
      m.PushBool( a.Int() <= b.Int() )
    case VAR_FLOAT:
      m.PushBool( a.Float() <= b.Float() )
    case VAR_STRING:
      m.PushBool( a.String() <= b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func greaterThanEqualHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_INT:
      m.PushBool( a.Int() >= b.Int() )
    case VAR_FLOAT:
      m.PushBool( a.Float() >= b.Float() )
    case VAR_STRING:
      m.PushBool( a.String() >= b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func greaterThanHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_INT:
      m.PushBool( a.Int() > b.Int() )
    case VAR_FLOAT:
      m.PushBool( a.Float() > b.Float() )
    case VAR_STRING:
      m.PushBool( a.String() > b.String() )
    default:
      m.PushBool( false )
  }

  return nil
}

func betweenHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  // The values to be between
  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  // The value to test
  c, err := m.Pop()
  if err != nil {
    return err
  }

  switch a.Type() {
    case VAR_INT:
      ci := c.Int()
      m.PushBool( ci >= a.Int() && ci <= b.Int() )
    case VAR_FLOAT:
      cf := c.Float()
      m.PushBool( cf >= a.Float() && cf <= b.Float() )
    default:
      m.PushBool( false )
  }
  return nil
}

func andHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  m.PushBool( a.Bool() && b.Bool() )
  return nil
}

func orHandler( m *Context, n *Node ) error {
  err := n.Invoke2(m)
  if err != nil {
    return err
  }

  a, b, err := m.Pop2()
  if err != nil {
    return err
  }

  m.PushBool( a.Bool() || b.Bool() )
  return nil
}

func notHandler( m *Context, n *Node ) error {
  err := n.InvokeLhs(m)
  if err != nil {
    return err
  }

  a, err := m.Pop()
  if err != nil {
    return err
  }

  m.PushBool( !a.Bool() )
  return nil
}
