// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    projectTeamMembers, err := UnmarshalProjectTeamMembers(bytes)
//    bytes, err = projectTeamMembers.Marshal()

package models

import "encoding/json"

type ProjectTeamMembers []ProjectTeamMember

func UnmarshalProjectTeamMembers(data []byte) (ProjectTeamMembers, error) {
	var r ProjectTeamMembers
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectTeamMembers) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ProjectTeamMember struct {
	Role         string      `json:"role"`
	TeamID       string      `json:"team_id"`
	User         User        `json:"user"`
	Permissions  interface{} `json:"permissions"`
	Accepted     bool        `json:"accepted"`
	PayoutsSplit interface{} `json:"payouts_split"`
	Ordering     int64       `json:"ordering"`
}
