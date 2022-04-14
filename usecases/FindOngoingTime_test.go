package usecases_test

import (
	"context"
	"fmt"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/adapters"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/usecases"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("FindOngoingTime", func() {
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
			_, err = u.FindOngoingTime(ctx, username)
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
			_, err = u.FindOngoingTime(ctx, username)
			Expect(err).ToNot(BeNil())
		})
	})

	When("there is error in fetching entry", func() {
		It("fails", func() {
			username := "test-username"
			userId := 1
			user := &models.User{
				Username: username,
				Model: gorm.Model{
					ID: uint(userId),
				},
			}
			mockPersistenceStore.EXPECT().FindUserByUsername(
				ctx,
				username,
			).Return(
				user,
				nil,
			).Times(1)

			mockPersistenceStore.EXPECT().FindOngoingTime(
				ctx,
				*user,
			).Return(
				nil,
				fmt.Errorf("db error"),
			).Times(1)

			_, err = u.FindOngoingTime(ctx, username)
			Expect(err).ToNot(BeNil())
		})
	})

	When("there is entry in db", func() {
		It("passes", func() {
			username := "test-username"
			userId := 1
			user := &models.User{
				Username: username,
				Model: gorm.Model{
					ID: uint(userId),
				},
			}
			mockPersistenceStore.EXPECT().FindUserByUsername(
				ctx,
				username,
			).Return(
				user,
				nil,
			).Times(1)

			mockPersistenceStore.EXPECT().FindOngoingTime(
				ctx,
				*user,
			).Return(
				&models.Entry{
					User: *user,
				},
				nil,
			).Times(1)

			e, err := u.FindOngoingTime(ctx, username)
			Expect(err).To(BeNil())
			Expect(e).ToNot(BeNil())
		})
	})
})
