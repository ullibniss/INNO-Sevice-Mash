package endpoint

type Endpoint struct {
	URL    string `yaml:"url"`
	Weight int    `yaml:"weight"`
}
