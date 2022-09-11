package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"checkoutpromo/graph/generated"
	"checkoutpromo/graph/model"
	"checkoutpromo/module/checkoutpromo/usecase"
	"context"
)

func (r *mutationResolver) AddCart(ctx context.Context, customerID string, productID string, qty int) (*model.ResponseData, error) {
	addCart, err := usecase.AddCart(customerID, productID, qty)

	if err != nil {
		return nil, err
	}

	return addCart, nil
}

func (r *mutationResolver) DeleteCart(ctx context.Context, customerID string, productID string) (*model.ResponseData, error) {
	delCart, err := usecase.DeleteCart(customerID, productID)

	if err != nil {
		return nil, err
	}

	return delCart, nil
}

func (r *mutationResolver) Checkout(ctx context.Context, customerID string) (*model.ResponseData, error) {
	orderNum, err := usecase.Checkout(customerID)

	if err != nil {
		return nil, err
	}

	hasil := &model.ResponseData{
		StatusCode: 200,
		Message:    "success checkout, order number : " + orderNum,
	}

	return hasil, nil
}

func (r *queryResolver) ProductList(ctx context.Context) ([]*model.Product, error) {
	productList, err := usecase.ProductList()

	if err != nil {
		return nil, err
	}

	return productList, nil
}

func (r *queryResolver) CustomerList(ctx context.Context) ([]*model.Customer, error) {
	customerList, err := usecase.CustomerList()

	if err != nil {
		return nil, err
	}

	return customerList, nil
}

func (r *queryResolver) ShoppingCart(ctx context.Context, customerID string) ([]*model.Cart, error) {
	cartList, err := usecase.CartList(customerID)

	if err != nil {
		return nil, err
	}

	return cartList, nil
}

func (r *queryResolver) OrderByCustomer(ctx context.Context, customerID string) ([]*model.Order, error) {
	orderList, err := usecase.OrderList(customerID)

	if err != nil {
		return nil, err
	}

	return orderList, nil
}

func (r *queryResolver) OrderDetail(ctx context.Context, orderNum string) ([]*model.OrderDetail, error) {
	orderDetail, err := usecase.OrderDetail(orderNum)

	if err != nil {
		return nil, err
	}

	return orderDetail, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
