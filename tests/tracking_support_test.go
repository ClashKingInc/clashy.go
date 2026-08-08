package clashy_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	clashy "github.com/clashkinginc/clashy.go"
)

func TestBattleLogEntryParsesTimestamp(t *testing.T) {
	t.Parallel()

	var entry clashy.BattleLogEntry
	if err := json.Unmarshal([]byte(`{"timestamp":"20260716T123456.000Z","stars":3}`), &entry); err != nil {
		t.Fatalf("unmarshal battle log entry: %v", err)
	}

	want := time.Date(2026, time.July, 16, 12, 34, 56, 0, time.UTC)
	if !entry.Timestamp.Equal(want) {
		t.Fatalf("unexpected timestamp: got %s, want %s", entry.Timestamp, want)
	}
	if entry.Stars != 3 {
		t.Fatalf("unexpected stars: got %d, want 3", entry.Stars)
	}
}

func TestTimestampMarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		timestamp clashy.Timestamp
		want      string
	}{
		{
			name:      "preserves raw API value",
			timestamp: clashy.Timestamp{RawTime: "20260716T123456.789Z"},
			want:      `"20260716T123456.789Z"`,
		},
		{
			name:      "formats parsed time",
			timestamp: clashy.Timestamp{Time: time.Date(2026, time.July, 16, 12, 34, 56, 0, time.FixedZone("CEST", 2*60*60))},
			want:      `"20260716T103456.000Z"`,
		},
		{
			name: "marshals zero value as null",
			want: "null",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := json.Marshal(test.timestamp)
			if err != nil {
				t.Fatalf("marshal timestamp: %v", err)
			}
			if string(got) != test.want {
				t.Fatalf("unexpected JSON: got %s, want %s", got, test.want)
			}
		})
	}
}

func TestResponseMetadataIsNotSerialized(t *testing.T) {
	t.Parallel()

	data, err := json.Marshal(clashy.RankedClan{})
	if err != nil {
		t.Fatalf("marshal ranked clan: %v", err)
	}
	if strings.Contains(string(data), "ResponseRetry") {
		t.Fatalf("internal response metadata leaked into JSON: %s", data)
	}
}
