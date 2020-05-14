package main

import (
	"io/ioutil"

	"github.com/kpango/glg"
)

func main() {
	data, err := ioutil.ReadFile("configs.yaml")
	config := Config{}
	config.Load(data)
	check("Error loading file: %s", err)
}

func check(message string, e error) {
	if e != nil {
		glg.Errorf(message, e.Error())
	}
}
