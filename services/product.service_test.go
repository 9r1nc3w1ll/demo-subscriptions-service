package services

import (
	"context"
	"database/sql"
	"fmt"
	"lithium-test/db/models"
	"lithium-test/pb"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	dbConfig := struct {
		User     string
		Password string
		Database string
	}{
		User:     "root",
		Password: "root",
		Database: "test_db",
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("could not connect to Docker: %s", err)
	}

	resource, err := pool.Run("postgres", "latest", []string{
		fmt.Sprintf("POSTGRES_DB=%s", dbConfig.Database),
		fmt.Sprintf("POSTGRES_USER=%s", dbConfig.User),
		fmt.Sprintf("POSTGRES_PASSWORD=%s", dbConfig.Password),
	})
	if err != nil {
		log.Fatalf("could not start resource: %s", err)
	}

	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("could not purge resource: %s", err)
		}
	}()

	pgPort := resource.GetPort("5432/tcp")

	dsn := fmt.Sprintf("host=localhost user=%s password=root dbname=test_db port=%s sslmode=disable TimeZone=UTC", dbConfig.User, pgPort)
	if err := pool.Retry(func() error {
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		log.Fatalf("could not connect to database: %s", err)
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to db: %s", err)
	}

	db.AutoMigrate(models.Product{})
	db.AutoMigrate(models.SubscriptionPlan{})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestCreateProduct(t *testing.T) {
	t.Run("Successfully create product", func(t *testing.T) {
		ctx := context.Background()

		md := metadata.New(map[string]string{
			"authorization": "Bearer VALID_TEST_TOKEN",
		})
		ctx = metadata.NewIncomingContext(ctx, md)
		service := NewProductService(db)

		response, err := service.CreateProduct(ctx, &pb.CreateProductInput{
			Name:        "Product 1",
			Type:        "PHYSICAL",
			Price:       10.50,
			Description: "Description of product 1",
		})

		assert.Nil(t, err)
		assert.NotEmpty(t, response)

		t.Cleanup(func() {
			service.DeleteProduct(ctx, &pb.DeleteProductInput{
				Id: response.Id,
			})
		})
	})
}

func TestGetProduct(t *testing.T) {
	ctx := context.Background()

	md := metadata.New(map[string]string{
		"authorization": "Bearer VALID_TEST_TOKEN",
	})
	ctx = metadata.NewIncomingContext(ctx, md)

	service := NewProductService(db)
	data, err := service.CreateProduct(ctx, &pb.CreateProductInput{
		Name:        "Product 1",
		Type:        "PHYSICAL",
		Price:       10.50,
		Description: "Description of product 1",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, data)

	t.Cleanup(func() {
		service.DeleteProduct(ctx, &pb.DeleteProductInput{
			Id: data.Id,
		})
	})

	t.Run("Successfully fetch an existing product", func(t *testing.T) {
		response, err := service.GetProduct(ctx, &pb.GetProductInput{
			Id: data.Id,
		})
		assert.Nil(t, err)
		assert.Equal(t, response.Id, data.Id)
	})

	t.Run("Fail fetch non-existing product", func(t *testing.T) {
		_, err := service.GetProduct(ctx, &pb.GetProductInput{
			Id: 20,
		})
		assert.NotNil(t, err)
	})
}

func TestListProducts(t *testing.T) {
	ctx := context.Background()

	md := metadata.New(map[string]string{
		"authorization": "Bearer VALID_TEST_TOKEN",
	})
	ctx = metadata.NewIncomingContext(ctx, md)

	service := NewProductService(db)
	data, err := service.CreateProduct(ctx, &pb.CreateProductInput{
		Name:        "Product 1",
		Type:        "PHYSICAL",
		Price:       10.50,
		Description: "Description of product 1",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, data)

	t.Cleanup(func() {
		service.DeleteProduct(ctx, &pb.DeleteProductInput{
			Id: data.Id,
		})
	})

	t.Run("Successfully return all existing products", func(t *testing.T) {
		response, err := service.ListProducts(ctx, &emptypb.Empty{})
		assert.Nil(t, err)
		assert.Equal(t, len(response.Data), 1)
	})
}

func TestDeleteProduct(t *testing.T) {
	ctx := context.Background()

	md := metadata.New(map[string]string{
		"authorization": "Bearer VALID_TEST_TOKEN",
	})
	ctx = metadata.NewIncomingContext(ctx, md)

	service := NewProductService(db)
	response, err := service.CreateProduct(ctx, &pb.CreateProductInput{
		Name:        "Product 1",
		Type:        "PHYSICAL",
		Price:       10.50,
		Description: "Description of product 1",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, response)

	t.Cleanup(func() {
		service.DeleteProduct(ctx, &pb.DeleteProductInput{
			Id: response.Id,
		})
	})

	t.Run("Successfully delete an existing product", func(t *testing.T) {
		_, err := service.DeleteProduct(ctx, &pb.DeleteProductInput{
			Id: response.Id,
		})
		assert.Nil(t, err)
	})

	t.Run("Fail delete non-existing product", func(t *testing.T) {
		_, err := service.DeleteProduct(ctx, &pb.DeleteProductInput{
			Id: response.Id,
		})
		assert.NotNil(t, err)
	})
}
