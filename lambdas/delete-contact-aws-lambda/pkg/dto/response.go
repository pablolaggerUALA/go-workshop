package dto

import "encoding/json"

type Response struct {
	Message string `json:"message"`
	Uuid    string `json:"deleted_uuid"`
}

func (c *Response) ToJsonStr() string {
	b, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(b)
}
