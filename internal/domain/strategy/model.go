package strategy

type Strategy struct {
	UUID           string
	Name           string
	Description    string
	Implementation []byte
}

func (s *Strategy) Tic() {

}
