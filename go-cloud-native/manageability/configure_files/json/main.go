package main

import (
	"encoding/json"
	"fmt"
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

	fmt.Println(string(bytes))

	bytes, _ = json.MarshalIndent(c, "", "	")

	fmt.Println(string(bytes))

	c = Config{}

	bytes = []byte(`{
		"Host": "127.0.0.1",
		"Port": 1234,
		"Tags": {
			"foo": "bar"
		}
	}`)

	err := json.Unmarshal(bytes, &c)

	if err != nil {
		fmt.Println(c)
	}

	bytes = []byte(`{
		"Host": "127.0.0.1",
		"Food": {
			"foo": "bar"
		}
	}`)

	err = json.Unmarshal(bytes, &c)

	if err != nil {
		fmt.Println(c)
	}
}
