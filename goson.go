package goson

import (
	"bytes"
	"encoding/json"
)

//Goson is goson base struct
type Goson struct {
	jsonObject interface{}
}

//NewGoson returns Goson instance
func NewGoson(data []byte) (*Goson, error) {
	g := new(Goson)

	if err := json.Unmarshal(data, &g.jsonObject); err != nil {
		return nil, err
	}
	return g, nil
}

//JSONObjectToPrettyJSONString convert json object to pretty json string
func (g *Goson) JSONObjectToPrettyJSONString() (string, error) {
	jsonData, err := json.Marshal(g.jsonObject)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := json.Indent(&buf, jsonData, "", " "); err != nil {
		return "", err
	}

	return buf.String(), nil
}

//Query return json value corresponding to keys. keys represents key of hierarchy of json
func (g *Goson) Query(keys ...string) (interface{}, error) {
	return nil, nil
}

func query(key string, object interface{}) (interface{}, error) {
	return nil, nil
}
