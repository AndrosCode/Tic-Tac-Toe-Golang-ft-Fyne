package main

import (
	"image/color"
	"io/ioutil"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Autor: Estuardo Cabrera
// Estudiante de Ingenieria mecanica Electrica
/* EJemplo de un programa con interfaz grafica en Golang
Es mi proyecto mas ambicioso en Go "Soy novato XD"
*/

// Sigue mi Git Hub para mas ejemplos

var a = app.New()
var tamaño = fyne.NewSize(200, 200)
var c = color.RGBA{0x00, 0xff, 0x0f, 0xff}
var textMain = widget.NewLabel("Es el turno de: jugador 1(x)")
var rectu = canvas.NewRectangle(color.NRGBA{R: 0, G: 100, B: 255, A: 200})
var tries int = 1
var values = [3][3]string{
	{"  ", " ", " "},
	{" ", "  ", " "},
	{" ", " ", "   "},
}

func main() {

	var w = a.NewWindow("title")
	w.SetMaster()
	r, _ := LoadResourceFromPath("C:/preXamp/htdocs/golang/fyne/logo.png")
	a.SetIcon(r)
	w.SetIcon(r)
	w.Resize(fyne.NewSize(500, 500))

	w.SetContent(
		container.NewVBox(
			newer(),
			container.NewCenter(
				container.NewGridWithColumns(
					3,
					container.NewGridWrap(tamaño, createBoton(&values[0][0])),
					container.NewGridWrap(tamaño, createBoton(&values[0][1])),
					container.NewGridWrap(tamaño, createBoton(&values[0][2])),
					container.NewGridWrap(tamaño, createBoton(&values[1][0])),
					container.NewGridWrap(tamaño, createBoton(&values[1][1])),
					container.NewGridWrap(tamaño, createBoton(&values[1][2])),
					container.NewGridWrap(tamaño, createBoton(&values[2][0])),
					container.NewGridWrap(tamaño, createBoton(&values[2][1])),
					container.NewGridWrap(tamaño, createBoton(&values[2][2])),
				),
			),
		),
	)

	w.CenterOnScreen()
	w.ShowAndRun()
}

func createBoton(v *string) *fyne.Container {
	text := widget.NewLabel(*v)
	colorA := color.NRGBA{R: 200, G: 200, B: 200, A: 255}
	mainColor := color.NRGBA{R: 200, G: 200, B: 200, A: 255}
	rect := canvas.NewRectangle(colorA)
	btn := widget.NewButton(" ", func() {

		if colorA == mainColor {
			if tries%2 == 0 {
				*v = "O"
				text.SetText(*v)
				colorA = color.NRGBA{R: 255, G: 100, B: 0, A: 255}
				tries++
				updater()
			} else {
				*v = "X"
				text.SetText(*v)
				colorA = color.NRGBA{R: 0, G: 100, B: 255, A: 255}
				tries++
				updater()
			}
		}
		rect.FillColor = colorA
		rect.Refresh()
	})

	contenedor := container.NewMax(
		btn,
		rect,
		container.NewCenter(
			text,
			text,
		),
	)
	btn.Resize(tamaño)
	rect.Resize(tamaño)
	rect.Show()
	return (contenedor)
}
func newer() *fyne.Container {
	contenedora := container.NewMax(
		rectu,
		container.NewCenter(
			textMain,
			textMain,
		),
	)
	return contenedora
}
func updater() {
	if tries%2 == 0 {
		rectu.FillColor = color.NRGBA{R: 255, G: 100, B: 0, A: 255}
		textMain.SetText("Es el turno de: jugador 2(O)")
		rectu.Refresh()
	} else {
		rectu.FillColor = color.NRGBA{R: 0, G: 100, B: 255, A: 255}
		textMain.SetText("Es el turno de: jugador 1(X)")
		rectu.Refresh()
	}
	matcher()
}
func matcher() {

	if values[0][0] == values[0][1] && values[0][1] == values[0][2] {

		if values[0][1] == "X" {
			winnerWindow("X")
		} else if values[0][1] == "O" {
			winnerWindow("O")
		}
	} else if values[1][0] == values[1][1] && values[1][1] == values[1][2] {
		if values[1][1] == "X" {
			winnerWindow("X")
		} else if values[1][1] == "O" {
			winnerWindow("O")
		}
	} else if values[2][0] == values[2][1] && values[2][1] == values[2][2] {
		if values[2][1] == "X" {
			winnerWindow("X")
		} else if values[2][1] == "O" {
			winnerWindow("O")
		}
	} else if values[0][0] == values[1][1] && values[1][1] == values[2][2] {
		if values[1][1] == "X" {
			winnerWindow("X")
		} else if values[1][1] == "O" {
			winnerWindow("O")
		}
	} else if values[2][0] == values[1][1] && values[1][1] == values[0][2] {
		if values[0][2] == "X" {
			winnerWindow("X")
		} else if values[0][2] == "O" {
			winnerWindow("O")
		}

	} else if values[0][0] == values[1][0] && values[1][0] == values[2][0] {
		if values[1][0] == "X" {
			winnerWindow("X")
		} else if values[1][0] == "O" {
			winnerWindow("O")
		}
	} else if values[0][1] == values[1][1] && values[1][1] == values[2][1] {
		if values[0][1] == "X" {
			winnerWindow("X")
		} else if values[0][1] == "O" {
			winnerWindow("O")
		}
	} else if values[0][2] == values[1][2] && values[1][2] == values[2][2] {
		if values[1][2] == "X" {
			winnerWindow("X")
		} else if values[1][2] == "O" {
			winnerWindow("O")
		}
	}

}
func winnerWindow(text string) {
	winner := a.NewWindow("Felicidades!")
	winner.Resize(fyne.NewSize(500, 200))
	if text == "X" {
		background := canvas.NewRectangle(color.NRGBA{R: 0, G: 100, B: 255, A: 255})
		label := widget.NewLabel("Felicidades Jugador 1 Has Ganado")
		winner.SetContent(
			container.NewMax(
				background,
				container.NewCenter(
					label,
					label,
				),
			),
		)

	} else if text == "O" {
		background := canvas.NewRectangle(color.NRGBA{R: 255, G: 100, B: 0, A: 255})
		label := widget.NewLabel("Felicidades Jugador 2 Has Ganado")
		winner.SetContent(
			container.NewMax(
				background,
				container.NewCenter(
					label,
					label,
				),
			),
		)
	}
	r, _ := LoadResourceFromPath("C:/preXamp/htdocs/golang/fyne/logo.png")
	winner.SetIcon(r)
	winner.SetFixedSize(true)
	winner.CenterOnScreen()
	winner.SetMaster()
	winner.Show()
}

type Resource interface {
	Name() string
	Content() []byte
}

/*
El siguiente codigo proviene de StackOverFlow
y permite modificar el icono de la ventana con una imagen
*/

// StaticResource is a bundled resource compiled into the application.
// These resources are normally generated by the fyne_bundle command included in
// the Fyne toolkit.
type StaticResource struct {
	StaticName    string
	StaticContent []byte
}

// Name returns the unique name of this resource, usually matching the file it
// was generated from.
func (r *StaticResource) Name() string {
	return r.StaticName
}

// Content returns the bytes of the bundled resource, no compression is applied
// but any compression on the resource is retained.
func (r *StaticResource) Content() []byte {
	return r.StaticContent
}

// NewStaticResource returns a new static resource object with the specified
// name and content. Creating a new static resource in memory results in
// sharable binary data that may be serialised to the location returned by
// CachePath().
func NewStaticResource(name string, content []byte) *StaticResource {
	return &StaticResource{
		StaticName:    name,
		StaticContent: content,
	}
}

// LoadResourceFromPath creates a new StaticResource in memory using the contents of the specified file.
func LoadResourceFromPath(path string) (Resource, error) {
	bytes, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	name := filepath.Base(path)
	return NewStaticResource(name, bytes), nil
}

/*
Si puedes imaginarlo,
Puedes programarlo

	-Programacion ATS

*/
