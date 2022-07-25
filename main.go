package main

import (
	con "animalcrosing/constans"
	"animalcrosing/functions"
)

func main() {
	functions.Init()
	for con.Running {
		functions.Input()
		functions.Update()
		functions.Render()
	}

	functions.Quit()
}
