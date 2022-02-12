package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		// Id
		field.Uint64("id"),

		// Name
		field.String("name"),

		// slug
		field.String("slug"),

		// description
		field.String("description").
			Optional(),

		// Level
		field.Uint("level"),

		// Created at
		field.Time("created_at").
			Default(time.Now),

		// Updated at
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}

func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("slug").
			Unique(),
	}
}
