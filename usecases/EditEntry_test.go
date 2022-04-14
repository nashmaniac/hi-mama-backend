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

var _ = Describe("Edit Entry", func() {
	var ctx context.Context
	var u adapters.Usecases
	var err error
	BeforeEach(func() {
		ctx = context.Background()
		u, err = usecases.NewUsecases(ctx, mockPersistenceStore, configuration)
	})

	When("fetching from db fails", func() {
		It("fails", func() {
			id := uint(1)
			currentTime := time.Now()
			clockInTime := currentTime
			clockOutTime := currentTime

			mockPersistenceStore.EXPECT().FindEntryById(
				ctx,
				id,
			).Return(
				nil,
				fmt.Errorf("db error"),
			).Times(1)
			_, err = u.EditEntry(
				ctx,
				id,
				clockInTime,
				&clockOutTime,
			)
			Expect(err).ToNot(BeNil())
		})
	})

	When("there is no entry with id", func() {
		It("fails", func() {
			id := uint(1)
			currentTime := time.Now()
			clockInTime := currentTime
			clockOutTime := currentTime

			mockPersistenceStore.EXPECT().FindEntryById(
				ctx,
				id,
			).Return(
				nil,
				nil,
			).Times(1)
			_, err = u.EditEntry(
				ctx,
				id,
				clockInTime,
				&clockOutTime,
			)
			Expect(err).ToNot(BeNil())
		})
	})

	Context("when there is entry", func() {
		When("updating fails in db", func() {
			It("fails", func() {
				id := uint(1)
				currentTime := time.Now()
				clockInTime := currentTime
				clockOutTime := currentTime

				entry := &models.Entry{
					Model: gorm.Model{
						ID: id,
					},
					UserID:       id,
					ClockInTime:  clockInTime,
					ClockOutTime: &clockOutTime,
				}

				mockPersistenceStore.EXPECT().FindEntryById(
					ctx,
					id,
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

				_, err = u.EditEntry(
					ctx,
					id,
					clockInTime,
					&clockOutTime,
				)
				Expect(err).ToNot(BeNil())
			})
		})

		When("updating passes in db", func() {
			It("passes", func() {
				id := uint(1)
				currentTime := time.Now()
				clockInTime := currentTime
				clockOutTime := currentTime

				entry := &models.Entry{
					Model: gorm.Model{
						ID: id,
					},
					UserID:       id,
					ClockInTime:  clockInTime,
					ClockOutTime: &clockOutTime,
				}

				mockPersistenceStore.EXPECT().FindEntryById(
					ctx,
					id,
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

				e, err := u.EditEntry(
					ctx,
					id,
					clockInTime,
					&clockOutTime,
				)
				Expect(err).To(BeNil())
				Expect(e).ToNot(BeNil())
			})
		})

	})

})
