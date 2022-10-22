package config

type Mysql struct {
	// ,inline 意思应该就是继承上一级的注解
	GeneralDB `yaml:",inline"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
