package game

type GameStateUpdate struct {
	ClientId      string
	Notifications []string
	ActorUpdates  []Actor
}

func newGameStateUpdate(gs *GameState) GameStateUpdate {
	var actors []Actor

	for _, player := range gs.Players {
		actors = append(actors, player.Actor)
	}

	for _, projectile := range gs.Projectiles {
		actors = append(actors, projectile.Actor)
	}

	if gs.Wave != nil {
		for _, enemy := range gs.Wave.Enemies {
			actors = append(actors, enemy.Actor)
		}
	}

	return GameStateUpdate{
		ClientId: "",
		Notifications: make([]string, 0),
		ActorUpdates: actors,
	}
}

func (gsu *GameStateUpdate) AddNotification(message string) {
	gsu.Notifications = append(gsu.Notifications, message)
}

func (gsu *GameStateUpdate) SetClientId(id string) {
	gsu.ClientId = id
}

func (gsu *GameStateUpdate) ClearClientId() {
	gsu.ClientId = ""
}