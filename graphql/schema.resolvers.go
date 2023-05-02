package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"commitado/database"
	"commitado/graphql/generated"
	"commitado/graphql/model"
	"context"
)

// CreateBill is the resolver for the createBill field.
func (r *mutationResolver) CreateBill(ctx context.Context, input model.CreateBillInput) (*model.Bill, error) {
	return db.CreateBill(input), nil
}

// UpdateBill is the resolver for the updateBill field.
func (r *mutationResolver) UpdateBill(ctx context.Context, id string, input model.UpdateBillInput) (*model.Bill, error) {
	return db.UpdateBill(id, input), nil
}

// DeleteBill is the resolver for the deleteBill field.
func (r *mutationResolver) DeleteBill(ctx context.Context, id string) (*model.DeleteBillResponse, error) {
	return db.DeleteBill(id), nil
}

// Bills is the resolver for the bills field.
func (r *queryResolver) Bills(ctx context.Context) ([]*model.Bill, error) {
	return db.GetBills(), nil
}

// Bill is the resolver for the bill field.
func (r *queryResolver) Bill(ctx context.Context, id string) (*model.Bill, error) {
	return db.GetBill(id), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()
