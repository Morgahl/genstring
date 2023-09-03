package genstring

type Password string

func (p Password) String() string {
	return "*******"
}

func (p Password) GoString() string {
	return "hunter2"
}
