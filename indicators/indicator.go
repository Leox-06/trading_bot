package indicator

type Candle struct {
	Open   float32 `json:"open"`
	High   float32 `json:"high"`
	Low    float32 `json:"low"`
	Close  float32 `json:"close"`
	Volume float32 `json:"volume"`
	Time   float32 `json:"time"`
}

type Graph struct {
	Candles  []Candle
	Interval int
	Limit    int
}

func (graph Graph) ATR(length int) (ATR []float32) {
	ATR = make([]float32, graph.Limit)
	prev_candle := Candle{0, 0, 0, 0, 0, 0}
	prev_ATR := float32(0)

	for n := graph.Limit - 1; n >= 0; n-- {
		curr_candle := graph.Candles[n]

		// TR := max((curr_candle.High - curr_candle.Low), math.Abs(curr_candle.High - prev_candle.Close), math.Abs(curr_candle.Low - prev_candle.Close))
		TR := max(curr_candle.High, prev_candle.Close) - min(curr_candle.Low, prev_candle.Close)

		prev_ATR = (prev_ATR*(float32(length)-1) + TR) / float32(length)
		prev_candle = curr_candle
		ATR[n] = prev_ATR
	}

	return
}

func (graph Graph) SuperTrend(length int, factor float32) (SuperTrend []float32, direction []int) {
	SuperTrend = make([]float32, graph.Limit)

	direction = make([]int, graph.Limit)

	prev_SP_upper_band := float32(3.4e+38)
	prev_SP_lower_band := float32(0)
	prev_SuperTrend := float32(0)
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
