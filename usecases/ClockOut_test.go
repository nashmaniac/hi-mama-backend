package usecases_test

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/adapters"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/usecases"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("ClockOut", func() {
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
			_, err = u.ClockOut(ctx, username)
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
			_, err = u.ClockOut(ctx, username)
			Expect(err).ToNot(BeNil())
		})
	})

	When("there is no timing present", func() {
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
				gorm.ErrRecordNotFound,
			).Times(1)

			_, err = u.ClockOut(ctx, username)
			Expect(err).ToNot(BeNil())
		})
	})

	When("there is no entry present", func() {
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
				nil,
			).Times(1)

			_, err = u.ClockOut(ctx, username)
			Expect(err).ToNot(BeNil())
		})
	})

	Context("there is existing timing present", func() {
		When("updating", func() {
			It("fails", func() {
				username := "test-username"
				userId := 1
				user := &models.User{
					Username: username,
					Model: gorm.Model{
						ID: uint(userId),
					},
				}

				entry := &models.Entry{
					UserID: user.ID,
					User:   *user,
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
					entry,
					nil,
				).Times(1)

				mockPersistenceStore.EXPECT().SaveEntry(
					ctx,
					gomock.Any(),
				).Return(
					nil,
					fmt.Errorf("db error"),
				).Times(1)

				_, err = u.ClockOut(ctx, username)
				Expect(err).ToNot(BeNil())
			})
		})

		When("saving is successful", func() {
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
				currentTime := time.Now()
				entry := &models.Entry{
					UserID:       user.ID,
					User:         *user,
					ClockOutTime: &currentTime,
				}
				mockPersistenceStore.EXPECT().FindOngoingTime(
					ctx,
					*user,
				).Return(
					entry,
					nil,
				).Times(1)

				mockPersistenceStore.EXPECT().SaveEntry(
					ctx,
					gomock.Any(),
				).Return(
					entry,
					nil,
				).Times(1)

				entry, err = u.ClockOut(ctx, username)
				Expect(err).To(BeNil())
				Expect(entry).ToNot(BeNil())
			})
		})

	})

})
