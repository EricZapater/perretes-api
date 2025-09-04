package server

import (
	"database/sql"
	"perretes-api/config"
	"perretes-api/internal/auth"
	"perretes-api/internal/courses"
	"perretes-api/internal/customers"
	"perretes-api/internal/health"
	"perretes-api/internal/users"
	"perretes-api/middleware"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	cfg    *config.Config
	db     *sql.DB
}

func NewServer(cfg *config.Config, db *sql.DB) *Server {
	return &Server{
		router: gin.Default(),
		cfg:    cfg,
		db:     db,
	}
}

func (s *Server) Setup() error {
	// CORS middleware
	s.router.Use(middleware.SetupCORS())
	

	// JWT middleware
	authMiddleware, err := middleware.SetupJWT(s.cfg)
	if err != nil {
		return err
	}

	// Action log middleware
	actionLogMiddleware := middleware.NewActionLogMiddleware(s.db)
	
	// Inicialitzar repositoris
	userRepo := users.NewUserRepository(s.db)
	customerRepo := customers.NewCustomerRepository(s.db)
	coursesRepo := courses.NewCourseRepository(s.db)

	// Inicialitzar serveis
	userService := users.NewUserService(userRepo)
	authService := auth.NewAuthService(userRepo, authMiddleware)
	customerService := customers.NewCustomerService(customerRepo, userService)
	coursesService := courses.NewCourseService(coursesRepo)



	// Inicialitzar handlers
	userHandler := users.NewUserHandler(userService)
	authHandler := auth.NewAuthHandler(authService, authMiddleware)
	customerHandler := customers.NewCustomerHandler(customerService)
	coursesHandler := courses.NewCourseHandler(coursesService)


	
	// Configurar les rutes públiques (sense autenticació)
	public := s.router.Group("/auth")
	public.Use(actionLogMiddleware.LogAction())
	public.GET("/health", health.CheckHealth)
	users.RegisterPublicRoutes(public, userHandler)
	auth.RegisterRoutes(public, authHandler, authMiddleware)


	// Configurar les rutes protegides (amb autenticació JWT)
	protected := s.router.Group("/api")
	protected.Use(authMiddleware.MiddlewareFunc())
	protected.Use(actionLogMiddleware.LogAction())
	

	// Registrar les rutes protegides
	users.RegisterRoutes(protected, userHandler)
	customers.RegisterRoutes(protected, customerHandler)
	courses.RegisterRoutes(protected, coursesHandler)

	
	return nil
}

func (s *Server) Run() error {
	//return s.router.RunTLS(":" + s.cfg.ApiPort, "./certs/cert.pem", "./certs/key.pem")
	return s.router.Run(":" + s.cfg.ApiPort)
}