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

package fr

// expBySqrtExp is equivalent to z.Exp(x, /Users/gbotrel/dev/go/src/github.com/consensys/gnark-crypto/internal/generator/addchain/183227397098d014dc2822db40c0ac2e9419f4243cdcb848a1f0fac9f)
//
// uses github.com/mmcloughlin/addchain v0.4.0 to generate a shorter addition chain
func (z *Element) expBySqrtExp(x *Element) *Element {
	// addition chain:
	//
	//	_10    = 2*1
	//	_11    = 1 + _10
	//	_101   = _10 + _11
	//	_111   = _10 + _101
	//	_1001  = _10 + _111
	//	_1011  = _10 + _1001
	//	_1101  = _10 + _1011
	//	_1111  = _10 + _1101
	//	_11000 = _1001 + _1111
	//	_11111 = _111 + _11000
	//	i26    = ((_11000 << 4 + _11) << 3 + 1) << 7
	//	i36    = ((_1001 + i26) << 2 + _11) << 5 + _111
	//	i53    = (2*(i36 << 6 + _1011) + 1) << 8
	//	i64    = (2*(_1001 + i53) + 1) << 7 + _1101
	//	i84    = ((i64 << 10 + _101) << 6 + _1101) << 2
	//	i100   = ((_11 + i84) << 7 + _101) << 6 + 1
	//	i117   = ((i100 << 7 + _1011) << 5 + _1101) << 3
	//	i137   = ((_101 + i117) << 8 + _11) << 9 + _101
	//	i153   = ((i137 << 3 + _11) << 8 + _1011) << 3
	//	i168   = ((_101 + i153) << 5 + _101) << 7 + _11
	//	i187   = ((i168 << 7 + _11111) << 2 + 1) << 8
	//	i204   = ((_1001 + i187) << 8 + _1111) << 6 + _1101
	//	i215   = 2*((i204 << 2 + _11) << 6 + _1011)
	//	i232   = ((1 + i215) << 8 + _1001) << 6 + _101
	//	i257   = ((i232 << 9 + _11111) << 9 + _11111) << 5
	//	return   ((_1011 + i257) << 3 + 1) << 7 + _11111
	//
	// Operations: 221 squares 49 multiplies

	// Allocate Temporaries.
	var t0, t1, t2, t3, t4, t5, t6, t7 Element
	// Step 1: z = x^0x2
	z.Square(x)

	// Step 2: t3 = x^0x3
	t3.Mul(x, z)

	// Step 3: t1 = x^0x5
	t1.Mul(z, &t3)

	// Step 4: t6 = x^0x7
	t6.Mul(z, &t1)

	// Step 5: t2 = x^0x9
	t2.Mul(z, &t6)

	// Step 6: t0 = x^0xb
	t0.Mul(z, &t2)

	// Step 7: t4 = x^0xd
	t4.Mul(z, &t0)

	// Step 8: t5 = x^0xf
	t5.Mul(z, &t4)

	// Step 9: t7 = x^0x18
	t7.Mul(&t2, &t5)

	// Step 10: z = x^0x1f
	z.Mul(&t6, &t7)

	// Step 14: t7 = x^0x180
	for s := 0; s < 4; s++ {
		t7.Square(&t7)
	}

	// Step 15: t7 = x^0x183
	t7.Mul(&t3, &t7)

	// Step 18: t7 = x^0xc18
	for s := 0; s < 3; s++ {
		t7.Square(&t7)
	}

	// Step 19: t7 = x^0xc19
	t7.Mul(x, &t7)

	// Step 26: t7 = x^0x60c80
	for s := 0; s < 7; s++ {
		t7.Square(&t7)
	}

	// Step 27: t7 = x^0x60c89
	t7.Mul(&t2, &t7)

	// Step 29: t7 = x^0x183224
	for s := 0; s < 2; s++ {
		t7.Square(&t7)
	}

	// Step 30: t7 = x^0x183227
	t7.Mul(&t3, &t7)

	// Step 35: t7 = x^0x30644e0
	for s := 0; s < 5; s++ {
		t7.Square(&t7)
	}

	// Step 36: t6 = x^0x30644e7
	t6.Mul(&t6, &t7)

	// Step 42: t6 = x^0xc19139c0
	for s := 0; s < 6; s++ {
		t6.Square(&t6)
	}

	// Step 43: t6 = x^0xc19139cb
	t6.Mul(&t0, &t6)

	// Step 44: t6 = x^0x183227396
	t6.Square(&t6)

	// Step 45: t6 = x^0x183227397
	t6.Mul(x, &t6)

	// Step 53: t6 = x^0x18322739700
	for s := 0; s < 8; s++ {
		t6.Square(&t6)
	}

	// Step 54: t6 = x^0x18322739709
	t6.Mul(&t2, &t6)

	// Step 55: t6 = x^0x30644e72e12
	t6.Square(&t6)

	// Step 56: t6 = x^0x30644e72e13
	t6.Mul(x, &t6)

	// Step 63: t6 = x^0x1832273970980
	for s := 0; s < 7; s++ {
		t6.Square(&t6)
	}

	// Step 64: t6 = x^0x183227397098d
	t6.Mul(&t4, &t6)

	// Step 74: t6 = x^0x60c89ce5c263400
	for s := 0; s < 10; s++ {
		t6.Square(&t6)
	}

	// Step 75: t6 = x^0x60c89ce5c263405
	t6.Mul(&t1, &t6)

	// Step 81: t6 = x^0x183227397098d0140
	for s := 0; s < 6; s++ {
		t6.Square(&t6)
	}

	// Step 82: t6 = x^0x183227397098d014d
	t6.Mul(&t4, &t6)

	// Step 84: t6 = x^0x60c89ce5c26340534
	for s := 0; s < 2; s++ {
		t6.Square(&t6)
	}

	// Step 85: t6 = x^0x60c89ce5c26340537
	t6.Mul(&t3, &t6)

	// Step 92: t6 = x^0x30644e72e131a029b80
	for s := 0; s < 7; s++ {
		t6.Square(&t6)
	}

	// Step 93: t6 = x^0x30644e72e131a029b85
	t6.Mul(&t1, &t6)

	// Step 99: t6 = x^0xc19139cb84c680a6e140
	for s := 0; s < 6; s++ {
		t6.Square(&t6)
	}

	// Step 100: t6 = x^0xc19139cb84c680a6e141
	t6.Mul(x, &t6)

	// Step 107: t6 = x^0x60c89ce5c263405370a080
	for s := 0; s < 7; s++ {
		t6.Square(&t6)
	}

	// Step 108: t6 = x^0x60c89ce5c263405370a08b
	t6.Mul(&t0, &t6)

	// Step 113: t6 = x^0xc19139cb84c680a6e141160
	for s := 0; s < 5; s++ {
		t6.Square(&t6)
	}

	// Step 114: t6 = x^0xc19139cb84c680a6e14116d
	t6.Mul(&t4, &t6)

	// Step 117: t6 = x^0x60c89ce5c263405370a08b68
	for s := 0; s < 3; s++ {
		t6.Square(&t6)
	}

	// Step 118: t6 = x^0x60c89ce5c263405370a08b6d
	t6.Mul(&t1, &t6)

	// Step 126: t6 = x^0x60c89ce5c263405370a08b6d00
	for s := 0; s < 8; s++ {
		t6.Square(&t6)
	}

	// Step 127: t6 = x^0x60c89ce5c263405370a08b6d03
	t6.Mul(&t3, &t6)

	// Step 136: t6 = x^0xc19139cb84c680a6e14116da0600
	for s := 0; s < 9; s++ {
		t6.Square(&t6)
	}

	// Step 137: t6 = x^0xc19139cb84c680a6e14116da0605
	t6.Mul(&t1, &t6)

	// Step 140: t6 = x^0x60c89ce5c263405370a08b6d03028
	for s := 0; s < 3; s++ {
		t6.Square(&t6)
	}

	// Step 141: t6 = x^0x60c89ce5c263405370a08b6d0302b
	t6.Mul(&t3, &t6)

	// Step 149: t6 = x^0x60c89ce5c263405370a08b6d0302b00
	for s := 0; s < 8; s++ {
		t6.Square(&t6)
	}

	// Step 150: t6 = x^0x60c89ce5c263405370a08b6d0302b0b
	t6.Mul(&t0, &t6)

	// Step 153: t6 = x^0x30644e72e131a029b85045b681815858
	for s := 0; s < 3; s++ {
		t6.Square(&t6)
	}

	// Step 154: t6 = x^0x30644e72e131a029b85045b68181585d
	t6.Mul(&t1, &t6)

	// Step 159: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba0
	for s := 0; s < 5; s++ {
		t6.Square(&t6)
	}

	// Step 160: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba5
	t6.Mul(&t1, &t6)

	// Step 167: t6 = x^0x30644e72e131a029b85045b68181585d280
	for s := 0; s < 7; s++ {
		t6.Square(&t6)
	}

	// Step 168: t6 = x^0x30644e72e131a029b85045b68181585d283
	t6.Mul(&t3, &t6)

	// Step 175: t6 = x^0x183227397098d014dc2822db40c0ac2e94180
	for s := 0; s < 7; s++ {
		t6.Square(&t6)
	}

	// Step 176: t6 = x^0x183227397098d014dc2822db40c0ac2e9419f
	t6.Mul(z, &t6)

	// Step 178: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba5067c
	for s := 0; s < 2; s++ {
		t6.Square(&t6)
	}

	// Step 179: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d
	t6.Mul(x, &t6)

	// Step 187: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d00
	for s := 0; s < 8; s++ {
		t6.Square(&t6)
	}

	// Step 188: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d09
	t6.Mul(&t2, &t6)

	// Step 196: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d0900
	for s := 0; s < 8; s++ {
		t6.Square(&t6)
	}

	// Step 197: t5 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d090f
	t5.Mul(&t5, &t6)

	// Step 203: t5 = x^0x183227397098d014dc2822db40c0ac2e9419f4243c0
	for s := 0; s < 6; s++ {
		t5.Square(&t5)
	}

	// Step 204: t4 = x^0x183227397098d014dc2822db40c0ac2e9419f4243cd
	t4.Mul(&t4, &t5)

	// Step 206: t4 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d090f34
	for s := 0; s < 2; s++ {
		t4.Square(&t4)
	}

	// Step 207: t3 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d090f37
	t3.Mul(&t3, &t4)

	// Step 213: t3 = x^0x183227397098d014dc2822db40c0ac2e9419f4243cdc0
	for s := 0; s < 6; s++ {
		t3.Square(&t3)
	}

	// Step 214: t3 = x^0x183227397098d014dc2822db40c0ac2e9419f4243cdcb
	t3.Mul(&t0, &t3)

	// Step 215: t3 = x^0x30644e72e131a029b85045b68181585d2833e84879b96
	t3.Square(&t3)

	// Step 216: t3 = x^0x30644e72e131a029b85045b68181585d2833e84879b97
	t3.Mul(x, &t3)

	// Step 224: t3 = x^0x30644e72e131a029b85045b68181585d2833e84879b9700
	for s := 0; s < 8; s++ {
		t3.Square(&t3)
	}

	// Step 225: t2 = x^0x30644e72e131a029b85045b68181585d2833e84879b9709
	t2.Mul(&t2, &t3)

	// Step 231: t2 = x^0xc19139cb84c680a6e14116da06056174a0cfa121e6e5c240
	for s := 0; s < 6; s++ {
		t2.Square(&t2)
	}

	// Step 232: t1 = x^0xc19139cb84c680a6e14116da06056174a0cfa121e6e5c245
	t1.Mul(&t1, &t2)

	// Step 241: t1 = x^0x183227397098d014dc2822db40c0ac2e9419f4243cdcb848a00
	for s := 0; s < 9; s++ {
		t1.Square(&t1)
	}

	// Step 242: t1 = x^0x183227397098d014dc2822db40c0ac2e9419f4243cdcb848a1f
	t1.Mul(z, &t1)

	// Step 251: t1 = x^0x30644e72e131a029b85045b68181585d2833e84879b9709143e00
	for s := 0; s < 9; s++ {
		t1.Square(&t1)
	}

	// Step 252: t1 = x^0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f
	t1.Mul(z, &t1)

	// Step 257: t1 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d090f372e12287c3e0
	for s := 0; s < 5; s++ {
		t1.Square(&t1)
	}

	// Step 258: t0 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d090f372e12287c3eb
	t0.Mul(&t0, &t1)

	// Step 261: t0 = x^0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f58
	for s := 0; s < 3; s++ {
		t0.Square(&t0)
	}

	// Step 262: t0 = x^0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f59
	t0.Mul(x, &t0)

	// Step 269: t0 = x^0x183227397098d014dc2822db40c0ac2e9419f4243cdcb848a1f0fac80
	for s := 0; s < 7; s++ {
		t0.Square(&t0)
	}

	// Step 270: z = x^0x183227397098d014dc2822db40c0ac2e9419f4243cdcb848a1f0fac9f
	z.Mul(z, &t0)

	return z
}

// expByLegendreExp is equivalent to z.Exp(x, /Users/gbotrel/dev/go/src/github.com/consensys/gnark-crypto/internal/generator/addchain/183227397098d014dc2822db40c0ac2e9419f4243cdcb848a1f0fac9f8000000)
//
// uses github.com/mmcloughlin/addchain v0.4.0 to generate a shorter addition chain
func (z *Element) expByLegendreExp(x *Element) *Element {
	// addition chain:
	//
	//	_10    = 2*1
	//	_11    = 1 + _10
	//	_101   = _10 + _11
	//	_111   = _10 + _101
	//	_1001  = _10 + _111
	//	_1011  = _10 + _1001
	//	_1101  = _10 + _1011
	//	_1111  = _10 + _1101
	//	_11000 = _1001 + _1111
	//	_11111 = _111 + _11000
	//	i26    = ((_11000 << 4 + _11) << 3 + 1) << 7
	//	i36    = ((_1001 + i26) << 2 + _11) << 5 + _111
	//	i53    = (2*(i36 << 6 + _1011) + 1) << 8
	//	i64    = (2*(_1001 + i53) + 1) << 7 + _1101
	//	i84    = ((i64 << 10 + _101) << 6 + _1101) << 2
	//	i100   = ((_11 + i84) << 7 + _101) << 6 + 1
	//	i117   = ((i100 << 7 + _1011) << 5 + _1101) << 3
	//	i137   = ((_101 + i117) << 8 + _11) << 9 + _101
	//	i153   = ((i137 << 3 + _11) << 8 + _1011) << 3
	//	i168   = ((_101 + i153) << 5 + _101) << 7 + _11
	//	i187   = ((i168 << 7 + _11111) << 2 + 1) << 8
	//	i204   = ((_1001 + i187) << 8 + _1111) << 6 + _1101
	//	i215   = 2*((i204 << 2 + _11) << 6 + _1011)
	//	i232   = ((1 + i215) << 8 + _1001) << 6 + _101
	//	i257   = ((i232 << 9 + _11111) << 9 + _11111) << 5
	//	i270   = ((_1011 + i257) << 3 + 1) << 7 + _11111
	//	return   (2*i270 + 1) << 27
	//
	// Operations: 249 squares 50 multiplies

	// Allocate Temporaries.
	var t0, t1, t2, t3, t4, t5, t6, t7 Element
	// Step 1: z = x^0x2
	z.Square(x)

	// Step 2: t3 = x^0x3
	t3.Mul(x, z)

	// Step 3: t1 = x^0x5
	t1.Mul(z, &t3)

	// Step 4: t6 = x^0x7
	t6.Mul(z, &t1)

	// Step 5: t2 = x^0x9
	t2.Mul(z, &t6)

	// Step 6: t0 = x^0xb
	t0.Mul(z, &t2)

	// Step 7: t4 = x^0xd
	t4.Mul(z, &t0)

	// Step 8: t5 = x^0xf
	t5.Mul(z, &t4)

	// Step 9: t7 = x^0x18
	t7.Mul(&t2, &t5)

	// Step 10: z = x^0x1f
	z.Mul(&t6, &t7)

	// Step 14: t7 = x^0x180
	for s := 0; s < 4; s++ {
		t7.Square(&t7)
	}

	// Step 15: t7 = x^0x183
	t7.Mul(&t3, &t7)

	// Step 18: t7 = x^0xc18
	for s := 0; s < 3; s++ {
		t7.Square(&t7)
	}

	// Step 19: t7 = x^0xc19
	t7.Mul(x, &t7)

	// Step 26: t7 = x^0x60c80
	for s := 0; s < 7; s++ {
		t7.Square(&t7)
	}

	// Step 27: t7 = x^0x60c89
	t7.Mul(&t2, &t7)

	// Step 29: t7 = x^0x183224
	for s := 0; s < 2; s++ {
		t7.Square(&t7)
	}

	// Step 30: t7 = x^0x183227
	t7.Mul(&t3, &t7)

	// Step 35: t7 = x^0x30644e0
	for s := 0; s < 5; s++ {
		t7.Square(&t7)
	}

	// Step 36: t6 = x^0x30644e7
	t6.Mul(&t6, &t7)

	// Step 42: t6 = x^0xc19139c0
	for s := 0; s < 6; s++ {
		t6.Square(&t6)
	}

	// Step 43: t6 = x^0xc19139cb
	t6.Mul(&t0, &t6)

	// Step 44: t6 = x^0x183227396
	t6.Square(&t6)

	// Step 45: t6 = x^0x183227397
	t6.Mul(x, &t6)

	// Step 53: t6 = x^0x18322739700
	for s := 0; s < 8; s++ {
		t6.Square(&t6)
	}

	// Step 54: t6 = x^0x18322739709
	t6.Mul(&t2, &t6)

	// Step 55: t6 = x^0x30644e72e12
	t6.Square(&t6)

	// Step 56: t6 = x^0x30644e72e13
	t6.Mul(x, &t6)

	// Step 63: t6 = x^0x1832273970980
	for s := 0; s < 7; s++ {
		t6.Square(&t6)
	}

	// Step 64: t6 = x^0x183227397098d
	t6.Mul(&t4, &t6)

	// Step 74: t6 = x^0x60c89ce5c263400
	for s := 0; s < 10; s++ {
		t6.Square(&t6)
	}

	// Step 75: t6 = x^0x60c89ce5c263405
	t6.Mul(&t1, &t6)

	// Step 81: t6 = x^0x183227397098d0140
	for s := 0; s < 6; s++ {
		t6.Square(&t6)
	}

	// Step 82: t6 = x^0x183227397098d014d
	t6.Mul(&t4, &t6)

	// Step 84: t6 = x^0x60c89ce5c26340534
	for s := 0; s < 2; s++ {
		t6.Square(&t6)
	}

	// Step 85: t6 = x^0x60c89ce5c26340537
	t6.Mul(&t3, &t6)

	// Step 92: t6 = x^0x30644e72e131a029b80
	for s := 0; s < 7; s++ {
		t6.Square(&t6)
	}

	// Step 93: t6 = x^0x30644e72e131a029b85
	t6.Mul(&t1, &t6)

	// Step 99: t6 = x^0xc19139cb84c680a6e140
	for s := 0; s < 6; s++ {
		t6.Square(&t6)
	}

	// Step 100: t6 = x^0xc19139cb84c680a6e141
	t6.Mul(x, &t6)

	// Step 107: t6 = x^0x60c89ce5c263405370a080
	for s := 0; s < 7; s++ {
		t6.Square(&t6)
	}

	// Step 108: t6 = x^0x60c89ce5c263405370a08b
	t6.Mul(&t0, &t6)

	// Step 113: t6 = x^0xc19139cb84c680a6e141160
	for s := 0; s < 5; s++ {
		t6.Square(&t6)
	}

	// Step 114: t6 = x^0xc19139cb84c680a6e14116d
	t6.Mul(&t4, &t6)

	// Step 117: t6 = x^0x60c89ce5c263405370a08b68
	for s := 0; s < 3; s++ {
		t6.Square(&t6)
	}

	// Step 118: t6 = x^0x60c89ce5c263405370a08b6d
	t6.Mul(&t1, &t6)

	// Step 126: t6 = x^0x60c89ce5c263405370a08b6d00
	for s := 0; s < 8; s++ {
		t6.Square(&t6)
	}

	// Step 127: t6 = x^0x60c89ce5c263405370a08b6d03
	t6.Mul(&t3, &t6)

	// Step 136: t6 = x^0xc19139cb84c680a6e14116da0600
	for s := 0; s < 9; s++ {
		t6.Square(&t6)
	}

	// Step 137: t6 = x^0xc19139cb84c680a6e14116da0605
	t6.Mul(&t1, &t6)

	// Step 140: t6 = x^0x60c89ce5c263405370a08b6d03028
	for s := 0; s < 3; s++ {
		t6.Square(&t6)
	}

	// Step 141: t6 = x^0x60c89ce5c263405370a08b6d0302b
	t6.Mul(&t3, &t6)

	// Step 149: t6 = x^0x60c89ce5c263405370a08b6d0302b00
	for s := 0; s < 8; s++ {
		t6.Square(&t6)
	}

	// Step 150: t6 = x^0x60c89ce5c263405370a08b6d0302b0b
	t6.Mul(&t0, &t6)

	// Step 153: t6 = x^0x30644e72e131a029b85045b681815858
	for s := 0; s < 3; s++ {
		t6.Square(&t6)
	}

	// Step 154: t6 = x^0x30644e72e131a029b85045b68181585d
	t6.Mul(&t1, &t6)

	// Step 159: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba0
	for s := 0; s < 5; s++ {
		t6.Square(&t6)
	}

	// Step 160: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba5
	t6.Mul(&t1, &t6)

	// Step 167: t6 = x^0x30644e72e131a029b85045b68181585d280
	for s := 0; s < 7; s++ {
		t6.Square(&t6)
	}

	// Step 168: t6 = x^0x30644e72e131a029b85045b68181585d283
	t6.Mul(&t3, &t6)

	// Step 175: t6 = x^0x183227397098d014dc2822db40c0ac2e94180
	for s := 0; s < 7; s++ {
		t6.Square(&t6)
	}

	// Step 176: t6 = x^0x183227397098d014dc2822db40c0ac2e9419f
	t6.Mul(z, &t6)

	// Step 178: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba5067c
	for s := 0; s < 2; s++ {
		t6.Square(&t6)
	}

	// Step 179: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d
	t6.Mul(x, &t6)

	// Step 187: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d00
	for s := 0; s < 8; s++ {
		t6.Square(&t6)
	}

	// Step 188: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d09
	t6.Mul(&t2, &t6)

	// Step 196: t6 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d0900
	for s := 0; s < 8; s++ {
		t6.Square(&t6)
	}

	// Step 197: t5 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d090f
	t5.Mul(&t5, &t6)

	// Step 203: t5 = x^0x183227397098d014dc2822db40c0ac2e9419f4243c0
	for s := 0; s < 6; s++ {
		t5.Square(&t5)
	}

	// Step 204: t4 = x^0x183227397098d014dc2822db40c0ac2e9419f4243cd
	t4.Mul(&t4, &t5)

	// Step 206: t4 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d090f34
	for s := 0; s < 2; s++ {
		t4.Square(&t4)
	}

	// Step 207: t3 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d090f37
	t3.Mul(&t3, &t4)

	// Step 213: t3 = x^0x183227397098d014dc2822db40c0ac2e9419f4243cdc0
	for s := 0; s < 6; s++ {
		t3.Square(&t3)
	}

	// Step 214: t3 = x^0x183227397098d014dc2822db40c0ac2e9419f4243cdcb
	t3.Mul(&t0, &t3)

	// Step 215: t3 = x^0x30644e72e131a029b85045b68181585d2833e84879b96
	t3.Square(&t3)

	// Step 216: t3 = x^0x30644e72e131a029b85045b68181585d2833e84879b97
	t3.Mul(x, &t3)

	// Step 224: t3 = x^0x30644e72e131a029b85045b68181585d2833e84879b9700
	for s := 0; s < 8; s++ {
		t3.Square(&t3)
	}

	// Step 225: t2 = x^0x30644e72e131a029b85045b68181585d2833e84879b9709
	t2.Mul(&t2, &t3)

	// Step 231: t2 = x^0xc19139cb84c680a6e14116da06056174a0cfa121e6e5c240
	for s := 0; s < 6; s++ {
		t2.Square(&t2)
	}

	// Step 232: t1 = x^0xc19139cb84c680a6e14116da06056174a0cfa121e6e5c245
	t1.Mul(&t1, &t2)

	// Step 241: t1 = x^0x183227397098d014dc2822db40c0ac2e9419f4243cdcb848a00
	for s := 0; s < 9; s++ {
		t1.Square(&t1)
	}

	// Step 242: t1 = x^0x183227397098d014dc2822db40c0ac2e9419f4243cdcb848a1f
	t1.Mul(z, &t1)

	// Step 251: t1 = x^0x30644e72e131a029b85045b68181585d2833e84879b9709143e00
	for s := 0; s < 9; s++ {
		t1.Square(&t1)
	}

	// Step 252: t1 = x^0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f
	t1.Mul(z, &t1)

	// Step 257: t1 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d090f372e12287c3e0
	for s := 0; s < 5; s++ {
		t1.Square(&t1)
	}

	// Step 258: t0 = x^0x60c89ce5c263405370a08b6d0302b0ba5067d090f372e12287c3eb
	t0.Mul(&t0, &t1)

	// Step 261: t0 = x^0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f58
	for s := 0; s < 3; s++ {
		t0.Square(&t0)
	}

	// Step 262: t0 = x^0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f59
	t0.Mul(x, &t0)

	// Step 269: t0 = x^0x183227397098d014dc2822db40c0ac2e9419f4243cdcb848a1f0fac80
	for s := 0; s < 7; s++ {
		t0.Square(&t0)
	}

	// Step 270: z = x^0x183227397098d014dc2822db40c0ac2e9419f4243cdcb848a1f0fac9f
	z.Mul(z, &t0)

	// Step 271: z = x^0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f593e
	z.Square(z)

	// Step 272: z = x^0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f
	z.Mul(x, z)

	// Step 299: z = x^0x183227397098d014dc2822db40c0ac2e9419f4243cdcb848a1f0fac9f8000000
	for s := 0; s < 27; s++ {
		z.Square(z)
	}

	return z
}
