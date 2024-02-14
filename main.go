package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"


	"gocv.io/x/gocv"
	"image"



	"strconv"
)

type spotVar struct{
	id int
	status bool 
	x int
	y int
}
 var confirmButton = false;
var spots []spotVar;

var curFrame image.Image;

func webCamR(img gocv.Mat, webcam *gocv.VideoCapture){
	
	webcam.Read(&img)
	
	

	img2,_ := img.ToImage()
			
	curFrame = img2

}

func appRunner(){
	
	

	a := app.New()
	w := a.NewWindow("Park Pallette")
	w.Resize(fyne.NewSize(750, 350))

	text := container.NewCenter( widget.NewLabel("--Park Pallette--"))
	

	input := widget.NewEntry()
	inputCon := container.NewVBox(input)
	inputCon.MinSize()


	input.SetPlaceHolder("Number of spots")
	
	

	

	webcam, _ := gocv.VideoCaptureDevice(0)
	img := gocv.NewMat()
	/*
	webCamR(img,webcam)

	image :=  canvas.NewImageFromImage(curFrame)
	image.SetMinSize(fyne.NewSize(100,100))
	image.FillMode = canvas.ImageFillContain*/

	button := widget.NewButton("confirm", func() {
		confirmButton = true
	})


	

	grid := container.NewGridWithRows(2,inputCon,button)








	content := container.NewGridWithColumns(1,grid)


	w.SetContent(content)

	inputT := 0
	
	
	
	setup := false
	
	
	go func(){
		
		
		for{
			spotText := ""

			// this chunk o' code checks if the number of spots has been filled out
			if(confirmButton){
				inputT ,_ = strconv.Atoi(input.Text)
				}else {inputT = 0}
			if(inputT !=0){
				if (!setup){
					
					

					for i := 1; i <= inputT; i++{
						var s spotVar
						
						s.id = i
						spots = append(spots,s)
					}


				}
			for i := 1; i <= inputT; i++{
				spotText += ("\nID: " + strconv.Itoa(spots[i-1].id) + "\n -Active: " + strconv.FormatBool(spots[i-1].status) + "\n -Coordinates: " + strconv.Itoa(spots[i-1].x) + ","  + strconv.Itoa(spots[i-1].y)+"\n")
			}
			
			text = container.NewCenter( widget.NewLabel(spotText))

			
				






			
			webCamR(img,webcam)
			image := canvas.NewImageFromImage(curFrame)
			image.SetMinSize(fyne.NewSize(200,200))
			image.FillMode = canvas.ImageFillContain
		
			content = container.NewGridWithColumns(2,image,text)
		
			
			w.SetContent(content)
		}}
	}()
	

	w.ShowAndRun()
}

func main(){
	fmt.Println("Runin")
	appRunner()
	fmt.Println("Ran")
}