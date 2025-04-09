package playerdata

type TeamIsFullError struct {
	Message string
}

func (t TeamIsFullError) Error() string {
	return t.Message
}
