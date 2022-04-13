package usecases_test

import (
	"context"
	"fmt"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/adapters"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/usecases"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClockIn", func() {
	var ctx context.Context
	var u adapters.Usecases
	var username string
	var err error
	BeforeEach(func() {
		username = "shetu"
		ctx = context.Background()
		u, err = usecases.NewUsecases(ctx, mockPersistenceStore, configuration)
	})

	When("there is error in fetching from db", func() {
		It("fails", func() {
			mockPersistenceStore.EXPECT().FindUserByUsername(
				ctx,
				username,
			).Return(
				nil,
				fmt.Errorf("error in db"),
			).Times(1)
			_, err = u.ClockIn(ctx, username)
			Expect(err).ToNot(BeNil())
		})
	})

	When("there is no user found", func() {
		It("fails", func() {
			mockPersistenceStore.EXPECT().FindUserByUsername(
				ctx,
				username,
			).Return(
				nil,
				fmt.Errorf("user not found"),
			).Times(1)
			_, err = u.ClockIn(ctx, username)
			Expect(err).ToNot(BeNil())
		})
	})
})
