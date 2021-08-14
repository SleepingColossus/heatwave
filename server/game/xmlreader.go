package game

import (
	"encoding/xml"
	"io/ioutil"
)

func ReadWaveData(path string) []*Wave {
	data, err :=ioutil.ReadFile(path)
	checkErr(err)

	var wt WavesTemplate
	if err := xml.Unmarshal(data, &wt); err != nil {
		checkErr(err)
	}

	return wt.toGameWave()
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}