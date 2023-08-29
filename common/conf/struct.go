package conf

// AppConf 服务配置
var AppConf app

type app struct {
	App struct {
		HostPort string
	}
	MySQL DbCommon
	Redis struct {
		HostPort string
		User     string
		Password string
		DB       int
	}
	//Exmail 腾讯企业邮
	Exmail struct {
		Host    string
		Port    int
		From    string
		Account EmailAccountInfo
	}
	Task       map[string]taskInfo
	SiteConfig map[string]ConfigInfo `toml:"conf"`
}

// DbCommon 数据库配置
type DbCommon struct {
	HostPort     string
	User         string
	PWD          string
	DataBaseName string
	Min          int
	Max          int
}

// EmailAccountInfo 邮件账户配置
type EmailAccountInfo struct {
	User     string
	Password string
}

type taskInfo struct {
	Spec   string
	Status string
}

// configInfo 站点基础配置信息
type ConfigInfo struct {
	Title      string
	SubTitle   string
	TitleEn    string
	SubTitleEn string
	CracDesc   string
	Cover      string
}
