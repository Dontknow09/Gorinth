// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    license, err := UnmarshalLicense(bytes)
//    bytes, err = license.Marshal()

package models

import "encoding/json"

func UnmarshalLicense(data []byte) (License, error) {
	var r License
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *License) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type License struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
