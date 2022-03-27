package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tmortx7/gql-ent-go/ent"
	"github.com/tmortx7/gql-ent-go/ent/schema/ulid"
	"github.com/tmortx7/gql-ent-go/graph/generated"
)

func (r *mutationResolver) CreatePet(ctx context.Context, input ent.CreatePetInput) (*ent.Pet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePet(ctx context.Context, input ent.UpdatePetInput) (*ent.Pet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *petResolver) CreatedAt(ctx context.Context, obj *ent.Pet) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *petResolver) UpdatedAt(ctx context.Context, obj *ent.Pet) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pet(ctx context.Context, id *ulid.ID) (*ent.Pet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pets(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.PetWhereInput) (*ent.PetConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Pet returns generated.PetResolver implementation.
func (r *Resolver) Pet() generated.PetResolver { return &petResolver{r} }

type petResolver struct{ *Resolver }
