package config

type Config struct {
	Listen  Listen
	Elastic Elastic
}

type Listen struct {
	Port string
}

type Elastic struct {
	Host          string
	Port          string
	IndexName     string
	MetaDataIndex string
}
