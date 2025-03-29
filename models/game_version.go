// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    gameVersion, err := UnmarshalGameVersion(bytes)
//    bytes, err = gameVersion.Marshal()

package models

import (
	"encoding/json"
	"time"
)

type GameVersion []GameVersionElement

func UnmarshalGameVersion(data []byte) (GameVersion, error) {
	var r GameVersion
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GameVersion) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GameVersionElement struct {
	Version     string      `json:"version"`
	VersionType VersionType `json:"version_type"`
	Date        time.Time   `json:"date"`
	Major       bool        `json:"major"`
}

type VersionType string

const (
	Alpha    VersionType = "alpha"
	Beta     VersionType = "beta"
	Release  VersionType = "release"
	Snapshot VersionType = "snapshot"
)
