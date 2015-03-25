package cp

type ThemeBlock struct {
	Name string
	Data interface{}
}

type ThemeLayout struct {
	Header,
	Sidebar,
	Footer,
	Content ThemeBlock
}
