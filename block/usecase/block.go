package usecase

import (
	"errors"

	"github.com/arielizuardi/golang-blockchain/block"
)

// BlockUsecase ...
type BlockUsecase struct {
	Repository block.Repository
}

// NewBlockUsecase returns new instance of block usecase
func NewBlockUsecase(repository block.Repository) block.Usecase {
	return &BlockUsecase{repository}
}

func (u *BlockUsecase) GetBlockChain() ([]block.Block, error) {
	return u.Repository.GetBlockChain()
}

func (u *BlockUsecase) WriteBlock(bpm int) error {
	blockchain, err := u.GetBlockChain()
	if err != nil {
		return err
	}

	lastBlock := blockchain[len(blockchain)-1]

	newBlock, err := block.GenerateBlock(lastBlock, bpm)
	if err != nil {
		return err
	}

	if !block.IsBlockValid(newBlock, lastBlock) {
		return errors.New("not valid")
	}

	return u.Repository.WriteBlock(newBlock)
}
