package model



/*type AdsConfig struct {
	ID  int  `json:"ID"`
	Ads Ads    `json:"ads"`
}*/

type Ads struct {
	ID  int  `json:"ID"`
	Category    [] Category     `json:"category"`
	SubCategory  []SubCategory  `json:"subCategory"`
	bottomBanner BottomBanner `json:"bottomBanner"`
	fullScreen   FullScreen   `json:"fullScreen"`
}

type Category struct {
	ID          string    `json:"id"`
	CategoryAds []AdsItem `json:"categoryAds"`
}

type SubCategory struct {
	ID          string    `json:"id"`
	CategoryAds []AdsItem `json:"categoryAds"`
}

type BottomBanner struct {
	AdsID       string `json:"adsId"`
	RequestType string `json:"requestType"`
}

type FullScreen struct {
	AdsID         string `json:"adsId"`
	RequestCount  string `json:"requestCount"`
	ShowInDetail  string `json:"showInDetail"`
	ShowInListing string `json:"showInListing"`
}

type AdsItem struct {
	Position int `json:"position"`
	AdsID    string `json:"adsId"`
	Type     string `json:"type"`
}


///////////////////////////
type TestItem struct {
	ID int
	Year   int
	Title  string
	Plot   string
	Rating float64
}