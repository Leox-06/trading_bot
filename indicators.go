package trading_bot

func (graph Graph) ATR(length int) (ATR []float64) {
	ATR = make([]float64, graph.Limit)
	prev_candle := Candle{0, 0, 0, 0, 0, 0}
	prev_ATR := float64(0)

	for n := graph.Limit - 1; n >= 0; n-- {
		curr_candle := graph.Candles[n]

		// TR := max((curr_candle.High - curr_candle.Low), math.Abs(curr_candle.High - prev_candle.Close), math.Abs(curr_candle.Low - prev_candle.Close))
		TR := max(curr_candle.High, prev_candle.Close) - min(curr_candle.Low, prev_candle.Close)

		prev_ATR = (prev_ATR*(float64(length)-1) + TR) / float64(length)
		prev_candle = curr_candle
		ATR[n] = prev_ATR
	}

	return
}

func (graph Graph) SuperTrend(length int, factor float64) (SuperTrend []float64, direction []int) {
	SuperTrend = make([]float64, graph.Limit)

	direction = make([]int, graph.Limit)

	prev_SP_upper_band := 3.4e+38
	prev_SP_lower_band := 0.0
	prev_SuperTrend := 0.0
	prev_direction := 0

	ATR := graph.ATR(length)
	for n := graph.Limit - 1; n >= 0; n-- {
		curr_candle := graph.Candles[n]
		curr_ATR := ATR[n]

		upper_band := (curr_candle.High+curr_candle.Low)/2 + factor*curr_ATR
		lower_band := (curr_candle.High+curr_candle.Low)/2 - factor*curr_ATR
		mid_band := (curr_candle.Open + curr_candle.Close) / 2

		prev_SP_upper_band = min(prev_SP_upper_band, upper_band)
		prev_SP_lower_band = max(prev_SP_lower_band, lower_band)
		if mid_band > prev_SP_upper_band {
			prev_SP_upper_band = upper_band
			prev_direction = 1
		} else if mid_band < prev_SP_lower_band {
			prev_SP_lower_band = lower_band
			prev_direction = 0
		}

		if prev_direction == 1 {
			prev_SuperTrend = prev_SP_lower_band
		} else {
			prev_SuperTrend = prev_SP_upper_band
		}

		SuperTrend[n] = prev_SuperTrend
		direction[n] = prev_direction
	}

	return
}
