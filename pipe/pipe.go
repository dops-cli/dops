package pipe

import "encoding/json"

var PipeModules Scheme

func Print() {
	if !IsPiped() {
		return
	}

	PipeModules.Print()
}

func AddModule(mod Module) {
	PipeModules.Modules = append(PipeModules.Modules, mod)
}

func GetLast() Module {
	return GetByIndex(len(PipeModules.Modules) - 1)
}

func GetByIndex(i int) Module {
	return PipeModules.Modules[i]
}

func GetSchemeFromJSON(jsonString string) *Scheme {
	newMod := &Scheme{}

	err := json.Unmarshal([]byte(jsonString), newMod)
	if err != nil {
		panic(err)
	}

	return newMod
}
