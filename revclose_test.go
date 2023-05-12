package revclose_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dragonzurfer/revclose"
)

type MockRevCandle struct {
	O, H, L, C float64
}

func (c *MockRevCandle) GetHigh() float64 {
	return c.H
}
func (c *MockRevCandle) GetLow() float64 {
	return c.L
}
func (c *MockRevCandle) GetClose() float64 {
	return c.C
}
func (c *MockRevCandle) GetOpen() float64 {
	return c.O
}

func (c *MockRevCandle) GetOHLC() (float64, float64, float64, float64) {
	return c.O, c.H, c.L, c.C
}

type MockRevCloseCandles struct {
	Candles []revclose.RevCandle
}

func (m *MockRevCloseCandles) GetCandle(index int) revclose.RevCandle {
	return m.Candles[index]
}

func (m *MockRevCloseCandles) GetCandlesLength() int {
	return len(m.Candles)
}

type MockLevel struct {
	ID      int
	Crossed bool
}

func (l *MockLevel) GetID() int {
	return l.ID
}

func (l *MockLevel) HasCrossed(value2 float64, value float64) bool {
	return l.Crossed
}

func signalsEqual(expected, actual revclose.Signal) bool {
	if expected.Reversal != actual.Reversal || expected.Close != actual.Close || expected.Value != actual.Value {
		return false
	}

	if len(expected.LevelsWithinRange) != len(actual.LevelsWithinRange) {
		return false
	}

	for i, expectedLevel := range expected.LevelsWithinRange {
		actualLevel := actual.LevelsWithinRange[i]

		expectedMockLevel, expectedOk := expectedLevel.(*MockLevel)
		actualMockLevel, actualOk := actualLevel.(*MockLevel)

		if !expectedOk || !actualOk || expectedMockLevel.ID != actualMockLevel.ID || expectedMockLevel.Crossed != actualMockLevel.Crossed {
			return false
		}
	}

	return true
}

func TestGetSignal(t *testing.T) {
	testCases := []string{
		"testcase1.json",
		"testcase2.json",
		"testcase3.json",
		"testcase4.json",
		"testcase5.json",
		"testcase6.json",
		"testcase7.json",
		// Add more test case file names here
	}

	for _, tc := range testCases {
		wd, err := os.Getwd()
		if err != nil {
			t.Fatalf("Error getting working directory: %v", err)
		}

		data, err := ioutil.ReadFile(filepath.Join(wd, "testcases", tc))
		if err != nil {
			t.Fatalf("Error reading test case file: %v", err)
		}

		var testCase struct {
			Input struct {
				Candles []MockRevCandle
				Levels  []MockLevel
				Close   float64
			}
			Output struct {
				Signal               revclose.Signal
				LevelsWithinRangeIDs []int
			}
		}

		err = json.Unmarshal(data, &testCase)
		if err != nil {
			t.Fatalf("Error unmarshaling test case data: %v", err)
		}

		candles := &MockRevCloseCandles{
			Candles: make([]revclose.RevCandle, len(testCase.Input.Candles)),
		}

		for i, c := range testCase.Input.Candles {
			candle := c // Create a new variable to store the value of c
			candles.Candles[i] = &candle
		}

		levels := make([]revclose.LevelInterface, len(testCase.Input.Levels))

		for i, l := range testCase.Input.Levels {
			level := l // Create a new variable to store the value of l
			levels[i] = &level
		}

		expectedSignal := testCase.Output.Signal
		expectedSignal.LevelsWithinRange = make([]revclose.LevelInterface, len(testCase.Output.LevelsWithinRangeIDs))
		for i, levelID := range testCase.Output.LevelsWithinRangeIDs {
			for _, level := range levels {
				mockLevel, _ := level.(*MockLevel)
				if mockLevel.ID == levelID {

					expectedSignal.LevelsWithinRange[i] = level
				}
			}
		}

		signal := revclose.GetSignal(candles, levels, testCase.Input.Close)
		// fmt.Printf("testcase signal: %+v\n Expected signal: %+v\n signal: %+v\n", testCase.Output.LevelsWithinRangeIDs, expectedSignal, signal)
		if !signalsEqual(expectedSignal, signal) {
			t.Errorf("The signal does not match the expected signal on %v\nExpected: %+v\nActual:   %+v", tc, expectedSignal, signal)
		}
	}
}
