// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package eth

import (
	"math/big"
	"sort"
	"sync"

	"github.com/gorievm/go-gori/common"
	"github.com/gorievm/go-gori/consensus"
	"github.com/gorievm/go-gori/consensus/ethash"
	"github.com/gorievm/go-gori/core"
	"github.com/gorievm/go-gori/core/rawdb"
	"github.com/gorievm/go-gori/core/txpool"
	"github.com/gorievm/go-gori/core/types"
	"github.com/gorievm/go-gori/core/vm"
	"github.com/gorievm/go-gori/crypto"
	"github.com/gorievm/go-gori/eth/downloader"
	"github.com/gorievm/go-gori/ethdb"
	"github.com/gorievm/go-gori/event"
	"github.com/gorievm/go-gori/params"
)

var (
	// testKey is a private key to use for funding a tester account.
	testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

	// testAddr is the Ori address of the tester account.
	testAddr = crypto.PubkeyToAddress(testKey.PublicKey)
)

// testTxPool is a mock transaction pool that blindly accepts all transactions.
// Its goal is to get around setting up a valid statedb for the balance and nonce
// checks.
type testTxPool struct {
	pool map[common.Hash]*types.Transaction // Hash map of collected transactions

	txFeed event.Feed   // Notification feed to allow waiting for inclusion
	lock   sync.RWMutex // Protects the transaction pool
}

// newTestTxPool creates a mock transaction pool.
func newTestTxPool() *testTxPool {
	return &testTxPool{
		pool: make(map[common.Hash]*types.Transaction),
	}
}

// Has returns an indicator whether txpool has a transaction
// cached with the given hash.
func (p *testTxPool) Has(hash common.Hash) bool {
	p.lock.Lock()
	defer p.lock.Unlock()

	return p.pool[hash] != nil
}

// Get retrieves the transaction from local txpool with given
// tx hash.
func (p *testTxPool) Get(hash common.Hash) *txpool.Transaction {
	p.lock.Lock()
	defer p.lock.Unlock()

	if tx := p.pool[hash]; tx != nil {
		return &txpool.Transaction{Tx: tx}
	}
	return nil
}

// Add appends a batch of transactions to the pool, and notifies any
// listeners if the addition channel is non nil
func (p *testTxPool) Add(txs []*txpool.Transaction, local bool, sync bool) []error {
	unwrapped := make([]*types.Transaction, len(txs))
	for i, tx := range txs {
		unwrapped[i] = tx.Tx
	}
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, tx := range unwrapped {
		p.pool[tx.Hash()] = tx
	}

	p.txFeed.Send(core.NewTxsEvent{Txs: unwrapped})
	return make([]error, len(unwrapped))
}

// Pending returns all the transactions known to the pool
func (p *testTxPool) Pending(enforceTips bool) map[common.Address][]*txpool.LazyTransaction {
	p.lock.RLock()
	defer p.lock.RUnlock()

	batches := make(map[common.Address][]*types.Transaction)
	for _, tx := range p.pool {
		from, _ := types.Sender(types.HomesteadSigner{}, tx)
		batches[from] = append(batches[from], tx)
	}
	for _, batch := range batches {
		sort.Sort(types.TxByNonce(batch))
	}
	pending := make(map[common.Address][]*txpool.LazyTransaction)
	for addr, batch := range batches {
		for _, tx := range batch {
			pending[addr] = append(pending[addr], &txpool.LazyTransaction{
				Hash:      tx.Hash(),
				Tx:        &txpool.Transaction{Tx: tx},
				Time:      tx.Time(),
				GasFeeCap: tx.GasFeeCap(),
				GasTipCap: tx.GasTipCap(),
			})
		}
	}
	return pending
}

// SubscribeNewTxsEvent should return an event subscription of NewTxsEvent and
// send events to the given channel.
func (p *testTxPool) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return p.txFeed.Subscribe(ch)
}

// testHandler is a live implementation of the Ori protocol handler, just
// preinitialized with some sane testing defaults and the transaction pool mocked
// out.
type testHandler struct {
	db      ethdb.Database
	chain   *core.BlockChain
	txpool  *testTxPool
	handler *handler
}

// newTestHandler creates a new handler for testing purposes with no blocks.
func newTestHandler() *testHandler {
	return newTestHandlerWithBlocks(0)
}

// newTestHandlerWithBlocks creates a new handler for testing purposes, with a
// given number of initial blocks.
func newTestHandlerWithBlocks(blocks int) *testHandler {
	// Create a database pre-initialize with a genesis block
	db := rawdb.NewMemoryDatabase()
	gspec := &core.Genesis{
		Config: params.TestChainConfig,
		Alloc:  core.GenesisAlloc{testAddr: {Balance: big.NewInt(1000000)}},
	}
	chain, _ := core.NewBlockChain(db, nil, gspec, nil, ethash.NewFaker(), vm.Config{}, nil, nil)

	_, bs, _ := core.GenerateChainWithGenesis(gspec, ethash.NewFaker(), blocks, nil)
	if _, err := chain.InsertChain(bs); err != nil {
		panic(err)
	}
	txpool := newTestTxPool()

	handler, _ := newHandler(&handlerConfig{
		Database:   db,
		Chain:      chain,
		TxPool:     txpool,
		Merger:     consensus.NewMerger(rawdb.NewMemoryDatabase()),
		Network:    1,
		Sync:       downloader.SnapSync,
		BloomCache: 1,
	})
	handler.Start(1000)

	return &testHandler{
		db:      db,
		chain:   chain,
		txpool:  txpool,
		handler: handler,
	}
}

// close tears down the handler and all its internal constructs.
func (b *testHandler) close() {
	b.handler.Stop()
	b.chain.Stop()
}
