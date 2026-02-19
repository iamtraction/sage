package config

func Run(args []string) int {
	switch {
	case len(args) == 0:
		return runShow()
	case len(args) == 1:
		return runGet(args[0])
	default:
		return runSet(args[0], args[1:])
	}
}
