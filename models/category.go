// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    category, err := UnmarshalCategory(bytes)
//    bytes, err = category.Marshal()

package models

import "encoding/json"

type Category []CategoryElement

func UnmarshalCategory(data []byte) (Category, error) {
	var r Category
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Category) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CategoryElement struct {
	Icon        string      `json:"icon"`
	Name        string      `json:"name"`
	ProjectType ProjectType `json:"project_type"`
	Header      Header      `json:"header"`
}

type Header string

const (
	Categories        Header = "categories"
	Features          Header = "features"
	PerformanceImpact Header = "performance impact"
	Resolutions       Header = "resolutions"
)

type ProjectType string

const (
	Mod          ProjectType = "mod"
	Modpack      ProjectType = "modpack"
	Resourcepack ProjectType = "resourcepack"
	Shader       ProjectType = "shader"
)
