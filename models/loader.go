// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    loader, err := UnmarshalLoader(bytes)
//    bytes, err = loader.Marshal()

package models

import "encoding/json"

type Loader []LoaderElement

func UnmarshalLoader(data []byte) (Loader, error) {
	var r Loader
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Loader) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type LoaderElement struct {
	Icon                  string                 `json:"icon"`
	Name                  string                 `json:"name"`
	SupportedProjectTypes []SupportedProjectType `json:"supported_project_types"`
}

type SupportedProjectType string

const (
	Datapack     SupportedProjectType = "datapack"
	Mod          SupportedProjectType = "mod"
	Modpack      SupportedProjectType = "modpack"
	Plugin       SupportedProjectType = "plugin"
	Project      SupportedProjectType = "project"
	Resourcepack SupportedProjectType = "resourcepack"
	Shader       SupportedProjectType = "shader"
)
