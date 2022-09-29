package ulid

import (
	"database/sql/driver"

	"github.com/oklog/ulid/v2"
)

type ULID struct {
	ulid.ULID
}

func Make() ULID {
	id := ulid.Make()

	return ULID{id}
}

func Parse(value string) (ULID, error) {
	id, err := ulid.Parse(value)
	if err != nil {
		return ULID{}, err
	}

	return ULID{id}, nil
}

func (id *ULID) Scan(value interface{}) error {
	switch v := value.(type) {
	case nil:
		return nil
	case string:
		return id.UnmarshalText([]byte(v))
	case []byte:
		return id.UnmarshalText(v)
	}

	return ulid.ErrScanValue
}

func (id ULID) Value() (driver.Value, error) {
	return id.String(), nil
}
