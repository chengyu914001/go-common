package permission

type Operator uint16

const (
	ReadList Operator = 1 << iota
	ReadDetail
	Update
	UpdateDetail
	Create
	Delete
	Use
	
)
