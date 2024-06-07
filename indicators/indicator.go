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

	for n := graph.Limit; n > 0; n-- {
		curr_candle := graph.Candles[n-1]
		// TR := max((curr_candle.High - curr_candle.Low), math.Abs(curr_candle.High - prev_candle.Close), math.Abs(curr_candle.Low - prev_candle.Close))
		TR := max(curr_candle.High, prev_candle.Close) - min(curr_candle.Low, prev_candle.Close)

		prev_ATR = (prev_ATR*(float32(length)-1) + TR) / float32(length)
		prev_candle = curr_candle

		ATR[n-1] = prev_ATR
	}

	return
}
