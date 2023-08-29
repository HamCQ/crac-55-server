package entities

// Config 站点配置
type Config struct {
	CurrentActiveYear int    `json:"current_active_year"`
	Title             string `json:"title"`
	SubTitle          string `json:"sub_title"`
	TitleEn           string `json:"title_en"`
	SubTitleEn        string `json:"sub_title_en"`
	Cover             string `json:"cover"`
	CracDesc          string `json:"crac_desc"`
}
