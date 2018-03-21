package applicationError

// CanNotMove error: Microorganism can not move
type CanNotMove struct {
}

func (error *CanNotMove) Error() string {
	return "Microorganism can't move"
}
