package services

import (
	"context"
	"fmt"
	"lithium-test/pb"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestCreateSubscriptionPlan(t *testing.T) {
	t.Run("Successfully create product", func(t *testing.T) {
		ctx := context.Background()
		productService := NewProductService(db)
		subscriptionService := NewSubscriptionService(db)

		createProductResponse, err := productService.CreateProduct(ctx, &pb.CreateProductInput{
			Name:        "Product 1",
			Type:        "SUBSCRIPTION",
			Price:       10.50,
			Description: "Description of product 1",
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, createProductResponse)

		t.Cleanup(func() {
			productService.DeleteProduct(ctx, &pb.DeleteProductInput{
				Id: createProductResponse.Id,
			})
		})

		createSubscriptionPlanResponse, err := subscriptionService.CreateSubscriptionPlan(ctx, &pb.CreateSubscriptionPlanInput{
			Name:        "Subscription 1",
			ProductId:   createProductResponse.Id,
			Duration:    30, // TODO: duration should come from product
			Price:       createProductResponse.Price,
			Description: fmt.Sprintf("Subscription for %s.", createProductResponse.Name),
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, createSubscriptionPlanResponse)

		t.Cleanup(func() {
			subscriptionService.DeleteSubscriptionPlan(ctx, &pb.DeleteSubscriptionPlanInput{
				Id: createSubscriptionPlanResponse.Id,
			})
		})
	})
}

func TestGetSubscriptionPlan(t *testing.T) {
	ctx := context.Background()
	productService := NewProductService(db)
	subscriptionService := NewSubscriptionService(db)

	createProductResponse, err := productService.CreateProduct(ctx, &pb.CreateProductInput{
		Name:        "Product 1",
		Type:        "SUBSCRIPTION",
		Price:       10.50,
		Description: "Description of product 1",
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, createProductResponse)

	t.Cleanup(func() {
		productService.DeleteProduct(ctx, &pb.DeleteProductInput{
			Id: createProductResponse.Id,
		})
	})

	createSubscriptionPlanResponse, err := subscriptionService.CreateSubscriptionPlan(ctx, &pb.CreateSubscriptionPlanInput{
		Name:        "Subscription 1",
		ProductId:   createProductResponse.Id,
		Duration:    30, // TODO: duration should come from product
		Price:       createProductResponse.Price,
		Description: fmt.Sprintf("Subscription for %s.", createProductResponse.Name),
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, createSubscriptionPlanResponse)

	t.Cleanup(func() {
		subscriptionService.DeleteSubscriptionPlan(ctx, &pb.DeleteSubscriptionPlanInput{
			Id: createSubscriptionPlanResponse.Id,
		})
	})

	t.Run("Successfully fetch an existing subscription plan", func(t *testing.T) {
		response, err := subscriptionService.GetSubscriptionPlan(ctx, &pb.GetSubscriptionPlanInput{
			Id: createSubscriptionPlanResponse.Id,
		})
		assert.Nil(t, err)
		assert.Equal(t, response.Id, createSubscriptionPlanResponse.Id)
	})

	t.Run("Fail fetch non-existing subscription plan", func(t *testing.T) {
		response, err := subscriptionService.GetSubscriptionPlan(ctx, &pb.GetSubscriptionPlanInput{
			Id: 20,
		})
		assert.NotNil(t, err)
		assert.Nil(t, response)
	})
}

func TestGetSubscriptionPlans(t *testing.T) {
	ctx := context.Background()
	productService := NewProductService(db)
	subscriptionService := NewSubscriptionService(db)

	createProductResponse, err := productService.CreateProduct(ctx, &pb.CreateProductInput{
		Name:        "Product 1",
		Type:        "SUBSCRIPTION",
		Price:       10.50,
		Description: "Description of product 1",
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, createProductResponse)

	t.Cleanup(func() {
		productService.DeleteProduct(ctx, &pb.DeleteProductInput{
			Id: createProductResponse.Id,
		})
	})

	createSubscriptionPlanResponse, err := subscriptionService.CreateSubscriptionPlan(ctx, &pb.CreateSubscriptionPlanInput{
		Name:        "Subscription 1",
		ProductId:   createProductResponse.Id,
		Duration:    30, // TODO: duration should come from product
		Price:       createProductResponse.Price,
		Description: fmt.Sprintf("Subscription for %s.", createProductResponse.Name),
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, createSubscriptionPlanResponse)

	t.Cleanup(func() {
		subscriptionService.DeleteSubscriptionPlan(ctx, &pb.DeleteSubscriptionPlanInput{
			Id: createSubscriptionPlanResponse.Id,
		})
	})

	t.Run("Successfully return all existing subscription plans", func(t *testing.T) {
		response, err := subscriptionService.ListSubscriptionPlans(ctx, &emptypb.Empty{})
		assert.Nil(t, err)
		assert.Equal(t, len(response.Data), 1)
	})
}

func TestDeleteSubscriptionPlan(t *testing.T) {
	ctx := context.Background()
	productService := NewProductService(db)
	subscriptionService := NewSubscriptionService(db)

	createProductResponse, err := productService.CreateProduct(ctx, &pb.CreateProductInput{
		Name:        "Product 1",
		Type:        "SUBSCRIPTION",
		Price:       10.50,
		Description: "Description of product 1",
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, createProductResponse)

	t.Cleanup(func() {
		productService.DeleteProduct(ctx, &pb.DeleteProductInput{
			Id: createProductResponse.Id,
		})
	})

	createSubscriptionPlanResponse, err := subscriptionService.CreateSubscriptionPlan(ctx, &pb.CreateSubscriptionPlanInput{
		Name:        "Subscription 1",
		ProductId:   createProductResponse.Id,
		Duration:    30, // TODO: duration should come from product
		Price:       createProductResponse.Price,
		Description: fmt.Sprintf("Subscription for %s.", createProductResponse.Name),
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, createSubscriptionPlanResponse)

	t.Cleanup(func() {
		subscriptionService.DeleteSubscriptionPlan(ctx, &pb.DeleteSubscriptionPlanInput{
			Id: createSubscriptionPlanResponse.Id,
		})
	})

	t.Run("Successfully delete an existing subscription plan", func(t *testing.T) {
		_, err := subscriptionService.DeleteSubscriptionPlan(ctx, &pb.DeleteSubscriptionPlanInput{
			Id: createSubscriptionPlanResponse.Id,
		})
		assert.Nil(t, err)
	})

	t.Run("Fail to delete a non-existing subscription plan", func(t *testing.T) {
		_, err := subscriptionService.DeleteSubscriptionPlan(ctx, &pb.DeleteSubscriptionPlanInput{
			Id: createSubscriptionPlanResponse.Id,
		})
		assert.NotNil(t, err)
	})
}
