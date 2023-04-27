package revclose

type RevCandle interface {
	GetOHLC() (float64, float64, float64, float64)
}

type RevCloseCandles interface {
	GetCandle(int) RevCandle
	GetCandlesLength() int
}

type LevelInterface interface {
	//reversal,close,level
	HasCrossed(float64, float64, interface{}) bool
}

type SignalValue string

const (
	Buy     SignalValue = "BUY"
	Sell    SignalValue = "SELL"
	Nuetral SignalValue = "NUETRAL"
)

type Signal struct {
	Reversal          float64
	Close             float64
	Value             SignalValue
	LevelsWithinRange []LevelInterface
}

//Two inputs, Candle data (OHLC), Prices (levels)
func GetSignal(candles RevCloseCandles, levels []LevelInterface, close float64) Signal {
	var signal Signal
	//get reversal of RevCloseCandles, which is william fractal of 5 candles
	reversal := getLatestReversal(candles)

	//check for hasCrossed within range (reversal to close) and add all within to signal.LevelsWithinRange
	crossedLevels := GetLevelsCrossedInRange(reversal, close, levels)

	//generate Signal
	signalValue := GetSignalValue(reversal, close, len(crossedLevels))

	signal.Reversal = reversal
	signal.Value = signalValue
	signal.Close = close
	signal.LevelsWithinRange = crossedLevels

	return signal
}
