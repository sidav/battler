package battler

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func unmarshalAllUnitsData() {
	dirname := "units_data"
	items, _ := ioutil.ReadDir(dirname)
	for _, currUnitDataFile := range items {
		if currUnitDataFile.IsDir() {
		} else {
			currUnitsList := make(map[string]*UnitData)
			read, _ := os.ReadFile(dirname + "/" + currUnitDataFile.Name())
			json.Unmarshal(read, &currUnitsList)

			//fmt.Printf("%v\n", string(read))
			//json, _ := json.MarshalIndent(currUnitsList, "", "\t")
			//fmt.Printf("%v\n", string(json))

			for k, v := range currUnitsList {
				UNITS_DATA[k] = v
				UNITS_DATA[k].normalize(k)
			}
		}
	}
}

func marshalAndSaveUnitsData(data map[string]*UnitData, name string) {
	// open output file
	fo, err := os.Create("units_data/" + name + ".json")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	json, _ := json.MarshalIndent(data, "", "\t")
	if _, err := fo.Write(json); err != nil {
		panic(err)
	}
}
