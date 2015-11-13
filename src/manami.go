package main

import (
	"runtime"
	"time"

	sf "github.com/manyminds/gosfml"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	// 0 - CONSTANTES Y VARIABLES GLOBALES

	const (
		winWidth  = 1080
		winHeight = 720
	)

	// para sincronizar el framerate a 60fps
	ticker := time.NewTicker(time.Second / 60)

	// 1 - INICIAR EL JUEGO

	// creamos la ventana del juego
	window := sf.NewRenderWindow(
		sf.VideoMode{Width: winWidth, Height: winHeight, BitsPerPixel: 32},
		"Manami - Simple Game Skeleton for Go/SFML2",
		sf.StyleTitlebar|sf.StyleClose,
		sf.DefaultContextSettings())

	// cerrado de ventana al final de la ejecucion
	defer window.Close()

	// activamos la sincronizacion vertical
	window.SetVSyncEnabled(true)

	// centramos la ventana
	window.SetPosition(sf.Vector2i{
		X: int((sf.GetDesktopVideoMode().Width - window.GetSize().X) / 2),
		Y: int((sf.GetDesktopVideoMode().Height - window.GetSize().Y) / 2),
	})

	// cargamos la textura de la imagen de fondo
	texture, err := sf.NewTextureFromFile(ResourcePath("gfx", "manami_logo.png"), nil)
	if err != nil {
		panic(err)
	}
	// creamos un sprite a partir de la textura
	sprite, err := sf.NewSprite(texture)
	if err != nil {
		panic(err)
	}

	// cargamos la musica de fondo
	music, err := sf.NewMusicFromFile(ResourcePath("sfx", "bg_music.ogg"))
	if err != nil {
		panic(err)
	}
	// reproducimos la musica de fondo
	music.SetLoop(true)
	music.Play()

	// 2 - BUCLE PRINCIPAL DEL JUEGO

	gameloop := true
	for gameloop && window.IsOpen() {
		select {
		case <-ticker.C: // cada 60 fps
			// 2.1 - PROCESA LA ENTRADA
			for event := window.PollEvent(); event != nil; event = window.PollEvent() {
				switch ev := event.(type) {
				case sf.EventClosed:
					gameloop = false

				case sf.EventKeyReleased:
					switch ev.Code {
					case sf.KeyEscape:
						gameloop = false
					}
				}
			}

			// 2.2 - ACTUALIZAR EL ESTADO DEL JUEGO

			// 2.3 - RENDERIZAR EL JUEGO
			window.Clear(sf.ColorBlack())
			window.Draw(sprite, sf.DefaultRenderStates())
			window.Display()
		}
	}

	// 3 - FINALIZAR EL JUEGO
	// defer's de go ya se encargan de liberar y finalizar todo! ;)

}
