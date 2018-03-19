package applicationError

type CanNotMove struct {
}

func (error *CanNotMove) Error() string {
	return "Bacterium can't move"
}
