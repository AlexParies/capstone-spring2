package main
//use this for whatever
import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"image"
	"strconv"

	"gocv.io/x/gocv"
)

type spotVar struct {
	id     int
	status bool
	x      int
	y      int
}

var confirmButton = false
var spots []spotVar
var inputT int
var curFrame image.Image
var spotCount int
var setupFin = false

func main() {

	fmt.Println("yay")

	webcam, _ := gocv.VideoCaptureDevice(0)
	img := gocv.NewMat()

	c1 := make(chan string)
	setup := make(chan bool)

	go func() {
		for { //This is the webcam
			webcam.Read(&img)
			img2, _ := img.ToImage()
			curFrame = img2


			c1 <- " "
		}
	}()

	a := app.New()
	w := a.NewWindow("Park Pallette")
	w.Resize(fyne.NewSize(750, 350))

	input := widget.NewEntry()
	inputCon := container.NewVBox(input)
	inputCon.MinSize()

	input.SetPlaceHolder("Number of spots")

	button := widget.NewButton("confirm", func() {
		inputT, _ = strconv.Atoi(input.Text)
		if inputT > 0 {
			setup <- true
		}
	})

	grid := container.NewGridWithRows(2, inputCon, button)

	content := container.NewGridWithColumns(1, grid)

	w.SetContent(content)

	go func() {
		spotText := "" // the thing to check if setup finishes

		for i := 0; i <= 0; {
			select {
			case setupVal := <-setup: //checks for status of setup
				fmt.Println("Setup:", inputT, setupVal)
				for l := 1; l <= inputT; l++ { //makes spot variables
					var s spotVar

					s.id = l
					spots = append(spots, s)
				}
				for v := 1; v <= inputT; v++ {
					spotText += ("\nID: " + strconv.Itoa(spots[v-1].id) + "\n -Active: " + strconv.FormatBool(spots[v-1].status) + "\n -Coordinates: " + strconv.Itoa(spots[v-1].x) + "," + strconv.Itoa(spots[v-1].y) + "\n")
				}
				fmt.Println(spots)

				text := container.NewCenter(widget.NewLabel(spotText))
				image := canvas.NewImageFromImage(curFrame)
				image.SetMinSize(fyne.NewSize(200, 200))
				image.FillMode = canvas.ImageFillContain

				content = container.NewGridWithColumns(2, image, text)

				w.SetContent(content)

				setupFin = true
				break
			}
		}
	}()

	go func(){ //this is the screen after setup
		       //pretty much the same as the one above
		for{
			if(setupFin){
			spotText := ""

			select {
			case wow := <-c1:

				for v := 1; v <= inputT; v++ {
					spotText += ("\nID:" +wow + strconv.Itoa(spots[v-1].id) + "\n -Active: " + strconv.FormatBool(spots[v-1].status) + "\n -Coordinates: " + strconv.Itoa(spots[v-1].x) + "," + strconv.Itoa(spots[v-1].y) + "\n")
				}

				text := container.NewCenter(widget.NewLabel(spotText))
				image := canvas.NewImageFromImage(curFrame)
				image.SetMinSize(fyne.NewSize(200,200))
				image.FillMode = canvas.ImageFillContain

				content = container.NewGridWithColumns(2,image,text)


				w.SetContent(content)
			}
			}
		}
	}()

	w.ShowAndRun()

}
