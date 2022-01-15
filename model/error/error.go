package error

type ValidationError struct {
	Loc  []string `json:"loc,omitempty"`
	Msg  string   `json:"msg,omitempty"`
	Type string   `json:"type,omitempty"`
}
