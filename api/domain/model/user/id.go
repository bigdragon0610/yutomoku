package user

type ID uint64

func NewID(id uint64) ID {
	return ID(id)
}
