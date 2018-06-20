package exec

import (
  "strconv"
)

const (
  _ = iota
  VAR_NULL
  VAR_BOOL
  VAR_INT
  VAR_FLOAT
  VAR_STRING
)

// An imutable Value of some kind.
type Value struct {
  varType    int
  boolVal    bool
  intVal     int64
  floatVal   float64
  stringVal  string
}

var nullValue Value = Value{ varType: VAR_NULL }

func NullValue() *Value {
  return &nullValue
}

// The type of this value
func (v *Value) Type() int {
  return v.varType
}

func BoolValue( i bool ) *Value {
  return &Value{ varType: VAR_BOOL, boolVal: i }
}

func IntValue( i int64 ) *Value {
  return &Value{ varType: VAR_INT, intVal: i }
}

func FloatValue( i float64 ) *Value {
  return &Value{ varType: VAR_FLOAT, floatVal: i }
}

func StringValue( i string ) *Value {
  return &Value{ varType: VAR_STRING, stringVal: i }
}

func (v *Value) IsNull() bool {
  return v.varType == VAR_NULL
}

func (v *Value) IsZero() bool {
  switch v.varType {
    case VAR_NULL:
      return true
    case VAR_BOOL:
      return !v.boolVal
    case VAR_INT:
      return v.intVal == 0
    case VAR_FLOAT:
      return v.floatVal == 0.0
    case VAR_STRING:
      return v.stringVal == ""
    default:
      return false
  }
}

// Bool returns the value as a bool
func (v *Value) Bool() bool {
  switch v.varType {
    case VAR_BOOL:
      return v.boolVal
    case VAR_INT:
      if v.intVal == 0 {
        return false
      }
      return true
    case VAR_FLOAT:
      if v.floatVal == 0 {
        return false
      }
      return true
    case VAR_STRING:
      if v.stringVal == "" || v.stringVal[0]=='f' || v.stringVal[0]=='F' {
        return false
      }
      return true
    default:
      return false
  }
}

// Int returns the value as an int64
func (v *Value) Int() int64 {
  switch v.varType {
    case VAR_BOOL:
      if v.boolVal {
        return 1
      }
      return 0
    case VAR_INT:
      return v.intVal
    case VAR_FLOAT:
      return int64(v.floatVal)
    case VAR_STRING:
      r, err := strconv.ParseInt( v.stringVal, 10, 64 )
      if err == nil {
        return r
      }
      // Panic instead?
      return 0
    default:
      return 0
  }
}

// Float returns the value as a float64
func (v *Value) Float() float64 {
  switch v.varType {
    case VAR_BOOL:
      if v.boolVal {
        return 1.0
      }
      return 0.0
    case VAR_INT:
      return float64(v.intVal)
    case VAR_FLOAT:
      return v.floatVal
    case VAR_STRING:
      r, err := strconv.ParseFloat( v.stringVal, 64 )
      if err == nil {
        return r
      }
      // Panic instead?
      return 0.0
    default:
      return 0.0
  }
}

// String returns the value as a string
func (v *Value) String() string {
  switch v.varType {
    case VAR_BOOL:
      if v.boolVal {
        return "true"
      }
      return "false"
    case VAR_INT:
      return strconv.FormatInt( v.intVal, 10 )
    case VAR_FLOAT:
      return strconv.FormatFloat( v.floatVal, 'f', 10, 16 )
    case VAR_STRING:
      return v.stringVal
    default:
      return ""
  }
}
