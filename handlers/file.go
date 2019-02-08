package handlers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gocv.io/x/gocv"

	"gopkg.in/gographics/imagick.v2/imagick"
)

type User struct {
	Email  string
	Name   string
	File   string
	Path   string
	Length int
}

func ConvertPDF2Image(file string, path string) int {
	imagick.Initialize()

	mw := imagick.NewMagickWand()
	err := mw.ReadImage(file)
	if err != nil {
		log.Fatal(err)
	}
	mw.SetImageFormat("jpg")
	mw.SetIteratorIndex(0)
	next := true

	err = os.Chdir(path)
	println(err)
	i := 0
	for next == true {
		img := "image" + strconv.Itoa(i) + ".jpg"
		// if err = mw.SetResolution(300, 300); err != nil {
		// 	println(err, 1)
		// }
		//mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_FLATTEN)
		//mw.SetCompressionQuality(95)
		mw.DeskewImage(255)
		fmt.Println(img)
		mw.WriteImage(img)
		jpg := gocv.IMRead(img, 0)
		//gocv.BitwiseNot(jpg, &jpg)
		//gocv.Threshold(jpg, &jpg, 0, 255, gocv.ThresholdBinary|gocv.ThresholdOtsu)
		gocv.IMWrite(img, jpg)
		next = mw.NextImage()
		i++
	}
	fmt.Println("done")
	mw.Destroy()
	imagick.Terminate()
	println("hello")
	return i

}

func (person *User) ReadFile() {

	i := len(person.File)
	for i > 0 && string(person.File[i-1]) != "." {
		// if string(person.File[i-1]) == "." {
		// 	fmt.Println(i, person.File)
		// 	person.File = strings.TrimPrefix(person.File, string(person.File[i-1]))
		// 	break
		// }
		// person.File = strings.TrimLeft(person.File, string(person.File[i-1]))
		// i--
		// fmt.Println(person.File, i)
		person.Path = person.File[:i-1]
		fmt.Println(person.Path)
		i--
	}
	person.Path = person.File[:i-1]
	fmt.Println(person.Path)
	os.Mkdir(person.Path, 0777)
	person.Length = ConvertPDF2Image(person.File, person.Path)
	println("done with files")
}
