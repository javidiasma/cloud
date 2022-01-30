package DAO

type History struct {
	Rank         int     `json:"rank"`
	Name         string  `json:"name"`
	Platform     string  `json:"platform"`
	Year         int     `json:"year"`
	Genre        string  `json:"genre"`
	Publisher    string  `json:"publisher"`
	NA_Sales     float32 `json:"na_sales"`
	EU_Sales     float32 `json:"eu_sales"`
	JP_Sales     float32 `json:"jp_sales"`
	Other_Sales  float32 `json:"other_sales"`
	Global_Sales float32 `json:"global_sales"`
}
