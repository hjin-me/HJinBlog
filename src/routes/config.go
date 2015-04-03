package routes

type config struct {
	Routes []r `yaml:"routes"`
}
type r struct {
	Rule   string   `yaml:"rule"`
	Method []string `yaml:"method"`
	Action string   `yaml:"action"`
}
