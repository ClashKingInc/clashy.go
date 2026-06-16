package clashy

import (
	"time"
)

// RankedClan is a clan ranking entry.
type RankedClan struct {
	Clan
	// Rank is the current ranking position.
	Rank int `json:"rank,omitempty"`
	// PreviousRank is the previous ranking position when the API provides it.
	PreviousRank int `json:"previousRank,omitempty"`
}

// RankedPlayer is a player ranking entry.
type RankedPlayer struct {
	Player
	// League is the player's ranking league when the endpoint includes it.
	League League `json:"league,omitempty"`
	// AttackWins is the player's attack win count in the ranking.
	AttackWins int `json:"attackWins,omitempty"`
	// DefenseWins is the player's defense win count in the ranking.
	DefenseWins int `json:"defenseWins,omitempty"`
	// Rank is the current ranking position.
	Rank int `json:"rank,omitempty"`
	// PreviousRank is the previous ranking position when the API provides it.
	PreviousRank int `json:"previousRank,omitempty"`
}

type responseMeta struct {
	ResponseRetry int
}

// RetryAfter returns the number of seconds the API says this response can be
// cached, derived from Cache-Control max-age.
func (m responseMeta) RetryAfter() int { return m.ResponseRetry }
func (m *responseMeta) setResponseMeta(meta responseMeta) {
	if m == nil {
		return
	}
	*m = meta
}

func applyResponseMeta(target any, retry int) {
	if target == nil {
		return
	}
	if setter, ok := target.(interface{ setResponseMeta(responseMeta) }); ok {
		setter.setResponseMeta(responseMeta{ResponseRetry: retry})
	}
}

// FromTimestamp parses a Clash API timestamp in 20060102T150405.000Z format.
func FromTimestamp(raw string) (time.Time, error) {
	if raw == "" {
		return time.Time{}, nil
	}
	return time.Parse("20060102T150405.000Z", raw)
}
