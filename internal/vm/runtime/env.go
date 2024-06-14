// Copyright 2023 The astranet Authors
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

package runtime

import (
	"github.com/astranetworld/ast/internal"
	"github.com/astranetworld/ast/internal/vm"
	"github.com/astranetworld/ast/internal/vm/evmtypes"
)

func NewEnv(cfg *Config) *vm.EVM {
	txContext := evmtypes.TxContext{
		Origin:   cfg.Origin,
		GasPrice: cfg.GasPrice,
	}

	blockContext := evmtypes.BlockContext{
		CanTransfer: internal.CanTransfer,
		Transfer:    internal.Transfer,
		GetHash:     cfg.GetHashFn,
		Coinbase:    cfg.Coinbase,
		BlockNumber: cfg.BlockNumber.Uint64(),
		Time:        cfg.Time.Uint64(),
		Difficulty:  cfg.Difficulty,
		GasLimit:    cfg.GasLimit,
		BaseFee:     cfg.BaseFee,
	}

	return vm.NewEVM(blockContext, txContext, cfg.State, cfg.ChainConfig, cfg.EVMConfig)
}
