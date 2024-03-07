// Copyright 2024 Justin Simmons.
//
// This file is part of pgxmap.
// pgxmap is free software: you can redistribute it and/or modify it under the terms of the GNU Lesser General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// pgxmap is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.
// You should have received a copy of the GNU Lesser General Public License along with pgxmap. If not, see <https://www.gnu.org/licenses/>.

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
