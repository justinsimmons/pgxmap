package pgxmap

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// UUID converts a google UUID to a pgtype.UUID.
// TODO: Should we allow uuid.Nil to be valid?
func UUID[T uuid.UUID | *uuid.UUID](id T) pgtype.UUID {
	uid := pgtype.UUID{}

	switch v := any(id).(type) {
	case uuid.UUID:
		uid.Bytes = v
		uid.Valid = true
	case *uuid.UUID:
		if v != nil {
			uid.Bytes = *v
			uid.Valid = true
		}
	}

	return uid
}

// UnwrapUUID unwraps the pgtype.UUID to a google UUID struct.
// If the value is null the function will return uuid.Nil.
// Use UnwrapUUIDPointer() if you want a pointer value.
func UnwrapUUID(id pgtype.UUID) uuid.UUID {
	if !id.Valid {
		return uuid.Nil
	}

	return uuid.UUID(id.Bytes)
}

// UnwrapUUIDPtr unwraps the pgtype.UUID to a google UUID pointer.
func UnwrapUUIDPtr(id pgtype.UUID) *uuid.UUID {
	if !id.Valid {
		return nil
	}

	nID := uuid.UUID(id.Bytes)

	return &nID
}
