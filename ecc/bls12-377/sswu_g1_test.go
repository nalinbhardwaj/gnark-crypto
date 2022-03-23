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

package bls12377

import (
	"github.com/consensys/gnark-crypto/ecc/bls12-377/fp"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/prop"
	"math/rand"
	"testing"
)

func TestG1SqrtRatio(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	if testing.Short() {
		parameters.MinSuccessfulTests = nbFuzzShort
	} else {
		parameters.MinSuccessfulTests = nbFuzz
	}

	properties := gopter.NewProperties(parameters)

	gen := genCoordElemG1(t)

	properties.Property("G1SqrtRatio must square back to the right value", prop.ForAll(
		func(u fp.Element, v fp.Element) bool {

			var seen fp.Element
			qr := g1SqrtRatio(&seen, &u, &v) == 0

			seen.
				Square(&seen).
				Mul(&seen, &v)

			var ref fp.Element
			if qr {
				ref = u
			} else {
				g1MulByZ(&ref, &u)
			}

			return seen.Equal(&ref)
		}, gen, gen))

	properties.TestingRun(t, gopter.ConsoleReporter(false))
}

func genCoordElemG1(t *testing.T) gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {

		genRandomElem := func() fp.Element {
			var a fp.Element

			if _, err := a.SetRandom(); err != nil {
				t.Error(err)
			}

			return a
		}
		a := genRandomElem()

		genResult := gopter.NewGenResult(a, gopter.NoShrinker)
		return genResult
	}
}

func g1TestMatchCoord(t *testing.T, coordName string, msg string, expectedStr string, seen *fp.Element) {
	var expected fp.Element

	expected.SetString(expectedStr)

	if !expected.Equal(seen) {
		t.Errorf("mismatch on \"%s\", %s:\n\texpected %s\n\tsaw      %s", msg, coordName, expected.String(), seen)
	}
}

func g1TestMatch(t *testing.T, c hashTestCase, seen *G1Affine) {
	g1TestMatchCoord(t, "x", c.msg, c.x, &seen.X)
	g1TestMatchCoord(t, "y", c.msg, c.y, &seen.Y)
}

func TestEncodeToCurveG1SSWU(t *testing.T) {
	t.Parallel()
	for _, c := range g1EncodeToCurveSSWUVector.cases {
		seen, err := EncodeToCurveG1SSWU([]byte(c.msg), g1EncodeToCurveSSWUVector.dst)
		if err != nil {
			t.Fatal(err)
		}
		g1TestMatch(t, c, &seen)
	}
}

func TestHashToCurveG1SSWU(t *testing.T) {
	t.Parallel()
	for _, c := range g1HashToCurveSSWUVector.cases {
		seen, err := HashToCurveG1SSWU([]byte(c.msg), g1HashToCurveSSWUVector.dst)
		if err != nil {
			t.Fatal(err)
		}
		g1TestMatch(t, c, &seen)
	}
	t.Log(len(g1HashToCurveSSWUVector.cases), "cases verified")
}

func BenchmarkG1EncodeToCurveSSWU(b *testing.B) {
	const size = 54
	bytes := make([]byte, size)
	dst := g1EncodeToCurveSSWUVector.dst
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		bytes[rand.Int()%size] = byte(rand.Int())

		if _, err := EncodeToCurveG1SSWU(bytes, dst); err != nil {
			b.Fail()
		}
	}
}

func BenchmarkG1HashToCurveSSWU(b *testing.B) {
	const size = 54
	bytes := make([]byte, size)
	dst := g1HashToCurveSSWUVector.dst
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		bytes[rand.Int()%size] = byte(rand.Int())

		if _, err := HashToCurveG1SSWU(bytes, dst); err != nil {
			b.Fail()
		}
	}
}

type hashTestVector struct {
	dst   []byte
	cases []hashTestCase
}

type hashTestCase struct {
	msg string
	x   string
	y   string
}

var g1HashToCurveSSWUVector hashTestVector
var g1EncodeToCurveSSWUVector hashTestVector
