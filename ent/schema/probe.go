package schema

import (
	"encoding/json" // Required for json.RawMessage
	"time"          // Required for time fields

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/Nathene/SyntheticMonitor/cmd/types"
)

// Probe holds the schema definition for the Probe entity.
type Probe struct {
	ent.Schema
}

// Fields of the Probe.
// We DO NOT define an 'id' field here. Ent creates an auto-incrementing
// integer 'id' field by default, which is usually what you want.
func (Probe) Fields() []ent.Field {
	return []ent.Field{
		// --- Probe Status ---
		field.Enum("status").
			// Reference the Go type from your types package
			GoType(types.ProbeStatus("")).
			// Set default using the constant from your types package
			Default(string(types.Disabled)),

		// --- Probe Type ---
		field.Enum("type").
			GoType(types.ProbeType("")).
			// Type cannot be changed after creation
			Immutable(),

		// --- Type-Specific Configuration ---
		// Stores varying config (URL/DSN etc.) as JSON
		// json.RawMessage delays unmarshalling until needed
		field.JSON("config", json.RawMessage{}).
			// Config might not apply to all types or is set later
			Optional(),

		// --- Timestamps ---
		field.Time("created_at").
			// Set automatically on creation
			Default(time.Now).
			// Cannot be changed after creation
			Immutable(),
		field.Time("updated_at").
			// Set automatically on creation
			Default(time.Now).
			// Set automatically on update
			UpdateDefault(time.Now),
		field.Time("last_heartbeat").
			// This timestamp is optional
			Optional().
			// Allows the field to be NULL in the DB (results in *time.Time in Go)
			Nillable(),
	}
}

// Edges of the Probe.
// We'll define relationships to other entities here later if needed.
func (Probe) Edges() []ent.Edge {
	return nil
}
