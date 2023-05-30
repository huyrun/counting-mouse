package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/huyrun/counting-mouse/src/config"
	"github.com/huyrun/counting-mouse/src/usecase"
)

func ProvideAppConfig() (*config.AppConfig, error) {
	return config.NewAppConfig()
}

func ProvideUseCase() usecase.UseCase {
	return usecase.NewUseCase()
}

func ProvideEngine() *gin.Engine {
	return gin.New()
}
