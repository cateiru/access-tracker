package types

import "encoding/json"

type Types struct {
	Text interface{}
}

func New(types interface{}) *Types {
	return &Types{
		Text: types,
	}
}

func (c *Types) GetJson() ([]byte, error) {
	bytes, err := json.Marshal(c.Text)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}
