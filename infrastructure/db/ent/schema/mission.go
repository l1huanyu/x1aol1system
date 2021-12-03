package schema

import "entgo.io/ent"

// Mission holds the schema definition for the Mission entity.
type Mission struct {
	ent.Schema
}

// Fields of the Mission.
func (Mission) Fields() []ent.Field {
	return nil
}

// Edges of the Mission.
func (Mission) Edges() []ent.Edge {
	return nil
}
