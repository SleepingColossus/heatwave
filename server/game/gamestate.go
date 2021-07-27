package game

type GameState struct {
	Phase   int
	Players map[string]*Actor
	Wave    *Wave
}

func NewGameState() *GameState {
	return &GameState{
		Phase:   pending,
		Players: make(map[string]*Actor),
		Wave:    waveData[0],
	}
}

func (gs *GameState) AddPlayer(player *Actor) {
	gs.Players[player.Id] = player

	if len(gs.Players) > 0 && gs.Phase == pending {
		gs.Phase = ready
	}
}

func (gs *GameState) DeletePlayer(id string) {
	delete(gs.Players, id)

	if len(gs.Players) == 0 {
		gs.Phase = over
	}
}

func (gs *GameState) StartGame() {
	if gs.Phase == ready {
		gs.Phase = started
		gs.Wave.Start()
	}
}
