// Copyright 2020 ConsenSys Software Inc.
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

// Code generated by consensys/gnark-crypto DO NOT EDIT

package mimc

import (
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
)

// Decompose interpret rawBytes as a bigInt x in big endian,
// and returns the digits of x (from LSB to MSB) when x is written
// in basis modulo.
func Decompose(rawBytes []byte) []fr.Element {

	rawBigInt := big.NewInt(0).SetBytes(rawBytes)
	modulo := fr.Modulus()

	// maximum number of chunks that a function
	maxNbChunks := len(rawBytes) / fr.Bytes

	res := make([]fr.Element, 0, maxNbChunks)
	var tmp fr.Element
	t := new(big.Int)
	for rawBigInt.Sign() != 0 {
		rawBigInt.DivMod(rawBigInt, modulo, t)
		tmp.SetBigInt(t)
		res = append(res, tmp)
	}

	return res
}
