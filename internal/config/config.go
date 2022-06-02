package config

type Config struct {
	Spreadsheet struct {
		Id     string `yaml:"id"`
		Ranges struct {
			Engineers      string `yaml:"engineers"`
			Holidays       string `yaml:"holidays"`
			Schedule       string `yaml:"schedule"`
			ScheduleInsert string `yaml:"scheduleinsert"`
		} `yaml:"ranges"`
	} `yaml:"spreadsheet"`
	Credentials struct {
		Path string `yaml:"path"`
	} `yaml:"credentials"`
	Debug bool `yaml:"debug"`
}
