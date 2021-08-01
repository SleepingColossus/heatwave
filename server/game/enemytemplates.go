package game

func newEnemyMeleeBasic() *Enemy {
	return newEnemy(enemyMeleeBasic, 3, 1)
}

func newEnemyMeleeFast() *Enemy {
	return newEnemy(enemyMeleeFast, 3, 3)
}

func newEnemyRangedBasic() *Enemy {
	return newEnemy(enemyRangedBasic, 2, 2)
}

func newEnemyRangedAdvanced() *Enemy {
	return newEnemy(enemyRangedAdvanced, 2, 2)
}

func newEnemyTank() *Enemy {
	return newEnemy(enemyTank, 10, 1)
}
