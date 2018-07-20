package config

type ApiConf struct {
	Host string `default:"127.0.0.1"`
	Port int    `default:"8000"`
}

type RedisConf struct {
	Host string `default:"192.168.7.211"`
	Port int    `default:"6379"`
	DB   int    `default:"2"`
}

type MysqlConf struct {
	Host      string `default:"192.168.7.119"`
	Port      int    `default:"3306"`
	User      string `default:"root"`
	Pass      string `default:"123456"`
	Name      string `default:"ppgo_api_demo_gin"`
	Charset   string `default:"utf8"`
	IdleConns int    `default:"200"`
	OpenConns int    `default:"500"`
}
