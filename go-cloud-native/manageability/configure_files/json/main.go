package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host string
	Port uint16
	Tags map[string]string
}

func main() {
	c := Config{
		Host: "localhost",
		Port: 1313,
		Tags: map[string]string{"env": "dev"},
	}

	bytes, _ := json.Marshal(c)

	bytesYaml, _ := yaml.Marshal(c)

	fmt.Println(string(bytes))
	fmt.Println(string(bytesYaml))

	bytes, _ = json.MarshalIndent(c, "", "	")

	fmt.Println(string(bytes))

	c2 := Config{}

	bytes = []byte(`{
		"Host": "127.0.0.1",
		"Port": 1234,
		"Tags": {
			"foo": "bar"
		}
	}`)

	err := json.Unmarshal(bytes, &c2)

	if err == nil {
		fmt.Println(c2)
	}

	c3 := Config{}

	bytes = []byte(`{
		"Host": "127.0.0.1",
		"Food": {
			"foo": "bar"
		}
	}`)

	err = json.Unmarshal(bytes, &c3)

	if err == nil {
		fmt.Println(c3)
	}
}
