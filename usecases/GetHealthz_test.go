package usecases_test

import (
	"context"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/adapters"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/usecases"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetHealthz", func() {
	var u adapters.Usecases
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
		u, _ = usecases.NewUsecases(ctx, mockPersistenceStore, configuration)
	})
	Context("whent the data is valid", func() {
		It("passes", func() {
			_, err := u.GetHealthz(ctx, "v1")
			Expect(err).To(BeNil())
		})
	})
})
