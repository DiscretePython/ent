// Code generated by ent, DO NOT EDIT.

package petusername

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the petusername type in the database.
	Label = "pet_user_name"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the petusername in the database.
	Table = "pet_user_names"
)

// Columns holds all SQL columns for petusername fields.
var Columns = []string{
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the PetUserName queries.
type OrderOption func(*sql.Selector)

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}