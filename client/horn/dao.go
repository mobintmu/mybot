package horn

type HornInterface interface {
	Call(message string)
}

type HornStruct struct {
	Url string
}
