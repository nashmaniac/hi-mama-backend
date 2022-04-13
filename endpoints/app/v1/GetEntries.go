package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/endpoints/app/v1/marshallers"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

func (v1 *apiV1) GetEntries(c *gin.Context) {
	currentUser := c.GetString("current-user-id")
	if currentUser == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, &models.ErrorResponse{
			Status:  http.StatusForbidden,
			Message: "forbidden",
		})
		return
	}

	entries, err := v1.Usecases.GetTimeList(c, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &models.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, marshallers.MarshalEntries(entries))
}
