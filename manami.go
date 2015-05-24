package main

import (
	"os"

	"github.com/eliukblau/go-sdl2/sdl"
	"github.com/eliukblau/go-sdl2/sdl_image"
	"github.com/eliukblau/go-sdl2/sdl_mixer"
	_ "github.com/eliukblau/go-sdl2/sdl_ttf"
)

func main() {
	// 0. VARIABLES GLOBALES
	var gameloop bool
	var window *sdl.Window
	var renderer *sdl.Renderer
	var texture *sdl.Texture
	var rectangle *sdl.Rect
	var music *mix.Music

	// 1. INICIAR EL JUEGO

	// 1.1 SDL2
	/*
		if code := sdl.Init(sdl.INIT_EVERYTHING); code == 0 {
			win, err := sdl.CreateWindow("Go-SDL2", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 800, 600, sdl.WINDOW_SHOWN)
			if err != nil {
				panic(err)
			} else {
				window = win
				gameloop = true
			}
		} else {
			os.Exit(code)
		}
	*/

	if code := sdl.Init(sdl.INIT_EVERYTHING); code == 0 {
		win, rend, err := sdl.CreateWindowAndRenderer(1080, 720, sdl.WINDOW_SHOWN|sdl.WINDOW_OPENGL)
		if err != nil {
			panic(err)
		} else {
			window = win
			renderer = rend
			gameloop = true
			window.SetTitle("Manami - Simple Game Skeleton for Go/SDL2")
		}
	} else {
		os.Exit(1)
	}

	// 1.2 SDL2_IMAGE
	if code := img.Init(img.INIT_PNG); code != 0 {
		tex, err := img.LoadTexture(renderer, "gfx/manami_logo.png")
		if err != nil {
			// luego hacer mejor con defer!
			renderer.Destroy()
			window.Destroy()
			sdl.Quit()
			panic(err)
		} else {
			texture = tex
			rectangle = new(sdl.Rect)
			rectangle.X, rectangle.Y = 0, 0
			_, _, w, h, err := texture.Query()
			if err != nil {
				texture.Destroy()
				renderer.Destroy()
				window.Destroy()
				sdl.Quit()
				panic(err)
			} else {
				rectangle.W, rectangle.H = int32(w), int32(h)
			}
		}
	} else {
		// luego hacer mejor con defer!
		renderer.Destroy()
		window.Destroy()
		sdl.Quit()
		os.Exit(2)
	}

	// 1.3 SDL2_MIXER
	if code := mix.OpenAudio(44100, sdl.AUDIO_S16SYS, 2, 4096); code {
		music = mix.LoadMUS("sfx/bg_music.ogg")
		if music == nil || !music.Play(-1) {
			// luego hacer mejor con defer!
			music.Free()
			texture.Destroy()
			renderer.Destroy()
			window.Destroy()
			mix.CloseAudio()
			sdl.Quit()
			os.Exit(2)
		}
	} else {
		// luego hacer mejor con defer!
		texture.Destroy()
		renderer.Destroy()
		window.Destroy()
		mix.CloseAudio()
		sdl.Quit()
		os.Exit(2)
	}

	// 2. BUCLE DEL JUEGO
	for gameloop {
		// 2.1 PROCESA LA ENTRADA
		switch event := sdl.PollEvent(); event.(type) {
		case *sdl.QuitEvent:
			gameloop = false
		}

		// 2.2 ACTUALIZAR EL ESTADO DEL JUEGO

		// 2.3 RENDERIZAR EL JUEGO
		renderer.Clear()
		renderer.Copy(texture, nil, rectangle)
		renderer.Present()
		sdl.Delay(10)
	}

	// 3. FINALIZAR EL JUEGO
	music.Free()
	texture.Destroy()
	renderer.Destroy()
	window.Destroy()
	mix.CloseAudio()
	sdl.Quit()
}
