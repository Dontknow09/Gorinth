// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    project, err := UnmarshalProject(bytes)
//    bytes, err = project.Marshal()

package models

import (
	"encoding/json"
	"time"
)

func UnmarshalProject(data []byte) (Project, error) {
	var r Project
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Project) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Project struct {
	ClientSide           string        `json:"client_side"`
	ServerSide           string        `json:"server_side"`
	GameVersions         []string      `json:"game_versions"`
	ID                   string        `json:"id"`
	Slug                 string        `json:"slug"`
	ProjectType          string        `json:"project_type"`
	Team                 string        `json:"team"`
	Organization         interface{}   `json:"organization"`
	Title                string        `json:"title"`
	Description          string        `json:"description"`
	Body                 string        `json:"body"`
	BodyURL              interface{}   `json:"body_url"`
	Published            time.Time     `json:"published"`
	Updated              time.Time     `json:"updated"`
	Approved             time.Time     `json:"approved"`
	Queued               time.Time     `json:"queued"`
	Status               string        `json:"status"`
	RequestedStatus      string        `json:"requested_status"`
	ModeratorMessage     interface{}   `json:"moderator_message"`
	License              License       `json:"license"`
	Downloads            int64         `json:"downloads"`
	Followers            int64         `json:"followers"`
	Categories           []string      `json:"categories"`
	AdditionalCategories []interface{} `json:"additional_categories"`
	Loaders              []string      `json:"loaders"`
	Versions             []string      `json:"versions"`
	IconURL              string        `json:"icon_url"`
	IssuesURL            string        `json:"issues_url"`
	SourceURL            interface{}   `json:"source_url"`
	WikiURL              interface{}   `json:"wiki_url"`
	DiscordURL           string        `json:"discord_url"`
	DonationUrls         []DonationURL `json:"donation_urls"`
	Gallery              []Gallery     `json:"gallery"`
	Color                int64         `json:"color"`
	ThreadID             string        `json:"thread_id"`
	MonetizationStatus   string        `json:"monetization_status"`
}

type DonationURL struct {
	ID       string `json:"id"`
	Platform string `json:"platform"`
	URL      string `json:"url"`
}

type Gallery struct {
	URL         string    `json:"url"`
	RawURL      string    `json:"raw_url"`
	Featured    bool      `json:"featured"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Ordering    int64     `json:"ordering"`
}

type License struct {
	ID   string      `json:"id"`
	Name string      `json:"name"`
	URL  interface{} `json:"url"`
}
