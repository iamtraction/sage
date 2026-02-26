package llm

type Request struct {
	System       string
	User         string
	Model        string
	OutputSchema map[string]any
}
