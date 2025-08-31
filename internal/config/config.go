package config

import "flag"

type Config struct {
	StorageType string
}

func NewConfig() *Config {
	storageType := flag.String("storage", "in-memory", "storage type (in-memory or postgres)")
	flag.Parse()

	if *storageType != "postgres" && *storageType != "in-memory" {
		panic("invalid storage type: " + *storageType + " , only postgres or in-memory")
	}

	return &Config{
		StorageType: *storageType,
	}
}
