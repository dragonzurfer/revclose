package revclose

import "math"

func GetSignalValue(reversal float64, close float64, numberOfCrossedLevels int) SignalValue {
	if numberOfCrossedLevels == 0 || reversal == math.SmallestNonzeroFloat64 {
		return Nuetral
	}
	switch {
	case reversal > close:
		return Sell
	case reversal < close:
		return Buy
	}
	return Nuetral
}
