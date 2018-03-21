package applicationError

// Stuck error: Microorganism can not move to the coordinates where another stands
type Stuck struct {
}

func (error *Stuck) Error() string {
	return "Microorganism can't move where another stands"
}
