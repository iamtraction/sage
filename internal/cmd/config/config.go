package config

func Run(args []string) int {
	switch {
	case len(args) == 1:
		return runGet(args[0])
	default:
		return runShow()
	}
}
