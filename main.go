package main

import "fmt"

type Soldier struct {
	name string
	rank string
}

func (s *Soldier) showForcePersonalName() {
	fmt.Println(s.name)
}

func (s *Soldier) showForcePersonalRank() {
	fmt.Println(s.rank)
}

type Police struct {
	name string
	rank string
}

func (p *Police) showForcePersonalName() {
	fmt.Println(p.name)
}

func (p *Police) showForcePersonalRank() {
	fmt.Println(p.rank)
}

type ForcePersonel interface {
	showForcePersonalName()
	showForcePersonalRank()
}

func showForcePersonelDetail(forcePersonel ForcePersonel) {
	forcePersonel.showForcePersonalName()
}

func main() {
	// kc := Soldier{
	// 	name: "kc",
	// 	rank: "major",
	// }

	lucky := Police{
		name: "luckky",
		rank: "inpsector",
	}

	showForcePersonelDetail(&lucky)
}
