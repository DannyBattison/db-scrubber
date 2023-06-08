package cfg

type ScrubConfig struct {
	Connection connectionConfig `yaml:"connection"`
	Scrub      scrubDataConfig  `yaml:"scrub"`
}

type connectionConfig struct {
	Dsn string `yaml:"dsn"`
}

type scrubDataConfig struct {
	Tables map[string]tableConfig `yaml:"tables"`
}

type tableConfig struct {
	Columns map[string]columnConfig `yaml:"columns"`
}

type columnConfig struct {
	Generate     string   `yaml:"generate"`
	ReferencedBy []string `yaml:"referencedBy"`
	Pattern      string   `yaml:"pattern"`
}
