package clashy_test

import (
	"strings"
	"testing"
)

func normalizeTag(tag string) string {
	tag = strings.ToUpper(tag)

	var b strings.Builder
	b.Grow(len(tag) + 1)
	b.WriteByte('#')

	for _, r := range tag {
		switch {
		case r == 'O':
			b.WriteByte('0')
		case r >= 'A' && r <= 'Z':
			b.WriteRune(r)
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		}
	}

	return b.String()
}

func snapshotKey(kind, tag string) string {
	return kind + ":" + normalizeTag(tag)
}

func TestBehaviorHelpers(t *testing.T) {
	t.Parallel()

	t.Run("tag_normalization", func(t *testing.T) {
		t.Parallel()

		cases := map[string]string{
			" 123aBc O": "#123ABC0",
			"#2pp":      "#2PP",
			"  #p0q  ":  "#P0Q",
		}
		for input, want := range cases {
			if got := normalizeTag(input); got != want {
				t.Fatalf("normalizeTag(%q) = %q, want %q", input, got, want)
			}
		}
	})

	t.Run("snapshot_keying", func(t *testing.T) {
		t.Parallel()

		if got := snapshotKey("clan", "#2PP"); got != "clan:#2PP" {
			t.Fatalf("unexpected clan snapshot key: %q", got)
		}
		if got := snapshotKey("player", " 2pp "); got != "player:#2PP" {
			t.Fatalf("unexpected player snapshot key: %q", got)
		}
	})
}
