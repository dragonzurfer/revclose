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
	Buy  SignalValue = "BUY"
	Sell SignalValue = "SELL"
)

type Signal struct {
	Reversal          float64
	Close             float64
	Value             SignalValue
	LevelsWithinRange []LevelInterface
}

//Two inputs, Candle data (OHLC), Prices (levels)
func GetSignal(candles RevCloseCandles, levels []LevelInterface, close float64) Signal {
	return Signal{Reversal: 17612.5, Close: 17631.9, Value: Buy, LevelsWithinRange: levels}
}
