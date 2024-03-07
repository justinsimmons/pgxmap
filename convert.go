package pgxmap

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// Text converts a native go string type to a pgtype.Text,
// required for DB operation on pgx/v5.
func Text[T string | *string](text T) pgtype.Text {
	txt := pgtype.Text{}

	switch v := any(text).(type) {
	case string:
		txt.String = v
		txt.Valid = true
	case *string:
		if v != nil {
			txt.String = *v
			txt.Valid = true
		}
	}

	return txt
}

// Timestamptz converts a native go time.Time type to a pgtype.Timestamptz,
// required for DB operation on pgx/v5.
// TODO: Support special values for infinity: (https://www.postgresql.org/docs/current/datatype-datetime.html#DATATYPE-DATETIME-SPECIAL-VALUES).
func Timestamptz[T time.Time | *time.Time](timestamp T) pgtype.Timestamptz {
	ts := pgtype.Timestamptz{}

	switch v := any(timestamp).(type) {
	case time.Time:
		ts.Time = v
		ts.Valid = true
	case *time.Time:
		if v != nil {
			ts.Time = *v
			ts.Valid = true
		}
	}

	return ts
}

// Bool converts a native go bool type to a pgtype.Bool,
// required for DB operation on pgx/v5.
func Bool[T bool | *bool](b T) pgtype.Bool {
	o := pgtype.Bool{}

	switch v := any(b).(type) {
	case bool:
		o.Bool = v
		o.Valid = true
	case *bool:
		if v != nil {
			o.Bool = *v
			o.Valid = true
		}
	}

	return o
}
