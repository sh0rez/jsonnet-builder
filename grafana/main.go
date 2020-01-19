package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	j "github.com/sh0rez/jsonnet-builder"
)

func main() {
	var dash Dashboard
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(data, &dash); err != nil {
		log.Fatalln(err)
	}

	o := j.Add("data",
		j.Import("", "grafonnet/grafana.libsonnet"),
		j.Object("",
			dash.Jsonnet("dash"),
		))

	d := j.Doc{
		Locals: []j.LocalType{
			j.Local(o),
		},
		Root: j.Ref("", "data.dash"),
	}

	fmt.Println(d.String())
}
