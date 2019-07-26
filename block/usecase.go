package block

// Usecase defines usecase contract
type Usecase interface {
	GetBlockChain() ([]Block, error)
	WriteBlock(bpm int) error
}
