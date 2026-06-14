// Package clashy provides a Go SDK for the Clash of Clans API and compatible
// ClashKing proxy-style APIs.
//
// The package is organized around the resource shapes exposed by the official
// API: clans, players, wars, Clan War League groups, locations, rankings,
// leagues, labels, gold pass seasons, raid weekends, and player battle logs.
// Client methods accept context.Context values for cancellation and deadlines,
// return typed Go models, and map common API failures onto typed errors that can
// be inspected with errors.As.
//
// A Client can authenticate with existing API tokens through LoginWithTokens or
// with developer-site credentials through Login. The underlying HTTP client
// rotates tokens, applies a simple concurrency limit, decodes compressed
// responses, honors cache headers for GET requests, and preserves retry metadata
// on response models that embed responseMeta.
//
// The package also embeds ClashKing static game data. Static helpers resolve
// troops, spells, heroes, pets, hero equipment, translations, and army-link
// payloads without forcing callers to learn the static JSON layout.
package clashy
