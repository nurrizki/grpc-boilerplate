package db

type DatabaseList struct {
	MySQL struct {
		Sdb         DatabaseConfig
		MiniProject DatabaseConfig
	}
	Redis struct {
		DB DatabaseConfig
	}
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Adapter  string
}
