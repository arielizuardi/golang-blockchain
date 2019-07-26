package server

import (
	"fmt"

	"github.com/arielizuardi/golang-blockchain/block"
	"github.com/arielizuardi/golang-blockchain/block/handler"
	"github.com/arielizuardi/golang-blockchain/block/repository"
	"github.com/arielizuardi/golang-blockchain/block/usecase"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

// Server represents http server
type Server struct {
	Port       int
	Blockchain []block.Block
}

// StartHTTPServer start http server
func (s *Server) StartHTTPServer() {
	e := echo.New()

	blockRepository := repository.NewInMemBlockRepository(s.Blockchain)
	blockUsecase := usecase.NewBlockUsecase(blockRepository)
	blockHandler := handler.NewBlockHTTPHandler(blockUsecase)

	e.GET("/blockchain", blockHandler.HandleGetBlockChain)
	logrus.Fatal(e.Start(fmt.Sprintf(":%d", s.Port)))
}
