package annotation

// ValueSet 用于定义允许的值集合
type ValueSet struct {
	Message string   `json:"message"`
	Groups  []string `json:"groups"`
	Payload []string `json:"payload"`
	Values  []string `json:"values"`
}

// NewValueSet 创建一个新的 ValueSet 实例
func NewValueSet(message string, groups []string, payload []string, values []string) *ValueSet {
	return &ValueSet{
		Message: message,
		Groups:  groups,
		Payload: payload,
		Values:  values,
	}
}
