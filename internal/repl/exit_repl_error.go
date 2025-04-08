package repl

type ExitReplError struct {
	message string
}

func (e ExitReplError) Error() string {
	return e.message
}
