package routes

import (
	"base-api/app/controllers"
	"base-api/app/repositories"
	"base-api/app/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
)

func InitializeRoutes(DB *gorm.DB) {

	participantRepository := repositories.ParticipantRepository(DB)
	participantService := services.ParticipantService(participantRepository)
	participantController := controllers.ParticipantController(participantService)

	//gin config
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	//participant
	api.GET("/participant", participantController.GetParticipants)
	api.GET("/participant/:id", participantController.GetParticipant)
	api.POST("/participant", participantController.CreateParticipant)
	api.PUT("/participant/:id", participantController.UpdateParticipant)
	api.DELETE("/participant/:id", participantController.DeleteParticipant)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
