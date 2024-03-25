package main

type Config struct {
	Version string
	Plugins []string
	Stat    map[string]int
}

func (cfg *Config) Clone() *Config {
	clone := &Config{
		Version: cfg.Version,
		Stat:    make(map[string]int, len(cfg.Stat)),
	}
	for k, v := range cfg.Stat {
		clone.Stat[k] = v
	}
	for _, v := range cfg.Plugins {
		clone.Plugins = append(clone.Plugins, v)
	}
	return clone
}
