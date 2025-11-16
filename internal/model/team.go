package model

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"slices"
	"time"
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
	candidates := make(TeamMembers, 0)
	for _, member := range *tm {
		if member.IsActive && !slices.Contains(excludeIds, member.UserID) {
			candidates = append(candidates, member)
		}
	}
	if maxCount <= 0 || len(candidates) == 0 {
		return make(TeamMembers, 0)
	}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rnd.Shuffle(len(candidates), func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})
	if len(candidates) > maxCount {
		return candidates[:maxCount]
	}
	return candidates
}

type Team struct {
	TeamName string      `json:"team_name" db:"team_name" validate:"required,max=64"`
	Members  TeamMembers `json:"members" db:"members" validate:"dive"`
}
