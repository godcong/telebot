// Code generated by entc, DO NOT EDIT.

package command

const (
	// Label holds the string label denoting the command type in the database.
	Label = "command"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the command in the database.
	Table = "commands"
)

// Columns holds all SQL columns for command fields.
var Columns = []string{
	FieldID,
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
