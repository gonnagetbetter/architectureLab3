package main

import (
	"net/http"

	"github.com/gonnagetbetter/architecture-lab-3/painter"
	"github.com/gonnagetbetter/architecture-lab-3/painter/lang"
	"github.com/gonnagetbetter/architecture-lab-3/ui"
)

func main() {
	var (
		pv ui.Visualizer // Візуалізатор створює вікно та малює у ньому.
		opLoop painter.Loop // Цикл обробки команд.

		state = lang.NewCanvasState()
		parser = lang.NewParserWithState(state)
	)

	//pv.Debug = true
	pv.Title = "Simple painter"

	pv.OnScreenReady = opLoop.Start
	opLoop.Receiver = &pv

	go func() {
		http.Handle("/", lang.HttpHandler(&opLoop, parser))
		_ = http.ListenAndServe("localhost:17000", nil)
	}()

	pv.Main()
	opLoop.StopAndWait()
}
