package main

import (
	"path/filepath"

	"github.com/eliukblau/go-sdl2/sdl"
	"github.com/eliukblau/go-sdl2/sdl_image"
	"github.com/eliukblau/go-sdl2/sdl_mixer"
	_ "github.com/eliukblau/go-sdl2/sdl_ttf"
)

func main() {
	// 1 - INICIAR EL JUEGO

	// 1.1 - SDL2
	if code := sdl.Init(sdl.INIT_EVERYTHING); code != 0 {
		panic(sdl.GetError())
	}
	defer sdl.Quit()

	// creamos la ventana del juego
	window, err := sdl.CreateWindow(
		"Manami - Simple Game Skeleton for Go/SDL2",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		1080,
		720,
		sdl.WINDOWEVENT_HIDDEN|sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// creamos el renderer para la ventana
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	// 1.2 - SDL2_IMAGE
	if code := img.Init(img.INIT_PNG); code == 0 { // ojo! cero si falla
		panic(img.GetError())
	}

	// obtenemos la textura de la imagen de fondo
	texture, err := img.LoadTexture(renderer, filepath.Join("gfx", "manami_logo.png"))
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	// obtenemos algunos datos de la textura
	_, _, w, h, err := texture.Query()
	if err != nil {
		panic(err)
	}

	// usamos los datos de la textura para crear un rectangulo
	rectangle := new(sdl.Rect)
	rectangle.X, rectangle.Y = 0, 0
	rectangle.W, rectangle.H = int32(w), int32(h)

	// 1.3 - SDL2_MIXER
	if ok := mix.OpenAudio(44100, sdl.AUDIO_S16SYS, 2, 4096); !ok {
		panic(sdl.GetError())
	}
	defer mix.CloseAudio()

	// iniciamos la musica de fondo
	music := mix.LoadMUS(filepath.Join("sfx", "bg_music.ogg"))
	if music == nil || !music.Play(-1) {
		panic(sdl.GetError())
	}
	defer music.Free()

	// levantamos la ventana justo antes de comenzar el gameloop
	window.Show()
	window.Raise()

	// 2 - BUCLE PRINCIPAL DEL JUEGO
	gameloop := true
	for gameloop {
		// 2.1 - PROCESA LA ENTRADA
		var event sdl.Event
		for event == nil {
			event = sdl.PollEvent()
		}

		switch event.(type) {
		case *sdl.QuitEvent: // ojo! debe ser un puntero (*)
			gameloop = false
		}

		// 2.2 - ACTUALIZAR EL ESTADO DEL JUEGO

		// 2.3 - RENDERIZAR EL JUEGO
		renderer.Clear()
		renderer.Copy(texture, nil, rectangle)
		renderer.Present()
		sdl.Delay(10)
	}

	// 3 - FINALIZAR EL JUEGO
	// defer's de go ya se encargan de liberar y finalizar todo! ;)
}
