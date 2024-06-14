// Copyright 2022 The astranet Authors
// This file is part of the astranet library.
//
// The astranet library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The astranet library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the astranet library. If not, see <http://www.gnu.org/licenses/>.

package common

import (
	"context"
	"github.com/astranetworld/ast/common/block"
	"github.com/astranetworld/ast/common/message"
	"github.com/astranetworld/ast/common/transaction"
	"github.com/astranetworld/ast/common/types"
	"github.com/holiman/uint256"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"google.golang.org/protobuf/proto"
)

// Service is a struct that can be registered into a ServiceRegistry for
// easy dependency management.
type Service interface {
	// Start spawns any goroutines required by the service.
	//Start()
	// Stop terminates all goroutines belonging to the service,
	// blocking until they are all terminated.
	Stop() error
	// Status returns error if the service is not considered healthy.
	//Status() error
}

type IDownloader interface {
	SyncHeader() error
	SyncBody() error
	SyncTx() error
	Start() error
	Close() error
	IsDownloading() bool
	ConnHandler([]byte, peer.ID) error
}

type ConnHandler func([]byte, peer.ID) error

type ProtocolHandshakeFn func(peer IPeer, genesisHash types.Hash, currentHeight *uint256.Int) (Peer, bool)
type ProtocolHandshakeInfo func() (types.Hash, *uint256.Int, error)

type INetwork interface {
	WriterMessage(messageType message.MessageType, payload []byte, peer peer.ID) error
	//BroadcastMessage(messageType message.MessageType, payload []byte) (int, error)
	SetHandler(message.MessageType, ConnHandler) error
	ClosePeer(id peer.ID) error
	Start() error
	Host() host.Host
	PeerCount() int
	Bootstrapped() bool
}

type IPeer interface {
	ID() peer.ID
	Write(msg message.IMessage) error
	WriteMsg(messageType message.MessageType, payload []byte) error
	SetHandler(message.MessageType, ConnHandler) error
	ClearHandler(message.MessageType) error
	Close() error
}

type IPubSub interface {
	JoinTopic(topic string) (*pubsub.Topic, error)
	Publish(topic string, msg proto.Message) error
	GetTopics() []string
	Start() error
}

type IStateDB interface {
	CreateAccount(types.Address)

	SubBalance(addr types.Address, amount uint256.Int)
	AddBalance(addr types.Address, amount uint256.Int)
	GetBalance(addr types.Address) uint256.Int

	GetNonce(addr types.Address) uint64
	SetNonce(addr types.Address, nonce uint64)

	GetCodeHash(addr types.Address) types.Hash
	GetCode(addr types.Address) []byte
	SetCode(addr types.Address, code []byte)
	GetCodeSize(addr types.Address) int

	AddRefund(uint64)
	SubRefund(uint64)
	GetRefund() uint64

	GetCommittedState(types.Address, types.Hash) types.Hash
	GetState(types.Address, types.Hash) types.Hash
	SetState(types.Address, types.Hash, types.Hash)

	Suicide(types.Address) bool
	HasSuicided(types.Address) bool

	Exist(types.Address) bool
	Empty(types.Address) bool

	PrepareAccessList(sender types.Address, dest *types.Address, precompiles []types.Address, list transaction.AccessList)
	AddressInAccessList(addr types.Address) bool
	SlotInAccessList(addr types.Address, slot types.Hash) (addressOk bool, slotOk bool)
	AddAddressToAccessList(addr types.Address)
	AddSlotToAccessList(addr types.Address, slot types.Hash)

	RevertToSnapshot(int)
	Snapshot() int

	AddLog(*block.Log)
	GetLogs(hash types.Hash, blockHash types.Hash) []*block.Log

	TxIndex() int
	Prepare(thash types.Hash, ti int)

	Error() error
}
type ChainStateReader interface {
	BalanceAt(ctx context.Context, account types.Address, blockNumber uint256.Int) (uint256.Int, error)
	StorageAt(ctx context.Context, account types.Address, key types.Hash, blockNumber uint256.Int) ([]byte, error)
	CodeAt(ctx context.Context, account types.Address, blockNumber uint256.Int) ([]byte, error)
	NonceAt(ctx context.Context, account types.Address, blockNumber uint256.Int) (uint64, error)
}

type ITxsPool interface {
	Service
	Has(hash types.Hash) bool
	Pending(enforceTips bool) map[types.Address][]*transaction.Transaction
	GetTransaction() ([]*transaction.Transaction, error)
	GetTx(hash types.Hash) *transaction.Transaction
	AddRemotes(txs []*transaction.Transaction) []error
	AddLocal(tx *transaction.Transaction) error
	Stats() (int, int, int, int)
	Nonce(addr types.Address) uint64
	Content() (map[types.Address][]*transaction.Transaction, map[types.Address][]*transaction.Transaction)
}
