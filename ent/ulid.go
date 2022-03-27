package ent

import (
	"context"
	"fmt"

	"github.com/tmortx7/gql-ent-go/ent/schema/ulid"
	"github.com/tmortx7/gql-ent-go/pkg/const/globalid"
)

var globalIDs = globalid.New()

// IDToType maps an ulid.ID to the underlaying table
func IDToType(_ context.Context, id ulid.ID) (string, error) {
	if len(id) < 4 {
		return "", fmt.Errorf("could not map id to type")
	}
	prefix := id[:4]
	t, err := globalIDs.FindTableByID(string(prefix))
	if err != nil {
		return "", fmt.Errorf("could not map id prefix '%s' to a type", prefix)
	}
	return t, nil
}
