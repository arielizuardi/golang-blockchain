package repository

import (
	"github.com/arielizuardi/golang-blockchain/block"
)

type inMemBlockRepository struct {
	blockchain []block.Block
}

// NewInMemBlockRepository returns new isntance of block Repository
func NewInMemBlockRepository(blockchain []block.Block) block.Repository {
	return &inMemBlockRepository{blockchain}
}

func (r *inMemBlockRepository) GetBlockChain() ([]block.Block, error) {
	return r.blockchain, nil
}

func (r *inMemBlockRepository) WriteBlock(b block.Block) error {
	r.blockchain = append(r.blockchain, b)
	return nil
}
