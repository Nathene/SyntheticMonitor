// Code generated by ent, DO NOT EDIT.

package probe

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Nathene/SyntheticMonitor/cmd/types"
	"github.com/Nathene/SyntheticMonitor/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Probe {
	return predicate.Probe(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Probe {
	return predicate.Probe(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Probe {
	return predicate.Probe(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Probe {
	return predicate.Probe(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Probe {
	return predicate.Probe(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Probe {
	return predicate.Probe(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Probe {
	return predicate.Probe(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Probe {
	return predicate.Probe(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Probe {
	return predicate.Probe(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldEQ(FieldUpdatedAt, v))
}

// LastHeartbeat applies equality check predicate on the "last_heartbeat" field. It's identical to LastHeartbeatEQ.
func LastHeartbeat(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldEQ(FieldLastHeartbeat, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v types.ProbeStatus) predicate.Probe {
	vc := v
	return predicate.Probe(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v types.ProbeStatus) predicate.Probe {
	vc := v
	return predicate.Probe(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...types.ProbeStatus) predicate.Probe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Probe(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...types.ProbeStatus) predicate.Probe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Probe(sql.FieldNotIn(FieldStatus, v...))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v types.ProbeType) predicate.Probe {
	vc := v
	return predicate.Probe(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v types.ProbeType) predicate.Probe {
	vc := v
	return predicate.Probe(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...types.ProbeType) predicate.Probe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Probe(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...types.ProbeType) predicate.Probe {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Probe(sql.FieldNotIn(FieldType, v...))
}

// ConfigIsNil applies the IsNil predicate on the "config" field.
func ConfigIsNil() predicate.Probe {
	return predicate.Probe(sql.FieldIsNull(FieldConfig))
}

// ConfigNotNil applies the NotNil predicate on the "config" field.
func ConfigNotNil() predicate.Probe {
	return predicate.Probe(sql.FieldNotNull(FieldConfig))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldLTE(FieldUpdatedAt, v))
}

// LastHeartbeatEQ applies the EQ predicate on the "last_heartbeat" field.
func LastHeartbeatEQ(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldEQ(FieldLastHeartbeat, v))
}

// LastHeartbeatNEQ applies the NEQ predicate on the "last_heartbeat" field.
func LastHeartbeatNEQ(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldNEQ(FieldLastHeartbeat, v))
}

// LastHeartbeatIn applies the In predicate on the "last_heartbeat" field.
func LastHeartbeatIn(vs ...time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldIn(FieldLastHeartbeat, vs...))
}

// LastHeartbeatNotIn applies the NotIn predicate on the "last_heartbeat" field.
func LastHeartbeatNotIn(vs ...time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldNotIn(FieldLastHeartbeat, vs...))
}

// LastHeartbeatGT applies the GT predicate on the "last_heartbeat" field.
func LastHeartbeatGT(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldGT(FieldLastHeartbeat, v))
}

// LastHeartbeatGTE applies the GTE predicate on the "last_heartbeat" field.
func LastHeartbeatGTE(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldGTE(FieldLastHeartbeat, v))
}

// LastHeartbeatLT applies the LT predicate on the "last_heartbeat" field.
func LastHeartbeatLT(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldLT(FieldLastHeartbeat, v))
}

// LastHeartbeatLTE applies the LTE predicate on the "last_heartbeat" field.
func LastHeartbeatLTE(v time.Time) predicate.Probe {
	return predicate.Probe(sql.FieldLTE(FieldLastHeartbeat, v))
}

// LastHeartbeatIsNil applies the IsNil predicate on the "last_heartbeat" field.
func LastHeartbeatIsNil() predicate.Probe {
	return predicate.Probe(sql.FieldIsNull(FieldLastHeartbeat))
}

// LastHeartbeatNotNil applies the NotNil predicate on the "last_heartbeat" field.
func LastHeartbeatNotNil() predicate.Probe {
	return predicate.Probe(sql.FieldNotNull(FieldLastHeartbeat))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Probe) predicate.Probe {
	return predicate.Probe(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Probe) predicate.Probe {
	return predicate.Probe(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Probe) predicate.Probe {
	return predicate.Probe(sql.NotPredicates(p))
}
