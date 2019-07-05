package JSON

import (
	"encoding/json"
	"log"
)

type H map[interface{}]interface{}

func (h *H) Stringify() string {
	origin := map[interface{}]interface{}(*h)
	var target = make(map[string]interface{})

	for key, value := range origin {
		if str, ok := key.(string); ok {
			target[str] = value
		}
	}

	jsonStr, err := json.Marshal(target)

	if err != nil {
		log.Fatal(err)
	}

	return string(jsonStr)
}