// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"time"

	"github.com/arielizuardi/golang-blockchain/block"
	"github.com/arielizuardi/golang-blockchain/server"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

var (
	httpServer *server.Server
	blockChain []block.Block
)

// NewRootCmd creates an instance of root cobra command
func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "",
		Short: "Golang Block Chain",
		Run: func(_ *cobra.Command, _ []string) {
			initDependency()
			httpServer.StartHTTPServer()
		},
	}
}

func initDependency() {
	now := time.Now()
	genesisBlock := block.Block{
		Timestamp: now.String(),
	}

	spew.Dump(genesisBlock)

	blockchain := []block.Block{genesisBlock}
	httpServer = &server.Server{
		Port:       8088,
		Blockchain: blockchain,
	}
}
