package model

import (
	"encoding/json"
	"fmt"
)

type TeamMember struct {
	UserID   string `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	IsActive bool   `json:"is_active" db:"is_active"`
}

type TeamMembers []TeamMember

func (tm *TeamMembers) Scan(src interface{}) error {
	if src == nil {
		*tm = make(TeamMembers, 0)
		return nil
	}

	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, tm)
	case string:
		return json.Unmarshal([]byte(v), tm)
	default:
		return fmt.Errorf("cannot scan %T into TeamMembers", src)
	}
}

type Team struct {
	TeamName string      `json:"team_name" db:"team_name"`
	Members  TeamMembers `json:"members" db:"team_members"`
}
