package cli

func Parse(args []string) (Command, error) {
	return NewServerCommand(), nil
}
