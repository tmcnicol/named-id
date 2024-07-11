package id

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/oklog/ulid/v2"
)

type ID struct {
	prefix prefix
	id     ulid.ULID
}

func (id ID) Prefix() string {
	return string(id.prefix)
}

func (id ID) String() string {
	result := strings.Join([]string{string(id.prefix), id.id.String()}, "_")
	return strings.ToLower(result)
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id ID) Nil() bool {
	def := ID{}
	return id == def
}

func New(key string) *ID {
	pref := prefixes[key]
	if pref == "" {
		slog.Error("ulid error no prefix", "prefix", key)
		return nil
	}
	return new(pref)
}

func MustNew(key string) ID {
	id := New(key)
	if id == nil {
		panic(fmt.Sprintf("could not create id: %s", key))
	}
	return *id
}

func MustParse(raw string) ID {
	id, err := Parse(raw)
	if err != nil {
		panic(fmt.Sprintf("could not parse id: %s", raw))
	}
	return id
}

func NewString(key string) string {
	pref := prefixes[key]
	if pref == "" {
		slog.Error("ulid error no prefix", "prefix", key)
		return ""
	}
	return new(pref).String()
}

func new(pref prefix) *ID {
	id := ulid.Make()
	return &ID{
		prefix: pref,
		id:     id,
	}
}

func Parse(raw string) (ID, error) {
	id, err := parse(raw)
	if err != nil {
		return ID{}, err
	}
	if !id.prefix.isValid() {
		return ID{}, ErrInvalidPrefix
	}
	return id, nil
}

func ParseType(raw string, acceptedKeys ...string) (ID, error) {
	id, err := parse(raw)
	if err != nil {
		return ID{}, err
	}

	if !id.prefix.isValid(acceptedKeys...) {
		return ID{}, ErrInvalidPrefix
	}

	return id, nil
}

func parse(raw string) (ID, error) {
	parts := strings.Split(raw, "_")
	if len(parts) != 2 {
		return ID{}, ErrCouldNotParse
	}
	pref := parts[0]
	id, err := ulid.Parse(parts[1])
	if err != nil {
		return ID{}, ErrCouldNotParse
	}

	return ID{
		prefix(pref),
		id,
	}, nil
}

var ErrCouldNotParse = fmt.Errorf("could not parse")
var ErrInvalidPrefix = fmt.Errorf("invalid prefix")
