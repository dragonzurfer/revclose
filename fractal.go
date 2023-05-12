package revclose

import "math"

func isBullishFractal(candles RevCloseCandles, index int) bool {
	if index < 2 || index > candles.GetCandlesLength()-3 {
		return false
	}
	_, high, _, _ := candles.GetCandle(index).GetOHLC()
	_, prevHigh1, _, _ := candles.GetCandle(index - 1).GetOHLC()
	_, prevHigh2, _, _ := candles.GetCandle(index - 2).GetOHLC()
	_, nextHigh1, _, _ := candles.GetCandle(index + 1).GetOHLC()
	_, nextHigh2, _, _ := candles.GetCandle(index + 2).GetOHLC()

	return high > prevHigh1 &&
		high > prevHigh2 &&
		high > nextHigh1 &&
		high > nextHigh2
}

func isBearishFractal(candles RevCloseCandles, index int) bool {
	if index < 2 || index > candles.GetCandlesLength()-3 {
		return false
	}
	_, _, low, _ := candles.GetCandle(index).GetOHLC()
	_, _, prevLow1, _ := candles.GetCandle(index - 1).GetOHLC()
	_, _, prevLow2, _ := candles.GetCandle(index - 2).GetOHLC()
	_, _, nextLow1, _ := candles.GetCandle(index + 1).GetOHLC()
	_, _, nextLow2, _ := candles.GetCandle(index + 2).GetOHLC()

	return low < prevLow1 &&
		low < prevLow2 &&
		low < nextLow1 &&
		low < nextLow2
}

func getLatestReversal(candles RevCloseCandles) (float64, ReversalType) {
	latestIndex := candles.GetCandlesLength() - 1
	for i := latestIndex; i >= 2; i-- {
		bullishFractalExist, bearishFractalExist := false, false
		price := 0.0
		var reversalType ReversalType
		if isBullishFractal(candles, i) {
			_, high, _, _ := candles.GetCandle(i).GetOHLC()
			bullishFractalExist = true
			price, reversalType = high, Bullish
		}
		if isBearishFractal(candles, i) {
			_, _, low, _ := candles.GetCandle(i).GetOHLC()
			bearishFractalExist = true
			price, reversalType = low, Bearish
		}
		if bullishFractalExist && bearishFractalExist {
			return math.SmallestNonzeroFloat64, Bullish
		}
		if bullishFractalExist || bearishFractalExist {
			return price, reversalType
		}
	}
	return math.SmallestNonzeroFloat64, Bullish
}
