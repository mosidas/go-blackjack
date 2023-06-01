package object

type Player struct {
	Name  string
	Money int
	Hand  []Hand
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

func (p *Player) AddHand(hand Hand) {
	p.Hand = append(p.Hand, hand)
}
