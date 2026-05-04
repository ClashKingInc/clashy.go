package clashy

import (
	"time"
)

type RankedClan struct {
	Clan
	Rank         int `json:"rank,omitempty"`
	PreviousRank int `json:"previousRank,omitempty"`
}

type RankedPlayer struct {
	Player
	Rank         int `json:"rank,omitempty"`
	PreviousRank int `json:"previousRank,omitempty"`
}

type responseMeta struct {
	ResponseRetry int
}

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

func FromTimestamp(raw string) (time.Time, error) {
	if raw == "" {
		return time.Time{}, nil
	}
	return time.Parse("20060102T150405.000Z", raw)
}
