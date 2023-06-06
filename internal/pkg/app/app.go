package app

import (
	"log"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rostis232/psycho/internal/app/handler"
	"github.com/rostis232/psycho/internal/app/repository"
	"github.com/rostis232/psycho/internal/app/service"
	"github.com/rostis232/psycho/internal/models"
)

type App struct {
	handler *handler.Handler
	service *service.Service
	repo    *repository.Repository
	server  *echo.Echo
}

func NewApp(dbconf repository.Config) *App {
	a := &App{}
	db, err := repository.NewPostgresDB(dbconf)
	if err != nil {
		log.Fatalln(err)
	}
	a.repo = repository.NewRepository(db)
	a.service = service.NewService(*a.repo)
	a.handler = handler.NewHandler(a.service)
	a.server = echo.New()
	a.server.Use(middleware.Logger())
	a.server.Use(middleware.Recover())
	a.server.Use(middleware.Static("./static"))
	a.server.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	a.server.Renderer = service.Tmps
	if err != nil {
		log.Println(err)
	}

	a.server.GET(models.Main.URI, a.handler.Home)
	a.server.GET(models.Login.URI, a.handler.LogInGet)
	a.server.POST(models.Login.URI, a.handler.LogInPost)
	a.server.GET(models.Help.URI, a.handler.Help)
	a.server.GET(models.Clients.URI, a.handler.Clients)
	a.server.GET(models.Journal.URI, a.handler.Journal)
	a.server.GET(models.Profile.URI, a.handler.Profile)
	a.server.GET(models.Logout.URI, a.handler.Logout)
	a.server.GET(models.Clients.URI+"/:id", a.handler.BeneficiaryPage)
	a.server.GET("/register", a.handler.RegistrationGET)
	a.server.POST("/register", a.handler.RegistrationPOST)
	return a
}

func (a *App) Run(port string) {
	err := a.server.Start(":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Server is running")
}
