package gorad

import (
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/oklog/ulid/v2"
)

type PULID struct {
	value string
}

// MustNewPULID returns a new PULID for time.Now() given a prefix. This uses the default entropy source.
func MustNewPULID(prefix string) *PULID {
	return &PULID{value: strings.ToUpper(prefix) + ulid.Make().String()}
}

func (u *PULID) String() string {
	return u.value
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *PULID) UnmarshalGQL(v any) error {
	return u.Scan(v)
}

// MarshalGQL implements the graphql.Marshaller interface
func (u *PULID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(u.String()))
}

// Scan implements the Scanner interface.
func (u *PULID) Scan(src any) error {
	if src == nil {
		return fmt.Errorf("pulid: expected a value")
	}
	s, ok := src.(string)
	if !ok {
		return fmt.Errorf("pulid: expected a string")
	}
	u.value = s
	return nil
}

// Value implements the driver Valuer interface.
func (u *PULID) Value() (driver.Value, error) {
	return u.String(), nil
}

func (u *PULID) Prefix() string {
	return u.String()[:3]
}

func (u *PULID) HasPrefix(prefix string) bool {
	return u.Prefix() == prefix
}
