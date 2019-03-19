package route

import (
	"encoding/json"
	"fmt"

	"github.com/nurlansu/go-chuck/models/database"
)

// Response is a struct returned as a response
type Response struct {
	Error string         `json:"error,omitempty"`
	Value *database.Item `json:"value,omitempty"`
}

// Serialize returns the json encoding of 'Response'
func (res *Response) Serialize() []byte {
	resp, err := json.Marshal(res)
	if err == nil {
		return resp
	}

	e := fmt.Sprintf("{\"error\":\"Error, marshaling json: %v\"", err)

	return []byte(e)
}
