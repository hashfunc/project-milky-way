package kakao

type Meta struct {
	TotalCount    int        `json:"total_count"`
	PageableCount int        `json:"pageable_count"`
	IsEnd         bool       `json:"is_end"`
	SameName      RegionInfo `json:"-"`
}

type RegionInfo struct {
	Region         []string `json:"region"`
	Keyword        string   `json:"keyword"`
	SelectedRegion string   `json:"selected_region"`
}
