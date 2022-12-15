package server

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/bouhartsev/amonic_airlines/server/docs"
)

func (s *Server) initRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(
		cors.Middleware(cors.Config{
			Origins:        "*",
			Methods:        "GET, PATCH, POST, DELETE",
			RequestHeaders: "Origin, Authorization, Content-Type",
			MaxAge:         50 * time.Second,
		}),
		gin.Logger(),
		getTokenMiddleware(),
		s.checkAuthorizationMiddleware(),
	)

	api := router.Group(`/api`)

	au := api.Group(`auth`)
	au.POST(`sign-in`, s.SignIn)

	users := api.Group(`users`)
	users.GET(``, s.GetUsers)
	users.POST(``, s.CreateUser)
	users.GET(`:user_id`, s.GetUser)
	users.PATCH(`:user_id`, s.UpdateUser)

	schedules := api.Group(`schedules`)
	schedules.GET(``, s.GetSchedules)
	schedules.PATCH(`:schedule_id`, s.UpdateSchedule)
	schedules.POST(`:schedule_id/confirm`, s.ConfirmSchedule)
	schedules.POST(`:schedule_id/unconfirm`, s.UnconfirmSchedule)

	api.GET(`countries`, s.GetCountries)
	api.GET(`cabin-types`, s.GetCabinTypes)
	api.GET(`offices`, s.GetOffices)
	api.GET(`airports`, s.GetAirports)

	api.POST(`tickets`, s.AddTicket)

	reviews := api.Group(`reviews`)

	reviews.POST(``, s.AddReview)
	reviews.GET(`brief`, s.GetBriefReviews)
	reviews.GET(`detailed`, s.GetDetailedReviews)

	// Documentation endpoint registration
	router.GET(`api/docs/*any`, ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
