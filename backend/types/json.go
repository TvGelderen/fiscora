package types

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func NewNullString(str sql.NullString) NullString {
	return NullString{
		NullString: str,
	}
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return []byte("null"), nil
}

func (s NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.Valid = false
		return nil
	}
	s.Valid = true
	return json.Unmarshal(data, &s.String)
}

func (s NullString) ToSQL() sql.NullString {
	return sql.NullString{
		String: s.String,
		Valid:  s.Valid,
	}
}

type NullInt struct {
	sql.NullInt32
}

func NewNullInt(i sql.NullInt32) NullInt {
    return NullInt{
        NullInt32: i,
    }
}

func (i NullInt) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Int32)
	}
	return []byte("null"), nil
}

func (i NullInt) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		i.Valid = false
		return nil
	}
	i.Valid = true
	return json.Unmarshal(data, &i.Int32)
}

func (i NullInt) ToSQL() sql.NullInt32 {
	return sql.NullInt32{
		Int32: i.Int32,
		Valid: i.Valid,
	}
}

type NullTime struct {
	sql.NullTime
}

func NewNullTime(t sql.NullTime) NullTime {
    return NullTime{
        NullTime: t,
    }
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time)
	}
	return []byte("null"), nil
}

func (t NullTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t.Valid = false
		return nil
	}
	t.Valid = true
	return json.Unmarshal(data, &t.Time)
}

func (t NullTime) ToSQL() sql.NullTime {
	return sql.NullTime{
		Time:  t.Time,
		Valid: t.Valid,
	}
}
