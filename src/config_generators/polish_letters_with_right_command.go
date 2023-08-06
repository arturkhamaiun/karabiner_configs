package config_generators

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"unicode"
)

func main() {
	rules := map[string]interface{}{
		"title": "0Polish letters with right command (ąćęłńóśźż)",
		"rules": []map[string]interface{}{
			{
				"description":  "0Polish letters with right command (ąćęłńóśźż)",
				"manipulators": getManipulators(),
			},
		},
	}

	jsonData, err := marshalOrderedJSON(rules)
	panicIfErr(err)

	panicIfErr(ioutil.WriteFile("./configs/polish_letters_with_right_command.json", []byte(jsonData), 0644))

	fmt.Println("Finished.")
}

func getManipulators() []map[string]interface{} {
	manipulators := []map[string]interface{}{}
	englishLettersTranslatableToPolish := []string{
		"a",
		"c",
		"e",
		"l",
		"n",
		"o",
		"s",
		"z",
		"x",
	}

	for _, englishLetter := range englishLettersTranslatableToPolish {
		manipulator := createPolishLetterWithRightCommandManipulator(englishLetter)
		manipulators = append(manipulators, manipulator)
	}

	for r := rune('a'); r <= rune('z'); r++ {
		if !unicode.IsLetter(r) {
			continue
		}

		var manipulator map[string]interface{}

		asciiLetter := string(r)

		if contains(englishLettersTranslatableToPolish, asciiLetter) {
			manipulator = createPolishLetterWithRightCommandManipulator(asciiLetter)
		} else {
			manipulator = createDisableRightCommandManipulatorManipulator(asciiLetter)
		}

		manipulators = append(manipulators, manipulator)
	}

	return manipulators
}

func createPolishLetterWithRightCommandManipulator(englishLetter string) map[string]interface{} {
	return map[string]interface{}{
		"type": "basic",
		"conditions": []map[string]interface{}{
			{
				"type": "input_source_if",
				"input_sources": []map[string]interface{}{
					{
						"input_source_id": "com.apple.keylayout.PolishPro",
					},
				},
			},
		},
		"from": map[string]interface{}{
			"key_code": englishLetter,
			"modifiers": map[string]interface{}{
				"mandatory": []string{"right_command"},
				"optional":  []string{"shift", "caps_lock"},
			},
		},
		"to": []map[string]interface{}{
			{
				"key_code":  englishLetter,
				"modifiers": []string{"option"},
			},
		},
	}
}

func createDisableRightCommandManipulatorManipulator(keyCode string) map[string]interface{} {
	return map[string]interface{}{
		"type": "basic",
		"from": map[string]interface{}{
			"key_code": keyCode,
			"modifiers": map[string]interface{}{
				"mandatory": []string{"right_command"},
			},
		},
		"to": []map[string]interface{}{
			{
				"key_code": keyCode,
			},
		},
	}
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func marshalOrderedJSON(data interface{}) ([]byte, error) {
	type kv struct {
		Key   string
		Value interface{}
	}

	var kvs []kv
	m := data.(map[string]interface{})
	for k, v := range m {
		kvs = append(kvs, kv{k, v})
	}

	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Key < kvs[j].Key
	})

	newMap := make(map[string]interface{})
	for _, kv := range kvs {
		newMap[kv.Key] = kv.Value
	}

	return json.MarshalIndent(newMap, "", "  ")
}
