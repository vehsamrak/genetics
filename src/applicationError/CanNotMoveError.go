package applicationError

// CanNotMove error: Bacterium can not move
type CanNotMove struct {
}

func (error *CanNotMove) Error() string {
	return "Bacterium can't move"
}
