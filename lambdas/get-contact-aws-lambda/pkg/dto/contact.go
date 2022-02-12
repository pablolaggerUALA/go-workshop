package dto

import "encoding/json"

type Contact struct {
	Uuid      string `dynamodbav:"Uuid"`
	FirstName string `dynamodbav:"FirstName"`
	LastName  string `dynamodbav:"LastName"`
	Status    string `dynamodbav:"Status"`
}

func (c *Contact) ToJsonStr() string {
	b, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(b)
}
