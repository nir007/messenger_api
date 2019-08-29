package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type config struct {
	Name   string
	Params map[interface{}]interface{}
}

// Get returns config by name
func Get(name string) (map[interface{}]interface{}, error) {
	for _, v := range conf {
		if v.Name == name {
			return v.Params, nil
		}
	}

	return nil, errors.New("undefined config for: " + name)
}

var conf []config

func init() {
	executingMode := gin.Mode()
	isRelease := executingMode == "release"

	root, _ := os.Getwd()
	dir, err := os.Open(root + "/config")
	files, _ := dir.Readdir(0)

	for _, f := range files {
		if f.IsDir() || !strings.Contains(f.Name(), ".yaml") {
			continue
		}

		if isRelease && strings.Contains(f.Name(), "-local.yaml") {
			continue
		}

		if !isRelease && !strings.Contains(f.Name(), "-local.yaml") {
			continue
		}

		name := strings.Replace(f.Name(), "-local.yaml", "", -1)
		name = strings.Replace(name, ".yaml", "", -1)

		c := &config{Name: name, Params: make(map[interface{}]interface{}, 0)}

		content, _ := ioutil.ReadFile(root + "/config/" + f.Name())

		err = yaml.Unmarshal(content, &c.Params)

		conf = append(conf, *c)

		if err != nil {
			panic(err)
		}
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("CONFIGS: ", conf)

}
