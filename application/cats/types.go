package cats

import (
	"base-go/domain/model"
)

// command
type AddCatIpt struct {
	Name string `json:"name"`
}

// query

// response
type GetCatResp = model.Cat
