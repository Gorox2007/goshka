package main

import "fmt"

type PepeSchnele struct {
	Speed    int
	Charisma int
	Wisdom   int
}

func NewPepeSchnele(speed, charisma, wisdom int) *PepeSchnele {
	return &PepeSchnele{
		Speed:    speed,
		Charisma: charisma,
		Wisdom:   wisdom,
	}
}

func (p *PepeSchnele) GetRating() int {
	return (p.Speed * 2) + (p.Charisma * 3) + p.Wisdom
}

func main() {
	p1 := NewPepeSchnele(1, 2, 3)
	p2 := NewPepeSchnele(4, 5, 6)

	fmt.Printf("Пепе Шнеле 1 [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d\n", p1.Speed, p1.Charisma, p1.Wisdom, p1.GetRating())
	fmt.Printf("Пепе Шнеле 2 [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d\n", p2.Speed, p2.Charisma, p2.Wisdom, p2.GetRating())

	if p1.GetRating() > p2.GetRating() {
		fmt.Printf("p1 круче\n")
	} else if p1.GetRating() == p2.GetRating() {
		fmt.Printf("Одинаково крутые\n")
	} else {
		fmt.Printf("p2 круче\n")
	}
}
