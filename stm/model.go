package stm

type Statement struct {
	Operator string       `json:"operator"`
	Field    string       `json:"field,omitempty"`
	Value    interface{}  `json:"value,omitempty"`
	Children []*Statement `json:"children,omitempty"`
}

func (s *Statement) Build() (string, []interface{}) {
	return buildStatement(s, make([]interface{}, 0))
}
