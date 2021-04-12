package collection

type Arguments interface {
	GetArgument(name string) (string, error)

	Import([]string)
	Export() []string
}
