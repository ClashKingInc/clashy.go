#!/usr/bin/env sh
set -eu

go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
gomarkdoc_bin="${GOBIN:-$(go env GOPATH)/bin}/gomarkdoc"

mkdir -p docs/api
find docs/api -type f -name '*.md' -delete

"$gomarkdoc_bin" --exclude-dirs ./tests --output 'docs/api/{{.Dir}}.md' ./...

# gomarkdoc's .Dir value for the module root is ".". Keep the generated page
# stable for MkDocs navigation while preserving the requested output template.
if [ -f docs/api/..md ]; then
	mv docs/api/..md docs/api/clashy.md
fi
if [ -f docs/api/.md ]; then
	mv docs/api/.md docs/api/clashy.md
fi

go run ./scripts/generate-api-sections.go
