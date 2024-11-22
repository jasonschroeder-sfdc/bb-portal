// Code generated by ent, DO NOT EDIT.

package build

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Build {
	return predicate.Build(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Build {
	return predicate.Build(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Build {
	return predicate.Build(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Build {
	return predicate.Build(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Build {
	return predicate.Build(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Build {
	return predicate.Build(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Build {
	return predicate.Build(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Build {
	return predicate.Build(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Build {
	return predicate.Build(sql.FieldLTE(FieldID, id))
}

// BuildURL applies equality check predicate on the "build_url" field. It's identical to BuildURLEQ.
func BuildURL(v string) predicate.Build {
	return predicate.Build(sql.FieldEQ(FieldBuildURL, v))
}

// BuildUUID applies equality check predicate on the "build_uuid" field. It's identical to BuildUUIDEQ.
func BuildUUID(v uuid.UUID) predicate.Build {
	return predicate.Build(sql.FieldEQ(FieldBuildUUID, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Build {
	return predicate.Build(sql.FieldEQ(FieldTimestamp, v))
}

// BuildURLEQ applies the EQ predicate on the "build_url" field.
func BuildURLEQ(v string) predicate.Build {
	return predicate.Build(sql.FieldEQ(FieldBuildURL, v))
}

// BuildURLNEQ applies the NEQ predicate on the "build_url" field.
func BuildURLNEQ(v string) predicate.Build {
	return predicate.Build(sql.FieldNEQ(FieldBuildURL, v))
}

// BuildURLIn applies the In predicate on the "build_url" field.
func BuildURLIn(vs ...string) predicate.Build {
	return predicate.Build(sql.FieldIn(FieldBuildURL, vs...))
}

// BuildURLNotIn applies the NotIn predicate on the "build_url" field.
func BuildURLNotIn(vs ...string) predicate.Build {
	return predicate.Build(sql.FieldNotIn(FieldBuildURL, vs...))
}

// BuildURLGT applies the GT predicate on the "build_url" field.
func BuildURLGT(v string) predicate.Build {
	return predicate.Build(sql.FieldGT(FieldBuildURL, v))
}

// BuildURLGTE applies the GTE predicate on the "build_url" field.
func BuildURLGTE(v string) predicate.Build {
	return predicate.Build(sql.FieldGTE(FieldBuildURL, v))
}

// BuildURLLT applies the LT predicate on the "build_url" field.
func BuildURLLT(v string) predicate.Build {
	return predicate.Build(sql.FieldLT(FieldBuildURL, v))
}

// BuildURLLTE applies the LTE predicate on the "build_url" field.
func BuildURLLTE(v string) predicate.Build {
	return predicate.Build(sql.FieldLTE(FieldBuildURL, v))
}

// BuildURLContains applies the Contains predicate on the "build_url" field.
func BuildURLContains(v string) predicate.Build {
	return predicate.Build(sql.FieldContains(FieldBuildURL, v))
}

// BuildURLHasPrefix applies the HasPrefix predicate on the "build_url" field.
func BuildURLHasPrefix(v string) predicate.Build {
	return predicate.Build(sql.FieldHasPrefix(FieldBuildURL, v))
}

// BuildURLHasSuffix applies the HasSuffix predicate on the "build_url" field.
func BuildURLHasSuffix(v string) predicate.Build {
	return predicate.Build(sql.FieldHasSuffix(FieldBuildURL, v))
}

// BuildURLEqualFold applies the EqualFold predicate on the "build_url" field.
func BuildURLEqualFold(v string) predicate.Build {
	return predicate.Build(sql.FieldEqualFold(FieldBuildURL, v))
}

// BuildURLContainsFold applies the ContainsFold predicate on the "build_url" field.
func BuildURLContainsFold(v string) predicate.Build {
	return predicate.Build(sql.FieldContainsFold(FieldBuildURL, v))
}

// BuildUUIDEQ applies the EQ predicate on the "build_uuid" field.
func BuildUUIDEQ(v uuid.UUID) predicate.Build {
	return predicate.Build(sql.FieldEQ(FieldBuildUUID, v))
}

// BuildUUIDNEQ applies the NEQ predicate on the "build_uuid" field.
func BuildUUIDNEQ(v uuid.UUID) predicate.Build {
	return predicate.Build(sql.FieldNEQ(FieldBuildUUID, v))
}

// BuildUUIDIn applies the In predicate on the "build_uuid" field.
func BuildUUIDIn(vs ...uuid.UUID) predicate.Build {
	return predicate.Build(sql.FieldIn(FieldBuildUUID, vs...))
}

// BuildUUIDNotIn applies the NotIn predicate on the "build_uuid" field.
func BuildUUIDNotIn(vs ...uuid.UUID) predicate.Build {
	return predicate.Build(sql.FieldNotIn(FieldBuildUUID, vs...))
}

// BuildUUIDGT applies the GT predicate on the "build_uuid" field.
func BuildUUIDGT(v uuid.UUID) predicate.Build {
	return predicate.Build(sql.FieldGT(FieldBuildUUID, v))
}

// BuildUUIDGTE applies the GTE predicate on the "build_uuid" field.
func BuildUUIDGTE(v uuid.UUID) predicate.Build {
	return predicate.Build(sql.FieldGTE(FieldBuildUUID, v))
}

// BuildUUIDLT applies the LT predicate on the "build_uuid" field.
func BuildUUIDLT(v uuid.UUID) predicate.Build {
	return predicate.Build(sql.FieldLT(FieldBuildUUID, v))
}

// BuildUUIDLTE applies the LTE predicate on the "build_uuid" field.
func BuildUUIDLTE(v uuid.UUID) predicate.Build {
	return predicate.Build(sql.FieldLTE(FieldBuildUUID, v))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Build {
	return predicate.Build(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Build {
	return predicate.Build(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Build {
	return predicate.Build(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Build {
	return predicate.Build(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Build {
	return predicate.Build(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Build {
	return predicate.Build(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Build {
	return predicate.Build(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Build {
	return predicate.Build(sql.FieldLTE(FieldTimestamp, v))
}

// TimestampIsNil applies the IsNil predicate on the "timestamp" field.
func TimestampIsNil() predicate.Build {
	return predicate.Build(sql.FieldIsNull(FieldTimestamp))
}

// TimestampNotNil applies the NotNil predicate on the "timestamp" field.
func TimestampNotNil() predicate.Build {
	return predicate.Build(sql.FieldNotNull(FieldTimestamp))
}

// HasInvocations applies the HasEdge predicate on the "invocations" edge.
func HasInvocations() predicate.Build {
	return predicate.Build(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InvocationsTable, InvocationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInvocationsWith applies the HasEdge predicate on the "invocations" edge with a given conditions (other predicates).
func HasInvocationsWith(preds ...predicate.BazelInvocation) predicate.Build {
	return predicate.Build(func(s *sql.Selector) {
		step := newInvocationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Build) predicate.Build {
	return predicate.Build(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Build) predicate.Build {
	return predicate.Build(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Build) predicate.Build {
	return predicate.Build(sql.NotPredicates(p))
}
