package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/tmortx7/gql-ent-go/ent/schema/ulid"
	"github.com/tmortx7/gql-ent-go/pkg/const/globalid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("role").
			NamedValues(
				"Client", "CLIENT",
				"Admin", "ADMIN",
			).
			Default("CLIENT"),
		field.String("id").
			GoType(ulid.ID("")).
			DefaultFunc(func() ulid.ID {
				return ulid.MustNew(globalid.New().User.Prefix)
			}),
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("email").NotEmpty().Unique(),
		field.String("password"),
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
