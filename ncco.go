package ncco

import (
	"encoding/json"
)

type Ncco struct {
	actions []interface{}
}

type Action interface {
	WithType() interface{}
}

func (n Ncco) Json() (string, error) {
	ncco, err := json.Marshal(n.actions)
	return string(ncco), err
}

func New(actions ...Action) Ncco {
	// The actions don't have an "action" key-value-pair yet. They need to be converted
	// and then added to the array inside of Ncco.
	translatedActions := make([]interface{}, 0)
	for _, action := range actions {
		translatedActions = append(translatedActions, action.WithType())
	}
	return Ncco{translatedActions}
}
