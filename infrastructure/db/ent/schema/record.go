package schema

import "entgo.io/ent"

// Record holds the schema definition for the Record entity.
type Record struct {
	ent.Schema
}

// Fields of the Record.
func (Record) Fields() []ent.Field {
	return nil
}

// Edges of the Record.
func (Record) Edges() []ent.Edge {
	return nil
}
