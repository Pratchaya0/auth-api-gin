package servers

import (
	"log"

	"github.com/Pratchaya0/auth-api-gin/configs"
	"github.com/Pratchaya0/auth-api-gin/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	App *gin.Engine
	Cfg *configs.Configs
	Db  *gorm.DB
}

func NewServer(cfg *configs.Configs, db *gorm.DB) *Server {
	return &Server{
		App: gin.Default(),
		Cfg: cfg,
		Db:  db,
	}
}

func (s *Server) Start() {
	if err := s.MapHandlers(); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	connectdsn, err := utils.ConnectionUrlBuilder("gin", s.Cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	host := s.Cfg.App.Host
	port := s.Cfg.App.Port
	log.Printf("server has been started on %s:%s ⚡", host, port)

	if err := s.App.Run(connectdsn); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}
