package Serialization

type Info struct {
	Version string `yaml:"version"`
}

type Openapi struct {
	Info Info `yaml:"info"`
}

type ServerInfo struct {
	Url string `yaml:"url"`
}

type Servers struct {
	Info []ServerInfo `yaml:"servers"`
}
