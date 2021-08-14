package game

import (
	"fmt"
	"log"
)

type GameState struct {
	Settings    *GameSettings
	Phase       int
	Players     map[string]*Player
	CurrentWave int
	Wave        *Wave
	Projectiles map[string]*Projectile
	Tick        int
}

func NewGameState(gs *GameSettings) *GameState {
	return &GameState{
		Settings:    gs,
		Phase:       pending,
		Players:     make(map[string]*Player),
		CurrentWave: 0,
		Wave:        nil,
		Projectiles: make(map[string]*Projectile),
		Tick:        0,
	}
}

func (gs *GameState) IsActive() bool {
	return gs.Phase == started
}

func (gs *GameState) AddPlayer(id string) GameStateUpdate {
	player := newPlayer(id)
	gs.Players[player.Id] = player

	if len(gs.Players) > 0 && gs.Phase == pending {
		gs.startGame()
	}

	log.Printf("player has joined: %s\n", player.Id)

	return newGameStateUpdate(gs)
}

func (gs *GameState) MarkPlayerForDeletion(id string) {
	gs.Players[id].State = actorDeleted
}

// delete actors marked for removal
// progress from "created" to "updated" state for next tick
func (gs *GameState) processActorStates() {
	// players
	var deletedPlayers []string

	for _, player := range gs.Players {
		if player.State == actorCreated {
			player.State = actorUpdated
		} else if player.State == actorDeleted {
			deletedPlayers = append(deletedPlayers, player.Id)
		}
	}

	for _, playerId := range deletedPlayers {
		delete(gs.Players, playerId)

		log.Printf("player has left: %s\n", playerId)
	}

	// end game if all players have left
	if len(gs.Players) == 0 {
		// TODO re-enable
		//gs.Phase = over
	}

	// projectiles
	var deletedProjectiles []string

	for _, projectile := range gs.Projectiles {
		if projectile.State == actorCreated {
			projectile.State = actorUpdated
		} else if projectile.State == actorDeleted {
			deletedProjectiles = append(deletedProjectiles, projectile.Id)
		}
	}

	for _, projectileId := range deletedProjectiles {
		delete(gs.Projectiles, projectileId)

		log.Printf("projectile deleted: %s\n", projectileId)
	}

	// enemies
	if gs.Wave != nil {
		var deletedEnemies []string

		for _, enemy := range gs.Wave.Enemies {
			if enemy.State == actorCreated {
				enemy.State = actorUpdated
			} else if enemy.State == actorDeleted {
				deletedEnemies = append(deletedEnemies, enemy.Id)
			}
		}

		for _, enemyId := range deletedEnemies {
			delete(gs.Wave.Enemies, enemyId)

			log.Printf("enemy deleted: %s\n", enemyId)
		}
	}
}

func (gs *GameState) startGame() {
	log.Println("game started")

	if gs.Phase == pending {
		gs.Phase = started
		gs.CurrentWave = 1
		gs.Wave = gs.Settings.WaveSettings[gs.CurrentWave - 1]
		gs.Wave.start(gs.getPlayers())
	}
}

func (gs *GameState) startNextWave() {
	log.Println("starting next wave")

	// is this the last wave?
	if gs.CurrentWave == len(gs.Settings.WaveSettings) {
		// yes - end the game
		gs.endGame("victory")
	} else {
		// no - start next wave
		gs.CurrentWave++
		gs.Wave = gs.Settings.WaveSettings[gs.CurrentWave - 1]
		gs.Wave.start(gs.getPlayers())
	}
}

func (gs *GameState) endGame(reason string) {
	log.Println("game over:", reason)

	gs.Phase = over
	gs.Wave.State = over
}

func (gs *GameState) Update() GameStateUpdate {
	for _, p := range gs.Players {
		p.update()
	}

	if gs.Wave != nil {
		for _, e := range gs.Wave.Enemies {
			newProjectile := e.update()

			if newProjectile != nil {
				gs.Projectiles[newProjectile.Id] = newProjectile
			}
		}
	}

	for _, pr := range gs.Projectiles {
		var enemies map[string]*Enemy

		if gs.Wave != nil {
			enemies = gs.Wave.Enemies
		} else {
			enemies = make(map[string]*Enemy)
		}

		pr.update(gs.Players, enemies)
	}

	// is the wave over?
	if (gs.Phase != over) && (gs.Wave != nil) && (len(gs.Wave.Enemies) == 0) {
		gs.startNextWave()
	}

	// are all players dead?
	if len(gs.Players) == 0 {
		gs.endGame("defeat")
	}

	update := newGameStateUpdate(gs)
	gs.processActorStates()
	gs.Tick++
	return update
}

func (gs *GameState) PlayerMove(id string, dirX, dirY int) error {
	player, ok := gs.Players[id]

	if !ok {
		return fmt.Errorf("player %s not found", id)
	}

	player.setDirection(dirX, dirY)

	return nil
}

func (gs *GameState) PlayerShoot(playerId string) error {
	player, ok := gs.Players[playerId]

	if !ok {
		return fmt.Errorf("player %s not found", playerId)
	}

	projectile := newFriendlyProjectile(player.Body2D)
	gs.Projectiles[projectile.Id] = projectile

	return nil
}

func (gs *GameState) getPlayers() []*Player {
	var p []*Player

	for _, player := range gs.Players {
		p = append(p, player)
	}

	return p
}
