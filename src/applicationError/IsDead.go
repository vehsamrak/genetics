package applicationError

// IsDead error: Microorganism can't process actions after death
type IsDead struct {
}

func (error *IsDead) Error() string {
	return "Microorganism can't process actions after death"
}
