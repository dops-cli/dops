package pipe

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
	return GetByIndex(len(PipeModules.Modules))
}

func GetByIndex(i int) Module {
	return PipeModules.Modules[i]
}
