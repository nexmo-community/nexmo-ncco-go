package ncco

import (
	"encoding/json"
)

type action string

const (
	RECORD action = "record"
)

type Ncco struct {
	Actions []interface{} `json:"omitempty"`
}

type Action interface {
	WithType() interface{}
}

func (n Ncco) Json() (string, error) {
	ncco, err := json.Marshal(n.Actions)
	return string(ncco), err
}

func New(actions ...Action) Ncco {
	translatedActions := make([]interface{}, 0)
	for _, action := range actions {
		translatedActions = append(translatedActions, action.WithType())
	}
	return Ncco{translatedActions}
}
