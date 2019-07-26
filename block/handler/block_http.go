package handler

import (
	"net/http"

	"github.com/arielizuardi/golang-blockchain/block"
	"github.com/labstack/echo"
)

// BlockHTTPHandler represents block http handler
type BlockHTTPHandler struct {
	Usecase block.Usecase
}

// NewBlockHTTPHandler returns new instance of block http handler
func NewBlockHTTPHandler(usecase block.Usecase) *BlockHTTPHandler {
	return &BlockHTTPHandler{usecase}
}

func (h *BlockHTTPHandler) HandleGetBlockChain(c echo.Context) error {
	blockchain, err := h.Usecase.GetBlockChain()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, blockchain)
}
