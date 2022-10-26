package main

import (
	"encoding/json"
	"fmt"
	"github.com/consensys/gnark-crypto/internal/generator/test_vector_utils"
	"github.com/consensys/gnark-crypto/internal/generator/test_vector_utils/small_rational"
	"github.com/consensys/gnark-crypto/internal/generator/test_vector_utils/small_rational/gkr"
	"github.com/consensys/gnark-crypto/internal/generator/test_vector_utils/small_rational/polynomial"
	"github.com/consensys/gnark-crypto/internal/generator/test_vector_utils/small_rational/sumcheck"
	"os"
	"path/filepath"
)

func main() {
	if err := Generate(); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

// TODO: Lots of copy-pasting here. Fix that.

func Generate() error {
	testDirPath, err := filepath.Abs("gkr/test_vectors")
	if err != nil {
		return err
	}

	fmt.Printf("generating GKR test cases: scanning directory %s for test specs\n", testDirPath)
	//debug.PrintStack()

	dirEntries, err := os.ReadDir(testDirPath)
	if err != nil {
		return err
	}
	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {

			if filepath.Ext(dirEntry.Name()) == ".json" {

				fmt.Println("\tprocessing", dirEntry.Name())

				path := filepath.Join(testDirPath, dirEntry.Name())

				var testCase *TestCase
				testCase, err = newTestCase(path)
				if err != nil {
					return err
				}

				testCase.Transcript.Update(0)
				proof := gkr.Prove(testCase.Circuit, testCase.FullAssignment, testCase.Transcript)

				testCase.Info.Proof = toPrintableProof(proof)
				var outBytes []byte
				if outBytes, err = json.MarshalIndent(testCase.Info, "", "\t"); err == nil {
					if err = os.WriteFile(path, outBytes, 0); err != nil {
						return err
					}
				} else {
					return err
				}
			}
		}
	}

	return nil
}

type TestCaseInfo struct {
	Hash    string          `json:"hash"`
	Circuit string          `json:"circuit"`
	Input   [][]interface{} `json:"input"`
	Output  [][]interface{} `json:"output"`
	Proof   PrintableProof  `json:"proof"`
}

type PrintableProof [][]PrintableSumcheckProof

type PrintableSumcheckProof struct {
	FinalEvalProof  interface{}     `json:"finalEvalProof"`
	PartialSumPolys [][]interface{} `json:"partialSumPolys"`
}

type TestCase struct {
	Circuit         gkr.Circuit
	Transcript      sumcheck.ArithmeticTranscript
	FullAssignment  gkr.WireAssignment
	InOutAssignment gkr.WireAssignment
	Info            *TestCaseInfo
}

func newTestCase(path string) (*TestCase, error) {
	path, err := filepath.Abs(path)

	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(path)

	var bytes []byte
	if bytes, err = os.ReadFile(path); err == nil {
		var info TestCaseInfo
		if err = json.Unmarshal(bytes, &info); err != nil {
			return nil, err
		}

		var circuit gkr.Circuit
		if circuit, err = getCircuit(filepath.Join(dir, info.Circuit)); err != nil {
			return nil, err
		}

		var hash test_vector_utils.HashMap
		if hash, err = test_vector_utils.GetHash(filepath.Join(dir, info.Hash)); err != nil {
			return nil, err
		}

		fullAssignment := make(gkr.WireAssignment)
		assignmentSize := len(info.Input[0])

		{
			i := len(circuit) - 1

			if len(circuit[i]) != len(info.Input) {
				return nil, fmt.Errorf("input layer not the same size as input vector")
			}

			for j := range circuit[i] {
				wire := &circuit[i][j]
				var wireAssignment []small_rational.SmallRational
				if wireAssignment, err = test_vector_utils.SliceToElementSlice(info.Input[j]); err == nil {
					fullAssignment[wire] = wireAssignment
				} else {
					return nil, err
				}

			}
		}

		for i := len(circuit) - 2; i >= 0; i-- {
			for j := range circuit[i] {
				wire := &circuit[i][j]
				assignment := make(polynomial.MultiLin, assignmentSize)
				in := make([]small_rational.SmallRational, len(wire.Inputs))
				for k := range assignment {
					for l, inputWire := range circuit[i][j].Inputs {
						in[l] = fullAssignment[inputWire][k]
					}
					assignment[k] = wire.Gate.Evaluate(in...)
				}

				fullAssignment[wire] = assignment
			}
		}

		if len(circuit[0]) != len(info.Output) {
			return nil, fmt.Errorf("output layer not the same size as output vector: %d ≠ %d", len(circuit[0]), len(info.Output))
		}
		for j := range circuit[0] {
			var outAssignment []small_rational.SmallRational
			if outAssignment, err = test_vector_utils.SliceToElementSlice(info.Output[j]); err == nil {
				if err = test_vector_utils.SliceEquals(outAssignment, fullAssignment[&circuit[0][j]]); err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		}

		return &TestCase{
			Circuit:        circuit,
			Transcript:     &test_vector_utils.MapHashTranscript{HashMap: hash},
			FullAssignment: fullAssignment,
			Info:           &info,
		}, nil

	} else {
		return nil, err
	}
}

type WireInfo struct {
	Gate   string  `json:"gate"`
	Inputs [][]int `json:"inputs"`
}

type CircuitInfo [][]WireInfo

var circuitCache = make(map[string]gkr.Circuit)

func getCircuit(path string) (gkr.Circuit, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	if circuit, ok := circuitCache[path]; ok {
		return circuit, nil
	}
	var bytes []byte
	if bytes, err = os.ReadFile(path); err == nil {
		var circuitInfo CircuitInfo
		if err = json.Unmarshal(bytes, &circuitInfo); err == nil {
			circuit := circuitInfo.toCircuit()
			circuitCache[path] = circuit
			return circuit, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (c CircuitInfo) toCircuit() (circuit gkr.Circuit) {
	isOutput := make(map[*gkr.Wire]interface{})
	circuit = make(gkr.Circuit, len(c))
	for i := len(c) - 1; i >= 0; i-- {
		circuit[i] = make(gkr.CircuitLayer, len(c[i]))
		for j, wireInfo := range c[i] {
			circuit[i][j].Gate = gates[wireInfo.Gate]
			circuit[i][j].Inputs = make([]*gkr.Wire, len(wireInfo.Inputs))
			isOutput[&circuit[i][j]] = nil
			for k, inputCoord := range wireInfo.Inputs {
				if len(inputCoord) != 2 {
					panic("circuit wire has two coordinates")
				}
				input := &circuit[inputCoord[0]][inputCoord[1]]
				input.NumOutputs++
				circuit[i][j].Inputs[k] = input
				delete(isOutput, input)
			}
			if (i == len(c)-1) != (len(circuit[i][j].Inputs) == 0) {
				panic("wire is input if and only if in last layer")
			}
		}
	}

	for k := range isOutput {
		k.NumOutputs = 1
	}

	return
}

var gates map[string]gkr.Gate

func init() {
	gates = make(map[string]gkr.Gate)
	gates["identity"] = identityGate{}
	gates["mul"] = mulGate{}
	gates["mimc"] = mimcCipherGate{} //TODO: Add ark
}

type identityGate struct{}

func (g identityGate) Evaluate(i ...small_rational.SmallRational) small_rational.SmallRational {
	if len(i) != 1 {
		panic("identity operates on one field element")
	}
	return i[0]
}

func (g identityGate) Degree() int {
	return 1
}

type mulGate struct{}

func (m mulGate) Evaluate(element ...small_rational.SmallRational) (result small_rational.SmallRational) {
	result.Mul(&element[0], &element[1])
	return
}

func (m mulGate) Degree() int {
	return 2
}

type mimcCipherGate struct {
	ark small_rational.SmallRational
}

func (m mimcCipherGate) Evaluate(input ...small_rational.SmallRational) (res small_rational.SmallRational) {
	var sum small_rational.SmallRational

	sum.
		Add(&input[0], &input[1]).
		Add(&sum, &m.ark)

	res.Square(&sum)    // sum^2
	res.Mul(&res, &sum) // sum^3
	res.Square(&res)    //sum^6
	res.Mul(&res, &sum) //sum^7

	return
}

func (m mimcCipherGate) Degree() int {
	return 7
}

func toPrintableProof(proof gkr.Proof) PrintableProof {
	res := make(PrintableProof, len(proof))

	for i, proofI := range proof {
		if proofI == nil {
			res[i] = nil
		} else {
			res[i] = make([]PrintableSumcheckProof, len(proofI))
			for j, proofIJ := range proofI {

				partialSumPolys := make([][]interface{}, len(proofIJ.PartialSumPolys))
				for k, partialK := range proofIJ.PartialSumPolys {
					partialSumPolys[k] = test_vector_utils.ElementSliceToInterfaceSlice(partialK)
				}

				res[i][j] = PrintableSumcheckProof{
					FinalEvalProof:  test_vector_utils.ElementSliceToInterfaceSlice(proofIJ.FinalEvalProof),
					PartialSumPolys: partialSumPolys,
				}
			}
		}
	}
	return res
}
