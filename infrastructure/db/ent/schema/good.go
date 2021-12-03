package schema

import "entgo.io/ent"

// Good holds the schema definition for the Good entity.
type Good struct {
	ent.Schema
}

// Fields of the Good.
func (Good) Fields() []ent.Field {
	return nil
}

// Edges of the Good.
func (Good) Edges() []ent.Edge {
	return nil
}
