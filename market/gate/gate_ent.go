package gate

type Depth struct {
	Time    int    `json:"time"`
	TimeMs  int64  `json:"time_ms"`
	Channel string `json:"channel"`
	Event   string `json:"event"`
	Result  struct {
		T            int64      `json:"t"`
		LastUpdateId int        `json:"lastUpdateId"`
		S            string     `json:"s"`
		Bids         [][]string `json:"bids"`
		Asks         [][]string `json:"asks"`
	} `json:"result"`
}
