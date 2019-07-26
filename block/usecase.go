package block

// Usecase defines usecase contract
type Usecase interface {
	GetBlockChain() ([]Block, error)
	WriteBlock(b Block) error
}
