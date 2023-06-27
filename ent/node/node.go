// Code generated by ent, DO NOT EDIT.

package node

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the node type in the database.
	Label = "node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// EdgePrev holds the string denoting the prev edge name in mutations.
	EdgePrev = "prev"
	// EdgeNext holds the string denoting the next edge name in mutations.
	EdgeNext = "next"
	// Table holds the table name of the node in the database.
	Table = "nodes"
	// PrevTable is the table that holds the prev relation/edge.
	PrevTable = "nodes"
	// PrevColumn is the table column denoting the prev relation/edge.
	PrevColumn = "node_next"
	// NextTable is the table that holds the next relation/edge.
	NextTable = "nodes"
	// NextColumn is the table column denoting the next relation/edge.
	NextColumn = "node_next"
)

// Columns holds all SQL columns for node fields.
var Columns = []string{
	FieldID,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "nodes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"node_next",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Node queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByPrevField orders the results by prev field.
func ByPrevField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPrevStep(), sql.OrderByField(field, opts...))
	}
}

// ByNextField orders the results by next field.
func ByNextField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNextStep(), sql.OrderByField(field, opts...))
	}
}
func newPrevStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, PrevTable, PrevColumn),
	)
}
func newNextStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, NextTable, NextColumn),
	)
}
