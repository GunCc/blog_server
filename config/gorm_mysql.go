package config

type Mysql struct {
	// ,inline 意思应该就是继承上一级的注解 后面这个一定要加 要不获取不到
	GeneralDB `yaml:",inline"  mapstructure:",squash"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
