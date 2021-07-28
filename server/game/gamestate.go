package game

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
		gs.Wave = waveData[0]
		gs.Wave.Start()
	}
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
