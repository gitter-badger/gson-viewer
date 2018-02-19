package goson

import (
	"bytes"
	"encoding/json"
)

//Goson is goson base struct
type Goson struct {
	jsonObject interface{}
}

//NewGoson - Returns Goson instance
func NewGoson(data []byte) (*Goson, error) {
	g := new(Goson)

	if err := json.Unmarshal(data, &g.jsonObject); err != nil {
		return nil, err
	}
	return g, nil
}

//StringIndent - Convert json object to pretty json string
func (g *Goson) StringIndent(prefix, indent string) (string, error) {
	jsonData, err := json.Marshal(g.jsonObject)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := json.Indent(&buf, jsonData, prefix, indent); err != nil {
		return "", err
	}

	return buf.String(), nil
}

//Search - Return json value corresponding to keys. keys represents key of hierarchy of json
func (g *Goson) Search(keys ...string) (interface{}, error) {
	return nil, nil
}

func search(key string, object interface{}) (interface{}, error) {
	return nil, nil
}
