package gene

type gene interface {
	Act() (ok bool, err error)
	Code() string
}
