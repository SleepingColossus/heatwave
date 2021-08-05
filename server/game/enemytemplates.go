package game

func newEnemyMeleeBasic() *Enemy {
	return newEnemy(enemyMeleeBasic, 3, 32, newVector2(1, 1))
}

func newEnemyMeleeFast() *Enemy {
	return newEnemy(enemyMeleeFast, 3, 32, newVector2(3, 3))
}

func newEnemyRangedBasic() *Enemy {
	return newEnemy(enemyRangedBasic, 2, 256, newVector2(2, 2))
}

func newEnemyRangedAdvanced() *Enemy {
	return newEnemy(enemyRangedAdvanced, 2, 512, newVector2(2, 2))
}

func newEnemyTank() *Enemy {
	return newEnemy(enemyTank, 10, 256, newVector2(1, 1))
}
