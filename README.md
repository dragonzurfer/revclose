# revclose

This is a package that returns signals based on close beyond levels and looks for reversal

  

# Input

- Candles 
	- Candle - OHLC(open, high, low, close)
- Prices
	- It takes in a bunch of price points

# Output
- Finds latest reversal
	- Reversal is defined as a candle high price point which remains the high for previous x candles and post x candles.
		- Example given High values of 5 candles [1,1,4,2,3] since 4 is higher than 1,1, and 2,3 we can say that 4 is a high and this  is a reversal
- Find Prices(input) that are within (Reversal to Latest Close) range if exists. Return reversal point and prices within reversal to close.

# Installation

```
	go get github.com/dragonzurfer/revclose
```
