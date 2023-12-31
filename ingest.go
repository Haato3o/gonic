package gonic

type Ingestable interface {
	Push()
	Pop()
	Count()
	Ping()
}
