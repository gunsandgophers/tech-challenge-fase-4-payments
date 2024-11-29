package app

import (
	"tech-challenge-fase-1/internal/infra/database"
	"tech-challenge-fase-1/internal/infra/events"
	httpserver "tech-challenge-fase-1/internal/infra/http"
	"tech-challenge-fase-1/internal/infra/repositories"
	"tech-challenge-fase-1/internal/infra/services"

	"github.com/gin-contrib/cors"
)

type APIApp struct {
	httpServer         *httpserver.GinHTTPServerAdapter
	connection         *database.PGXConnectionAdapter
	paymentRepository    *repositories.PaymentRepositoryDB
	mercadoPagoGateway *services.MercadoPagoGateway
	eventManager *events.EventManager
}

func NewAPIApp() *APIApp {
	app := &APIApp{}
	app.initGin()
	app.configCors()
	app.initConnectionDB()
	app.configRoutes()
	return app
}

func (app *APIApp) initGin() {
	app.httpServer = httpserver.NewGinHTTPServerAdapter()
	app.httpServer.SetTrustedProxies(nil)
}

func (app *APIApp) configCors() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"}
	app.httpServer.Engine.Use(cors.New(config))
}

func (app *APIApp) initConnectionDB() {
	app.connection = database.NewPGXConnectionAdapter()
	app.paymentRepository = repositories.NewOrderRepositoryDB(app.connection)
	app.eventManager = events.NewEventManager()
	app.mercadoPagoGateway = services.NewMercadoPagoGateway(app.eventManager)
}

func (app *APIApp) configRoutes() {
	registerRouters(app)
}

func (app *APIApp) Run() {
	app.httpServer.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (app *APIApp) Shutdown() {
	app.connection.Close()
}
