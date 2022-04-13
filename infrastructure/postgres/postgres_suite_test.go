package postgres_test

import (
	"context"
	"testing"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/adapters"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/infrastructure/postgres"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var stubRepository adapters.PeristenceStore

func TestInfrastructure(t *testing.T) {
	RegisterFailHandler(Fail)
	ctx := context.Background()
	stubRepository, _ = postgres.NewRepository(
		ctx,
		"postgres",
		"password",
		"testdb",
		"localhost",
		"5432",
	)
	RunSpecs(t, "Infrastructure Suite")
}
