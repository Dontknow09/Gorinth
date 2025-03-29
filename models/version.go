// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    version, err := UnmarshalVersion(bytes)
//    bytes, err = version.Marshal()

package models

import (
	"encoding/json"
	"time"
)

func UnmarshalVersion(data []byte) (Version, error) {
	var r Version
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Version) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Version struct {
	GameVersions    []string     `json:"game_versions"`
	Loaders         []string     `json:"loaders"`
	ID              string       `json:"id"`
	ProjectID       string       `json:"project_id"`
	AuthorID        string       `json:"author_id"`
	Featured        bool         `json:"featured"`
	Name            string       `json:"name"`
	VersionNumber   string       `json:"version_number"`
	Changelog       string       `json:"changelog"`
	ChangelogURL    interface{}  `json:"changelog_url"`
	DatePublished   time.Time    `json:"date_published"`
	Downloads       int64        `json:"downloads"`
	VersionType     string       `json:"version_type"`
	Status          string       `json:"status"`
	RequestedStatus interface{}  `json:"requested_status"`
	Files           []File       `json:"files"`
	Dependencies    []Dependency `json:"dependencies"`
}

type Dependency struct {
	VersionID      interface{} `json:"version_id"`
	ProjectID      string      `json:"project_id"`
	FileName       interface{} `json:"file_name"`
	DependencyType string      `json:"dependency_type"`
}

type File struct {
	Hashes   Hashes      `json:"hashes"`
	URL      string      `json:"url"`
	Filename string      `json:"filename"`
	Primary  bool        `json:"primary"`
	Size     int64       `json:"size"`
	FileType interface{} `json:"file_type"`
}

type Hashes struct {
	Sha512 string `json:"sha512"`
	Sha1   string `json:"sha1"`
}
