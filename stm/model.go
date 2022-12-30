package stm

// Statement is a struct that represents a statement in a SQL query.
// We do not use it directly.
type Statement struct {
	Operator string       `json:"operator"`
	Field    string       `json:"field,omitempty"`
	Value    interface{}  `json:"value,omitempty"`
	Children []*Statement `json:"children,omitempty"`
}

func (s *Statement) Build() (string, []interface{}) {
	return buildStatement(s, make([]interface{}, 0))
}
