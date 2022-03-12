package simpleDB

import "fmt"

type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                             // 服务器地址:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                       // 高级配置
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`                     // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                 // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                 // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`                  // 是否开启Gorm全局日志
	LogZap       string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`                     // 是否通过zap写入日志文件
}

func (this *Mysql) Dsn() string {

	//dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	return this.Username + ":" + this.Password + "@tcp(" + this.Path + ")/" + this.Dbname + "?" + this.Config
}

type Postgresql struct {
	Host                 string `mapstructure:"host" json:"host" yaml:"host"`
	Port                 string `mapstructure:"port" json:"port" yaml:"port"`
	Config               string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname               string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username             string `mapstructure:"username" json:"username" yaml:"username"`
	Password             string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns         int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns         int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	PreferSimpleProtocol bool   `mapstructure:"prefer-simple-protocol" json:"preferSimpleProtocol" yaml:"prefer-simple-protocol"`
	Logger               bool   `mapstructure:"logger" json:"logger" yaml:"logger"`
	SingularTable        bool   `mapstructure:"singular-table" json:"singularTable" yaml:"singular-table"`
	TablePrefix          string `mapstructure:"table-prefix" json:"tablePrefix" yaml:"table-prefix"`
}

func (this *Postgresql) Dsn() string {
	//dsn := "host=" + p.Host + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " " + p.Config
	//return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	return fmt.Sprintf(`host=%s  user=%s  password=%s  dbname=%s port=%s %s`,
		this.Host,
		this.Username,
		this.Password,
		this.Dbname,
		this.Port,
		this.Config,
	)
}

type Sqlite struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	Logger       bool   `mapstructure:"logger" json:"logger" yaml:"logger"`
}

type Sqlserver struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	Logger       bool   `mapstructure:"logger" json:"logger" yaml:"logger"`
}
