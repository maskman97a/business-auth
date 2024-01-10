package conf

import (
	"business-auth/conf/environment"
	"business-auth/conf/http_server"
	"business-auth/conf/server"
	"business-auth/internal/constants"
	"business-auth/internal/controller/controller_impl"
	"business-auth/internal/repository/repository_impl"
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App interface {
	Run()

	Init()

	AddController()
}

type app struct {
	env environment.Environment

	gin *gin.Engine

	gormDB *gorm.DB

	db *sql.DB

	httpServer server.HttpServer
}

func NewApp() App {
	return &app{}
}

func (app *app) Run() {

	app.httpServer.Start()
	defer app.httpServer.Stop()

	// Listen for OS signals to perform a graceful shutdown
	log.Println("listening signals...")
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	log.Println("graceful shutdown...")
}

func (app *app) Init() {
	app.env = environment.ConfigAppEnv()

	app.gin = gin.New()
	//gin.SetMode(gin.ReleaseMode)

	app.gin.Use(gin.Recovery())

	app.gormDB = repository_impl.NewGormDB(repository_impl.ProvideConfig())

	app.db, _ = repository_impl.NewSQLDB(app.gormDB)

	app.AddController()

	app.httpServer = server.NewHttpServer(
		app.gin,
		http_server.HttpServerConfig{
			Port: app.env.ServerPort,
		},
	)
}

func (app *app) AddController() {
	authController := controller_impl.NewAuthController(app.gin, app.gormDB)
	controller_impl.InitRouter(app.gin, constants.ContextPath, "/auth", "/sign-up", authController.SignUp, http.MethodPost)

	controller_impl.InitRouter(app.gin, constants.ContextPath, "/auth", "/login", authController.Login, http.MethodPost)

}
