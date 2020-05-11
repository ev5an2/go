package balance
type balance interface {
	DoBalance([]*Instance)(*Instance,error)
}
