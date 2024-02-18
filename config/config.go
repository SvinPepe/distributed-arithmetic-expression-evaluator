package config

type Config struct {
	Orchestrator struct {
		Version string `yaml:"version"`
		Port    int    `yaml:"port"`
	} `yaml:"app"`
	Demon struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"demon"`
}
