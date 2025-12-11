package stats

type GetStatsRespose struct {
	Period string `json:"period"`
	Sum    int    `json:"sum"`
}
