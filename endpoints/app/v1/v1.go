package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/adapters"
)

type ApiV1 interface {
	Healthz(c *gin.Context)
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
	Me(c *gin.Context)
	ClockIn(c *gin.Context)
	ClockOut(c *gin.Context)
	GetEntries(c *gin.Context)
	GetOngoing(c *gin.Context)
}

type apiV1 struct {
	Usecases adapters.Usecases
}

func V1Api(
	usecases adapters.Usecases,
) (ApiV1, error) {
	return &apiV1{
		Usecases: usecases,
	}, nil
}
