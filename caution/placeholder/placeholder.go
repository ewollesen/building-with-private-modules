package placeholder

type Thinger struct{}

func NewThinger() *Thinger {
	return &Thinger{}
}

func (t *Thinger) Thing() string {
	return "placeholder"
}
