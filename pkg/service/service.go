package service

type Image struct {
	PullPolicy string `yaml:"pullPolicy"`
	Repository string `yaml:"repository"`
	Tag        string `yaml:"tag"`
}

type Microservices map[string]*Microservice

type Microservice struct {
	Image Image `yaml:"image"`
}
