package block

// Repository ...
type Repository interface {
	GetBlockChain() ([]Block, error)
	WriteBlock(b Block) error
}
