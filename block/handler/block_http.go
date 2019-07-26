package handler

import (
	"net/http"

	"github.com/arielizuardi/golang-blockchain/block"
	"github.com/davecgh/go-spew/spew"
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

func (h *BlockHTTPHandler) HandleWriteBlockChain(c echo.Context) error {
	var m map[string]int
	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": err.Error(),
		})
	}

	if err := h.Usecase.WriteBlock(m["bpm"]); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	blockchain, err := h.Usecase.GetBlockChain()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	spew.Dump(blockchain)

	return c.JSON(http.StatusCreated, map[string]bool{
		"success": true,
	})
}
