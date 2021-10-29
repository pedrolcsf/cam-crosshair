package main

import (
	"image"
	"image/color"
	"log"

	"github.com/go-vgo/robotgo"

	"gocv.io/x/gocv"
)

type ImageCV struct {
	mat gocv.Mat
}

var (
	blue = color.RGBA{0, 0, 255, 0}
	red  = color.RGBA{240, 52, 52, 1}

	xx = 0
	yy = 0

	xx1 = 0
	yy1 = 0
)

func (icv *ImageCV) Resizee(width, height int) *ImageCV {
	resizeMat := gocv.NewMat()
	gocv.Resize(icv.mat, &resizeMat, image.Pt(width, height), 0, 0, gocv.InterpolationArea)
	_ = icv.mat.Close()
	icv.mat = resizeMat
	return icv
}

func main() {
	// open webcam. 0 is the default device ID, change it if your device ID is different
	webcam, err := gocv.VideoCaptureDevice(0)

	if err != nil {
		log.Fatalf("error opening web cam: %v", err)
	}

	defer webcam.Close()

	// prepare image matrix
	img := gocv.NewMat()

	// gocv.Rotate(img, &img, gocv.RotateFlag(80))

	// open display window
	window := gocv.NewWindow("CAM CROSSHAIR") // Cria uma nova janela e da o nome dela
	defer window.Close()

	go addKeysListen("mleft", img)

	for {
		// se nao achar a webcam, da erro
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Print("cannot read webcam")
			continue
		}

		if xx > 0 {
			gocv.Rectangle(&img, image.Rectangle{
				Min: image.Pt(xx-1300, yy),
				Max: image.Pt(xx-1301, yy-30),
			}, blue, 1)

			gocv.Rectangle(&img, image.Rectangle{
				Min: image.Pt(xx-1280, yy-16),
				Max: image.Pt(xx-1321, yy-15),
			}, blue, 1)

		}
		if xx > 0 && xx1 > 0 {
			//
			gocv.Rectangle(&img, image.Rectangle{
				Min: image.Pt(xx-1300, yy),
				Max: image.Pt(xx-1301, yy-30),
			}, blue, 1)

			gocv.Rectangle(&img, image.Rectangle{
				Min: image.Pt(xx-1280, yy-16),
				Max: image.Pt(xx-1321, yy-15),
			}, blue, 1)
			//

			//
			gocv.Rectangle(&img, image.Rectangle{
				Min: image.Pt(xx1-1300, yy1),
				Max: image.Pt(xx1-1301, yy1-30),
			}, blue, 1)

			gocv.Rectangle(&img, image.Rectangle{
				Min: image.Pt(xx1-1280, yy1-16),
				Max: image.Pt(xx1-1321, yy1-15),
			}, blue, 1)
			//
		}

		// // make a "x"

		window.IMShow(img)

		// fps
		window.WaitKey(1)

	}
}

func addKeysListen(key string, img gocv.Mat, arr ...string) {
	for {
		if ok := robotgo.AddEvent("mleft"); ok {

			x, y := robotgo.GetMousePos()

			if xx > 0 && xx1 > 0 {
				xx = 0
				xx1 = 0
				yy = 0
				yy1 = 0
			}

			if xx > 0 {
				xx1 = x
				yy1 = y
			} else {
				xx = x
				yy = y
			}

			// fmt.Println("pos: ", xx, yy)
		}
	}
}

func (icv *ImageCV) Resize(width, height int) *ImageCV {
	resizeMat := gocv.NewMat()
	gocv.Resize(icv.mat, &resizeMat, image.Pt(width, height), 0, 0, gocv.InterpolationArea)
	_ = icv.mat.Close()
	icv.mat = resizeMat
	return icv
}
