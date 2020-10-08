package pipe

import (
	"encoding/json"
	"github.com/pterm/pterm"
)

type Scheme struct {
	Modules []Module `json:"Module"`
}

type Output struct {
	Info    []string `json:"Info"`
	Success []string `json:"Success"`
	Warning []string `json:"Warning"`
	Error   []string `json:"Error"`
	Fatal   []string `json:"Fatal"`
	Other   []string `json:"Other"`
}

type Finished struct {
	Success []string `json:"Success"`
	Failed  []string `json:"Failed"`
}

type Files struct {
	Finished Finished `json:"Finished"`
	Todo     []string `json:"Todo"`
}

type Module struct {
	Name   string `json:"Name"`
	Output Output `json:"Output"`
	Files  Files  `json:"Files"`
}

func (p Scheme) Sprint() string {
	j, _ := json.Marshal(p)
	return string(j)
}

func (p *Scheme) Print() {
	pterm.Print(p.Sprint())
}
