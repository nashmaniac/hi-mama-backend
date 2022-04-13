package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/endpoints/app/v1/marshallers"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

func (v1 *apiV1) GetOngoing(c *gin.Context) {
	currentUser := c.GetString("current-user-id")
	if currentUser == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, &models.ErrorResponse{
			Status:  http.StatusForbidden,
			Message: "forbidden",
		})
		return
	}

	entry, err := v1.Usecases.FindOngoingTime(c, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &models.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if entry != nil {
		c.JSON(http.StatusOK, marshallers.MarshalEntry(*entry))
	} else {
		c.JSON(http.StatusOK, nil)
	}

}
