package main

import (
	"fmt"
	"github.com/raeanne-marks/huego"
	"github.com/veandco/go-sdl2/sdl"
	//sdlm "github.com/veandco/go-sdl2/sdl_mixer"
	//"time"
)

/*
func myfunc(ch chan bool) {
	for {

	}
}
*/

type LightShow struct {
	Bridge *huego.HueBridge
	Lights []huego.HueLight
}

func main() {
	fmt.Println("hello world")

	var ls LightShow
	var event sdl.Event
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error

	window, err = sdl.CreateWindow("Hue Light Show", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, 600, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	running := true

	setup(&ls)

	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			fmt.Println("Got event")
		}
	}

	/*
		sdlm.SetSoundFonts("")

		if sdlm.OpenAudio(sdlm.DEFAULT_FREQUENCY, sdlm.DEFAULT_FORMAT,
			sdlm.DEFAULT_CHANNELS, 4096) == false {
			fmt.Println("Error opening audio")
		}
		fmt.Println("open audio")
		fmt.Println(sdl.GetError())

		song := sdlm.LoadMUS("music/PusherLoveGirl.flac")
		if song == nil {
			fmt.Println("Couldn't load song :(")
			fmt.Println(sdl.GetError())
			return
		} else {
			fmt.Println("Loaded song yay")
		}

		_ = sdlm.SetMusicVolume(100)
		fmt.Println("volume")
		fmt.Println(sdl.GetError())
		t := song.Play(1)
		fmt.Println("play")
		fmt.Println(sdl.GetError())
		time.Sleep(time.Second * 10)

		fmt.Println(t)

		bridge := new(huego.HueBridge)
		err := bridge.SetupBridge()
		if err != nil {
			panic(err)
		}

		lights, err := bridge.GetLights()
		if err != nil {
			panic(err)
		}

		for _, l := range lights {
			l.Reset()
		}

		//	ch := make(chan bool)
		//	go myfunc(ch)
	*/
}
