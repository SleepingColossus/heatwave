package game

type GameSettings struct {
	WaveSettings []*Wave
}

func NewGameSettings(wavesPath string) *GameSettings {
	waves := ReadWaveData(wavesPath)

	return &GameSettings{
		waves,
	}
}
