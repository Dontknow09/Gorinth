// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    user, err := UnmarshalUser(bytes)
//    bytes, err = user.Marshal()

package models

import (
	"encoding/json"
	"time"
)

func UnmarshalUser(data []byte) (User, error) {
	var r User
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *User) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type User struct {
	ID                  string       `json:"id"`
	Username            string       `json:"username"`
	AvatarURL           string       `json:"avatar_url"`
	Bio                 *string      `json:"bio"`
	Created             time.Time    `json:"created"`
	Role                string       `json:"role"`
	Badges              int64        `json:"badges"`
	AuthProviders       interface{}  `json:"auth_providers"`
	Email               interface{}  `json:"email"`
	EmailVerified       interface{}  `json:"email_verified"`
	HasPassword         interface{}  `json:"has_password"`
	HasTotp             interface{}  `json:"has_totp"`
	PayoutData          interface{}  `json:"payout_data"`
	StripeCustomerID    *interface{} `json:"stripe_customer_id"`
	AllowFriendRequests *interface{} `json:"allow_friend_requests"`
	GithubID            interface{}  `json:"github_id"`
}
