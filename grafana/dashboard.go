package main

import (
	"fmt"
	"sort"

	j "github.com/sh0rez/jsonnet-builder"
)

type Dashboard struct {
	Annotations   Annotations `json:"annotations"`
	Description   string      `json:"description"`
	Editable      bool        `json:"editable"`
	GraphTooltip  int         `json:"graphTooltip"`
	ID            int         `json:"id"`
	UID           string      `json:"uid"`
	Links         []Link      `json:"links"`
	Panels        Panels      `json:"panels"`
	Refresh       string      `json:"refresh"`
	SchemaVersion int         `json:"schemaVersion"`
	Style         string      `json:"style"`
	Tags          []string    `json:"tags"`
	Time          Time        `json:"time"`
	Timepicker    Timepicker  `json:"timepicker"`
	Timezone      string      `json:"timezone"`
	Title         string      `json:"title"`
	Version       int         `json:"version"`
}

func (d Dashboard) Jsonnet(name string) j.Type {
	dash := j.Call("", "$.dashboard.new", j.Args(
		j.String("title", d.Title),
		j.String("description", d.Description),
		j.Bool("editable", d.Editable),
		j.Int("graphTooltip", d.GraphTooltip),
		j.String("refresh", d.Refresh),
		j.Int("schemaVersion", d.SchemaVersion),
		j.String("timezone", d.Timezone),
		j.String("uid", d.UID),
		j.String("time_from", d.Time.From),
		j.String("time_to", d.Time.To),
	))

	chain := []j.CallType{dash}
	for _, p := range d.Panels {
		chain = append(chain, p.Jsonnet())
	}

	return j.CallChain(name,
		chain...,
	)
}

type Timepicker struct{}

type Time struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Annotations struct {
}

type Link struct {
	Icon string
	Tags []string
	Type string
}

type Panels []Panel

type Panel struct {
	ID         int     `json:"id"`
	Datasource string  `json:"datasource"`
	GridPos    GridPos `json:"gridPos"`

	Type  string `json:"type"`
	Title string `json:"title"`

	Targets     []map[string]interface{} `json:"targets"`
	Options     map[string]interface{}   `json:"options"`
	Transparent bool                     `json:"transparent"`
	Fill        int                      `json:"fill"`
	Pointradius int                      `json:"pointradius"`
}

func (p Panel) Jsonnet() j.CallType {
	targets := []j.CallType{}

	sort.Slice(p.Targets, func(a, b int) bool {
		var aID, bID string
		if s, ok := p.Targets[a]["refId"]; ok {
			aID = s.(string)
		}

		if s, ok := p.Targets[b]["refId"]; ok {
			bID = s.(string)
		}

		return aID < bID
	})

	for _, t := range p.Targets {
		tgt := j.Call("", "addTarget", j.Args(
			j.Marshal("target", t)),
		)

		targets = append(targets, tgt)
	}

	panelNew := fmt.Sprintf("$.%sPanel.new", p.Type)
	rawPanel := j.Call("panel", panelNew, j.Args(
		j.String("title", p.Title),
		j.String("datasource", p.Datasource),
		j.Int("fill", p.Fill),
		j.Int("pointradius", p.Pointradius),
	))

	panel := j.CallChain("panel",
		append([]j.CallType{rawPanel}, targets...)...,
	)

	return j.Call("", "addPanel", j.Args(
		panel,
		p.GridPos.Jsonnet("gridPos"),
	))
}

type GridPos struct {
	H int `json:"h"`
	W int `json:"w"`
	X int `json:"x"`
	Y int `json:"y"`
}

func (g GridPos) Jsonnet(name string) j.Type {
	return j.Marshal(name, g)
}
