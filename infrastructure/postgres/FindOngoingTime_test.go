package postgres_test

import (
	"context"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FindOngoingTime", func() {
	When("there is not ongoing time", func() {
		It("should return empty", func() {
			ctx := context.Background()
			username := "test_user_name"
			password := "password"
			user := models.User{
				Password: password,
				Username: username,
			}
			stubRepository.CreateUser(ctx, &user)

			entry, err := stubRepository.FindOngoingTime(ctx, user)
			Expect(entry).To(BeNil())
			Expect(err).To(BeNil())

		})
	})
})
