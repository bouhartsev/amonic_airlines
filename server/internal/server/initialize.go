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
	gin.SetMode(gin.ReleaseMode)
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
	au.POST(`sign-out`, s.SignOut)
	au.POST(`report`, s.ReportLastLogoutError)

	users := api.Group(`users`)
	users.GET(``, s.GetUsers)
	users.POST(``, s.CreateUser)
	users.GET(`:user_id`, s.GetUser)
	users.PATCH(`:user_id`, s.UpdateUser)
	users.GET(`:user_id/logins`, s.GetUserLogins)
	users.POST(`:user_id/switch-status`, s.SwitchUserStatus)

	schedules := api.Group(`schedules`)
	schedules.GET(``, s.GetSchedules)
	schedules.PATCH(`:schedule_id`, s.UpdateSchedule)
	schedules.POST(`:schedule_id/switch-status`, s.SwitchScheduleStatus)
	schedules.POST(`/upload`, s.UpdateSchedulesFromFile)

	api.GET(`countries`, s.GetCountries)
	api.GET(`cabin-types`, s.GetCabinTypes)
	api.GET(`offices`, s.GetOffices)
	api.GET(`airports`, s.GetAirports)
	api.GET(`amenities`, s.GetAmenities)
	api.GET(`amenities/reports/brief`, s.GetAmenitiesBriefReport)

	tickets := api.Group(`tickets`)

	tickets.POST(``, s.AddTicket)
	tickets.GET(``, s.GetTickets)
	tickets.GET(`:ticket_id/amenities`, s.GetTicketAmenities)
	tickets.DELETE(`:ticket_id/amenities`, s.RemoveTicketAmenities)
	tickets.POST(`:ticket_id/amenities`, s.AddTicketAmenities)

	reviews := api.Group(`reviews`)

	reviews.POST(``, s.AddReview)
	reviews.GET(`brief`, s.GetBriefReviews)
	reviews.GET(`detailed`, s.GetDetailedReviews)

	api.GET(`report`, s.GetDetailedReport)

	// Documentation endpoint registration
	router.GET(`api/docs/*any`, ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
