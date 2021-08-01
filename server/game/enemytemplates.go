package game

func newEnemyMeleeBasic() *Enemy {
	return newEnemy(enemyMeleeBasic, 3, 1, 32)
}

func newEnemyMeleeFast() *Enemy {
	return newEnemy(enemyMeleeFast, 3, 3, 32)
}

func newEnemyRangedBasic() *Enemy {
	return newEnemy(enemyRangedBasic, 2, 2, 256)
}

func newEnemyRangedAdvanced() *Enemy {
	return newEnemy(enemyRangedAdvanced, 2, 2, 512)
}

func newEnemyTank() *Enemy {
	return newEnemy(enemyTank, 10, 1, 32)
}
