package conf

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
}

type DbCommon struct {
	HostPort     string
	User         string
	PWD          string
	DataBaseName string
	Min          int
	Max          int
}
