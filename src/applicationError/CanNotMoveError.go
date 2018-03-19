package applicationError

// Bacterium can not move error
type CanNotMove struct {
}

func (error *CanNotMove) Error() string {
	return "Bacterium can't move"
}
