package main

import (
	"path/filepath"

	"github.com/eliukblau/go-sdl2/sdl"
	"github.com/eliukblau/go-sdl2/sdl_image"
	"github.com/eliukblau/go-sdl2/sdl_mixer"
	_ "github.com/eliukblau/go-sdl2/sdl_ttf"
)

func main() {
	// 0 - CONSTANTES Y VARIABLES GLOBALES
	const (
		WinWidth  = 1080
		WinHeight = 720
	)

	// 1 - INICIAR EL JUEGO

	// 1.1 - SDL2
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// hint para el modo "fullscreen spaces" de mac
	sdl.SetHint(sdl.HINT_VIDEO_MAC_FULLSCREEN_SPACES, "1")
	// hint para escalar con mejores resultados visuales
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "best")

	// creamos la ventana del juego
	window, err := sdl.CreateWindow(
		"Manami - Simple Game Skeleton for Go/SDL2",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		WinWidth,
		WinHeight,
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
	rectangle.W, rectangle.H = w, h

	// 1.3 - SDL2_MIXER
	if err := mix.OpenAudio(44100, sdl.AUDIO_S16SYS, 2, 4096); err != nil {
		panic(sdl.GetError())
	}
	defer mix.CloseAudio()

	// iniciamos la musica de fondo
	music, err := mix.LoadMUS(filepath.Join("sfx", "bg_music.ogg"))
	if err != nil {
		panic(err)
	}
	if err := music.Play(-1); err != nil {
		panic(err)
	}
	defer music.Free()

	// levantamos la ventana justo antes de comenzar el gameloop
	window.Show()
	window.Raise()

	// 2 - BUCLE PRINCIPAL DEL JUEGO
	gameloop := true
	for gameloop {
		// 2.1 - PROCESA LA ENTRADA
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) { // ojo! los casos deben ser punteros (*)
			case *sdl.QuitEvent:
				gameloop = false
			case *sdl.KeyUpEvent:
				switch t.Keysym.Sym {
				case sdl.K_f:
					flags := window.GetFlags() ^ sdl.WINDOW_FULLSCREEN_DESKTOP
					if err := window.SetFullscreen(flags); err == nil {
						var displayMode sdl.DisplayMode
						sdl.GetDisplayMode(0, 0, &displayMode)

						if flags&sdl.WINDOW_FULLSCREEN_DESKTOP != 0 {
							rectangle.W, rectangle.H = displayMode.W, displayMode.H
						} else {
							_, _, rectangle.W, rectangle.H, _ = texture.Query()
						}
					}
				case sdl.K_ESCAPE:
					gameloop = false
				}
			}
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
