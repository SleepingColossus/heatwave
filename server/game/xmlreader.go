package game

import (
	"encoding/xml"
	"io/ioutil"
)

func ReadEnemyData(path string) map[int]EnemyTemplate {
	data, err := ioutil.ReadFile(path)
	checkErr(err)

	var et EnemiesTemplate
	if err := xml.Unmarshal(data, &et); err != nil {
		checkErr(nil)
	}

	etMap := make(map[int]EnemyTemplate)

	for _, enemyTemplate := range et.Enemies {
		etMap[enemyTemplate.Type] = enemyTemplate
	}

	return etMap
}

func ReadWaveData(path string, enemyTemplates map[int]EnemyTemplate) []*Wave {
	data, err :=ioutil.ReadFile(path)
	checkErr(err)

	var wt WavesTemplate
	if err := xml.Unmarshal(data, &wt); err != nil {
		checkErr(err)
	}

	return wt.toGameWave(enemyTemplates)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}