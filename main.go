package main

import (
	"encoding/json"
	"bufio"
	"fmt"
	"io"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	info, err := os.Stdin.Stat()
    if err != nil {
        panic(err)
    }

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: yaml2json < example.json > example.yaml")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	var obj interface{}
	err = json.Unmarshal([]byte(string(output)), &obj)
	if err != nil {
		fmt.Printf("Error unmarshalling input. Is it valid JSON? %s: %v\n", string(output), err)
		os.Exit(-1)
	}

	if obj != nil {
		yamlBytes, err := yaml.Marshal(obj)
		if err != nil {
			fmt.Printf("Error marshaling into YAML from JSON %s: %v\n", string(output), err)
			os.Exit(-1)
		}

		fmt.Println(string(yamlBytes))
	}

}
