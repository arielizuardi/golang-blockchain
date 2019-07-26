package usecase

import "github.com/arielizuardi/golang-blockchain/block"

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

func (u *BlockUsecase) WriteBlock(b block.Block) error {
	return u.Repository.WriteBlock(b)
}
