package main

import (
	"github.com/raeanne-marks/huego"
	"github.com/veandco/go-sdl2/sdl"
)

func setup(ls *LightShow) {
	i := sdl.Init(sdl.INIT_EVERYTHING)
	if i < 0 {
		panic(sdl.GetError())
	}

	ls.Bridge = new(huego.HueBridge)
	err := ls.Bridge.SetupBridge()
	if err != nil {
		panic(err)
	}

}
