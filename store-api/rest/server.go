package rest

import (
	"blog-api/rest/middleware"
	"blog-api/rest/req"
	"blog-api/service"
	"blog-api/tools/tokenmanager"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	server *http.Server
	router chi.Router

	tokenManager *tokenmanager.Tool
	service      *service.Service
}

type Config struct {
	Addr string
	Port string

	TokenManager *tokenmanager.Tool
	Service      *service.Service
}

func NewServer(config *Config) *Server {
	return &Server{
		server: &http.Server{
			Addr:    config.Addr + ":" + config.Port,
			Handler: http.NotFoundHandler(),
		},
		router: chi.NewRouter(),

		tokenManager: config.TokenManager,
		service:      config.Service,
	}
}

func (s *Server) RunServer() error {
	return s.server.ListenAndServe()
}

func (s *Server) SetupRouter() {
	// setup cors
	s.setupCors()

	// middleware
	mw := middleware.New(s.service, s.tokenManager)

	s.router.Route("/auth", func(r chi.Router) {
		r.Method("POST", "/signup", req.NewHandler(s.service.User.Signup))
		r.Method("POST", "/signin", req.NewHandler(s.service.User.Signin))
		r.Method("POST", "/refresh", req.NewHandler(s.service.User.Refresh))
	})

	s.router.Route("/user", func(r chi.Router) {
		r.Use(mw.Auth)
		r.Method("GET", "/check", req.NewHandler(s.service.User.Check))
		r.Method("POST", "/upload", req.NewHandler(s.service.Product.UploadImage))
	})

	s.router.Route("/assets", func(r chi.Router) {
		r.Method("GET", "/{file}", req.NewHandler(s.service.Product.SendFile))
	})

	s.router.Route("/product", func(r chi.Router) {
		r.Use(mw.Auth)
		r.Method("POST", "/new", req.NewHandler(s.service.Product.NewProduct))
	})

	s.router.Route("/purchase", func(r chi.Router) {
		r.Use(mw.Auth)
		r.Method("GET", "/generate", req.NewHandler(s.service.Purchase.GenerateBill))
	})

	s.router.Group(func(r chi.Router) {
		r.Method("GET", "/feed_content", req.NewHandler(s.service.Product.GetProductForFeed))
		r.Method("GET", "/search_product", req.NewHandler(s.service.Product.SearchProduct))
		r.Method("GET", "/product/{path}", req.NewHandler(s.service.Product.GetProduct))
		r.Method("GET", "/user/by_username", req.NewHandler(s.service.User.GetByUsername))
		r.Method("GET", "/purchase/success", req.NewHandler(s.service.Purchase.SuccessPurchase))
	})

	s.server.Handler = s.router
}

func (s *Server) setupCors() {
	s.router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}).Handler)
}
