package handlers

type Handler interface {
	HandleMsg(command string, args []string) (answer string, err error)
}
