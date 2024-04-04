package service

import (
	"encoding/json"
	inter "millieMind/internal/questions/interface"
)

func Questions(data inter.Resp) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	var response inter.Resp
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		panic(err)
	}

	for category, questions := range response {
		println("Category:", category)
		for _, questions := range questions {
			println("Age:", questions.Age)
			for _, p := range questions.Question {
				println("Number:", p.Number)
				println("Desc:", p.Desc)
			}
		}
	}

}
