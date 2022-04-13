package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/endpoints/app/v1/marshallers"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

type editEntryInput struct {
	ClockInTime  time.Time  `json:"clock_in"`
	ClockOutTime *time.Time `json:"clock_out"`
}

func (v1 *apiV1) EditEntry(
	c *gin.Context,
) {
	id, ok := c.Params.Get("id")
	if !ok || id == "" {
		c.AbortWithError(http.StatusBadRequest, &models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid data",
		})
		return
	}

	var userInput editEntryInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.AbortWithError(http.StatusBadRequest, &models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid data",
		})
		return
	}

	uid, _ := strconv.ParseUint(id, 10, 8)
	entry, err := v1.Usecases.EditEntry(c, uint(uid), userInput.ClockInTime, userInput.ClockOutTime)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, &models.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, marshallers.MarshalEntry(*entry))
}
