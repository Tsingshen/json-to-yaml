package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	yaml := flag.String("y", "", "special yaml file path")
	json := flag.String("j", "", "special json file path")
	flag.Parse()

	// json to yaml
	if *yaml != "" {
		y, err := ioutil.ReadFile(*yaml)
		if err != nil {
			fmt.Printf("从文件 %s 中读取 yaml 错误:\n\t%v\n", *yaml, err)
			os.Exit(1)
		}
		str, err := yamlToJson(y)
		if err != nil {
			fmt.Printf("yaml to json err: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(str)
	}
	// yaml to json
	if *json != "" {
		j, err := ioutil.ReadFile(*json)
		if err != nil {
			fmt.Printf("从文件 %s 中读取 json 错误:\n\t%v\n", *json, err)
			os.Exit(1)
		}
		str, err := jsonToYaml(j)
		if err != nil {
			fmt.Printf("json to yaml err: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(str)

	}

}

func jsonToYaml(b []byte) (string, error) {
	y, err := yaml.JSONToYAML(b)
	if err != nil {
		return "", err
	}
	return string(y), nil
}

func yamlToJson(b []byte) (string, error) {
	j, err := yaml.YAMLToJSON(b)

	if err != nil {
		return "", err
	}
	return string(j), nil

}
