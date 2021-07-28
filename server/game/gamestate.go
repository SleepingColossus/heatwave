package game

import (
	"log"
)

type GameState struct {
	Phase   int
	Players map[string]*Actor
	Wave    *Wave
	Channel chan interface{}
}

func NewGameState() *GameState {
	return &GameState{
		Phase:   pending,
		Players: make(map[string]*Actor),
		Wave:    nil,
	}
}

func (gs *GameState) IsPending() bool {
	return gs.Phase == pending
}

func (gs *GameState) AddPlayer(player *Actor) {
	gs.Players[player.Id] = player

	// TODO
	//if len(gs.Players) > 0 && gs.Phase == pending {
	//	gs.StartGame()
	//}

	log.Printf("player has joined: %s\n", player.Id)
}

func (gs *GameState) DeletePlayer(id string) {
	delete(gs.Players, id)

	if len(gs.Players) == 0 {
		gs.Phase = over
	}

	log.Printf("player has left: %s\n", id)
}

func (gs *GameState) StartGame() []map[string]string {
	if gs.Phase == pending {
		gs.Phase = started
		gs.Wave = waveData[0]
		gs.Wave.Start(gs.players())
	}

	log.Println("game started")

	return gs.toEnemyMap()
}

func (gs *GameState) Update() {
	for _, p := range gs.Players {
		p.Update()
	}

	if gs.Wave != nil {
		for _, e := range gs.Wave.Enemies {
			e.Update()
		}
	}
}

// returns map of all actors
// used after calls to update()
func (gs *GameState) ToActorMap() []map[string]string {
	var m []map[string]string

	for _, p := range gs.Players {
		m = append(m, p.toMap())
	}

	if gs.Wave != nil {
		for _, e := range gs.Wave.Enemies {
			m = append(m, e.toMap())
		}
	}

	return m
}

// returns map of all enemies
// used after new wave spawns
func (gs *GameState) toEnemyMap() []map[string]string {
	var m []map[string]string

	if gs.Wave != nil {
		for _, e := range gs.Wave.Enemies {
			m = append(m, e.toMap())
		}
	}

	return m
}

func (gs *GameState) players() []*Actor {
	var p []*Actor

	for _, player := range gs.Players {
		p = append(p, player)
	}

	return p
}
