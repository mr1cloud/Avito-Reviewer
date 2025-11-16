package model

import (
	"encoding/json"
	"fmt"
	"slices"
)

type TeamMember struct {
	UserID   string `json:"user_id" db:"user_id" validate:"required,max=32"`
	Username string `json:"username" db:"username" validate:"required,max=64"`
	IsActive bool   `json:"is_active" db:"is_active" validate:"required"`
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

func (tm *TeamMembers) Contains(userID string) bool {
	for _, member := range *tm {
		if member.UserID == userID {
			return true
		}
	}
	return false
}

func (tm *TeamMembers) GetMembersCount() int {
	return len(*tm)
}

func (tm *TeamMembers) GetActiveMembers(maxCount int, excludeIds ...string) TeamMembers {
	activeMembers := make(TeamMembers, 0, maxCount)
	for _, member := range *tm {
		if member.IsActive && !slices.Contains(excludeIds, member.UserID) {
			activeMembers = append(activeMembers, member)
			if len(activeMembers) >= maxCount {
				break
			}
		}
	}
	return activeMembers
}

type Team struct {
	TeamName string      `json:"team_name" db:"team_name" validate:"required,max=64"`
	Members  TeamMembers `json:"members" db:"members" validate:"dive"`
}
