package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/yerlanov/xmexercise/internal/company/handler"
	"github.com/yerlanov/xmexercise/internal/company/ipapi"
	"github.com/yerlanov/xmexercise/internal/company/service"
	"github.com/yerlanov/xmexercise/internal/company/storage"
	"github.com/yerlanov/xmexercise/internal/config"
	"log"
)

//func RunApp(cfg *config.Config) {
//
//	gin.SetMode(gin.ReleaseMode)
//	Router := gin.Default()
//
//	pg, err := postgresql.NewClient(context.TODO(), 3, cfg.DB.URI)
//	if err != nil {
//		log.Fatal("failed to connect: ", err)
//	}
//
//	store := storage.NewStore(pg)
//
//	service := company.NewService(store)
//
//	client := ipapi.NewClient(cfg.IpApiURL)
//
//	handler := company.Handler{Service: service, IpApi: client}
//
//	handler.Register(Router.Group("/api/v1/company"))
//
//	log.Println("start application")
//	start(Router, cfg)
//}

//func start(Router *gin.Engine, cfg *config.Config) {
//	var server *http.Server
//	var listener net.Listener
//
//	log.Printf("bind application to host: %s and port: %s", cfg.Listen.BindIP, cfg.Listen.Port)
//
//	var err error
//
//	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	server = &http.Server{
//		Handler:      Router,
//		WriteTimeout: 15 * time.Second,
//		ReadTimeout:  15 * time.Second,
//	}
//
//	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM},
//		server)
//
//	log.Println("application initialized and started")
//
//	if err = server.Serve(listener); err != nil {
//		log.Fatal(err)
//	}
//}

type Server struct {
	config *config.Config
	Router *gin.Engine
	Store  storage.Store
}

func NewServer(cfg *config.Config, store storage.Store) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	service := service.NewService(store)

	client := ipapi.NewClient(cfg.IpApiURL)

	handler := handler.Handler{Service: service, IpApi: client}

	handler.Register(router.Group("/api/v1/company"))

	log.Println("start application")

	server := &Server{
		config: cfg,
		Router: router,
	}
	return server
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
