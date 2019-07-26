package block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block represents single block in blockhain
type Block struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	BPM       int    `json:"bpm"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prev_hash"`
}

// CalculateHash is a helper to calculate hash for given block
func CalculateHash(b Block) string {
	record := string(b.Index) + b.Timestamp + string(b.BPM) + b.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// GenerateBlock generates new block from old block
func GenerateBlock(oldBlock Block, bpm int) (Block, error) {
	var newBlock Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = bpm
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock, nil
}

// IsBlockValid validates whether block is valid
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
