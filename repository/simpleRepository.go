package repository

type SimpleRepository struct {
	Error bool
}

func NewRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}
