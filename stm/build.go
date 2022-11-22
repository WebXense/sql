package stm

import "strings"

func buildStatement(s *Statement, values []interface{}) (string, []interface{}) {
	var stm string
	switch strings.ToUpper(s.Operator) {
	case OP_AND:
		stm, values = processHasChildrenClause(s, values)
	case OP_OR:
		stm, values = processHasChildrenClause(s, values)
	case OP_NULL:
		stm = s.Field + " " + s.Operator
	case OP_NOT_NULL:
		stm = s.Field + " " + s.Operator
	case OP_BETWEEN:
		cls := &Statement{
			Operator: OP_AND,
			Children: []*Statement{
				{
					Operator: OP_GREATER_EQUAL,
					Field:    s.Field,
					Value:    s.Value.([]interface{})[0],
				},
				{
					Operator: OP_LESS_EQUAL,
					Field:    s.Field,
					Value:    s.Value.([]interface{})[1],
				},
			},
		}
		stm, values = buildStatement(cls, values)
	default:
		stm = s.Field + " " + s.Operator + " ?"
		values = append(values, s.Value)
	}
	return stm, values
}

func processHasChildrenClause(cls *Statement, values []interface{}) (string, []interface{}) {
	var buf []string
	for _, child := range cls.Children {
		s, v := buildStatement(child, values)
		buf = append(buf, s)
		values = v
	}
	stm := "(" + strings.Join(buf, " "+cls.Operator+" ") + ")"
	return stm, values
}
