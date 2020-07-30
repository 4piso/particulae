package iexcloud

// Earnings definition struct
type Earnings struct {
	Symbol   string `json:"symbol"`
	Earnings []earnings
}

// earnings struct of the earning return type
type earnings struct {
	ActualEPS            float64 `json:"actualEPS"`
	ConsensusEPS         float64 `json:"consensusEPS"`
	AnnounceTime         string  `json:"announceTime"`
	NumberOfEstimates    int     `json:"numberOfEstimates"`
	EPSSurpriseDollar    float64 `json:"EPSSurpriseDollar"`
	EPSReportDate        string  `json:"EPSReportDate"`
	FiscalPeriod         string  `json:"fiscalPeriod"`
	FiscalEndDate        string  `json:"fiscalEndDate"`
	YearAgo              float64 `json:"yearAgo"`
	YearAgoChangePercent float64 `json:"yearAgoChangePercent"`
}

// Quote is the data struct for the quote api
type Quote struct {
	Symbol                 string  `json:"symbol"`
	CompanyName            string  `json:"companyName"`
	CalculationPrice       string  `json:"calculationPrice"`
	Open                   int     `json:"open"`
	OpenTime               int64   `json:"openTime"`
	Close                  float64 `json:"close"`
	CloseTime              int64   `json:"closeTime"`
	High                   float64 `json:"high"`
	Low                    float64 `json:"low"`
	LatestPrice            float64 `json:"latestPrice"`
	LatestSource           string  `json:"latestSource"`
	LatestTime             string  `json:"latestTime"`
	LatestUpdate           int64   `json:"latestUpdate"`
	LatestVolume           int     `json:"latestVolume"`
	Volume                 int     `json:"volume"`
	IexRealtimePrice       float64 `json:"iexRealtimePrice"`
	IexRealtimeSize        int     `json:"iexRealtimeSize"`
	IexLastUpdated         int64   `json:"iexLastUpdated"`
	DelayedPrice           float64 `json:"delayedPrice"`
	DelayedPriceTime       int64   `json:"delayedPriceTime"`
	OddLotDelayedPrice     float64 `json:"oddLotDelayedPrice"`
	OddLotDelayedPriceTime int64   `json:"oddLotDelayedPriceTime"`
	ExtendedPrice          float64 `json:"extendedPrice"`
	ExtendedChange         float64 `json:"extendedChange"`
	ExtendedChangePercent  float64 `json:"extendedChangePercent"`
	ExtendedPriceTime      int64   `json:"extendedPriceTime"`
	PreviousClose          float64 `json:"previousClose"`
	PreviousVolume         int     `json:"previousVolume"`
	Change                 float64 `json:"change"`
	ChangePercent          float64 `json:"changePercent"`
	IexMarketPercent       float64 `json:"iexMarketPercent"`
	IexVolume              int     `json:"iexVolume"`
	AvgTotalVolume         int     `json:"avgTotalVolume"`
	IexBidPrice            float64 `json:"iexBidPrice"`
	IexBidSize             int     `json:"iexBidSize"`
	IexAskPrice            float64 `json:"iexAskPrice"`
	IexAskSize             int     `json:"iexAskSize"`
	MarketCap              int64   `json:"marketCap"`
	Week52High             float64 `json:"week52High"`
	Week52Low              float64 `json:"week52Low"`
	YtdChange              float64 `json:"ytdChange"`
	PeRatio                float64 `json:"peRatio"`
	LastTradeTime          int64   `json:"lastTradeTime"`
	IsUSMarketOpen         bool    `json:"isUSMarketOpen"`
}
