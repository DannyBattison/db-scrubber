package cfg

import (
	"bufio"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*ScrubConfig, error) {
	config := ScrubConfig{}

	f, err := os.Open((path))

	if err != nil {
		return nil, err
	}

	defer f.Close()

	stats, err := f.Stat()
	if err != nil {
		return nil, err
	}

	bytes := make([]byte, stats.Size())

	bufr := bufio.NewReader(f)
	_, err = bufr.Read(bytes)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
