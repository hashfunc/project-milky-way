package kakao

type Client struct {
	config *Config
}

type Config struct {
	URL string `yaml:"url" config:"required"`
	Key string `yaml:"key" config:"required"`
}

func NewClient(config *Config) *Client {
	return &Client{config: &Config{
		URL: config.URL,
		Key: config.Key,
	}}
}
