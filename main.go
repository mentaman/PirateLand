package main

import (
	"flag"
	"fmt"
	"github.com/mentaman/PirateLand/Game"
	"github.com/vova616/garageEngine/engine"
	//"math"
	//"github.com/go-gl/gl"
	"os"
	//"runtime"
	"runtime/pprof"
	//"time"
)

var cpuprofile = flag.String("p", "", "write cpu profile to file")
var memprofile = flag.String("m", "", "write mem profile to file")

// 
func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Errorf("%s\n", err)
		}
		pprof.StartCPUProfile(f)

		defer pprof.StopCPUProfile()
	}

	file, _ := os.Create("./log.txt")
	os.Stdout = file
	os.Stderr = file
	os.Stdin = file
	defer file.Close()
	go Start()
	engine.Terminated()
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			fmt.Errorf("%s\n", err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}

// 
func Start() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p, engine.PanicPath())
		}

		engine.Terminate()
	}()
	engine.StartEngine()
	_ = Game.MenuSceneG

	/*
		Running local server.
	*/
	//go Server.StartServer()

	engine.LoadScene(Game.MenuSceneG)
	for engine.MainLoop() {

	}
}
