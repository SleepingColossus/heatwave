package game

import (
	"fmt"
	"log"
)

type GameState struct {
	Phase   int
	Players map[string]*Player
	Wave    *Wave
	Projectiles map[string]*Projectile
}

func NewGameState() *GameState {
	return &GameState{
		Phase:   pending,
		Players: make(map[string]*Player),
		Wave:    nil,
		Projectiles: make(map[string]*Projectile),
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
		}
	}
}

func (gs *GameState) startGame() {
	log.Println("game started")

	if gs.Phase == pending {
		gs.Phase = started
		gs.Wave = waveData[0]
		gs.Wave.start(gs.players())
	}
}

func (gs *GameState) Update() GameStateUpdate {
	for _, p := range gs.Players {
		p.update()
	}

	if gs.Wave != nil {
		for _, e := range gs.Wave.Enemies {
			e.update()
		}
	}

	for _, pr := range gs.Projectiles {
		pr.update()
	}

	update := newGameStateUpdate(gs)
	gs.processActorStates()
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

func (gs *GameState) players() []*Player {
	var p []*Player

	for _, player := range gs.Players {
		p = append(p, player)
	}

	return p
}
