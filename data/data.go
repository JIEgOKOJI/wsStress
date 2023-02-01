package data

type ParsedConfig struct {
	Addr        string `yaml:"addr"`
	Tokens      string `yaml:"tokens"`
	Connections int    `yaml:"connections"`
	Sendmsg     bool   `yaml:"sendmsg"`
}
type Tokens struct {
	Data map[string]string
}
