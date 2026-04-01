package clashy

import (
	"encoding/json"
	"reflect"
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
	Raw           json.RawMessage
}

func (m responseMeta) RetryAfter() int { return m.ResponseRetry }

func setRawJSON(target any, raw []byte, retry int) {
	if target == nil {
		return
	}
	value := reflect.ValueOf(target)
	if value.Kind() != reflect.Pointer || value.IsNil() {
		return
	}
	elem := value.Elem()
	field := elem.FieldByName("responseMeta")
	if !field.IsValid() || !field.CanSet() {
		return
	}
	field.Set(reflect.ValueOf(responseMeta{
		ResponseRetry: retry,
		Raw:           append(json.RawMessage(nil), raw...),
	}))
}

func FromTimestamp(raw string) (time.Time, error) {
	if raw == "" {
		return time.Time{}, nil
	}
	return time.Parse("20060102T150405.000Z", raw)
}

func applyPostDecode(client *Client, target any) {
	walkValue(client, reflect.ValueOf(target))
}

func walkValue(client *Client, value reflect.Value) {
	if !value.IsValid() {
		return
	}
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return
		}
		if value.CanInterface() {
			if finalizer, ok := value.Interface().(interface{ Finalize() }); ok {
				finalizer.Finalize()
			}
			if hook, ok := value.Interface().(interface{ postDecode(*Client) }); ok {
				hook.postDecode(client)
			}
		}
		walkValue(client, value.Elem())
		return
	}
	switch value.Kind() {
	case reflect.Struct:
		if value.CanAddr() {
			addr := value.Addr()
			if addr.CanInterface() {
				if finalizer, ok := addr.Interface().(interface{ Finalize() }); ok {
					finalizer.Finalize()
				}
				if hook, ok := addr.Interface().(interface{ postDecode(*Client) }); ok {
					hook.postDecode(client)
				}
			}
		}
		for i := 0; i < value.NumField(); i++ {
			walkValue(client, value.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walkValue(client, value.Index(i))
		}
	}
}
