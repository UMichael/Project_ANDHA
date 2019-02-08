package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"./handlers"
)

func main() {
	person := handlers.User{
		Email: "",
		Name:  "",
		File:  "pizza.pdf",
	}
	// person.ReadFile()
	// person.Create()

	person.Path, _ = ioutil.TempDir("", "unk")
	name, err := ioutil.TempFile(person.Path, "file.pdf")
	fmt.Println(name, err)
	//defer os.RemoveAll(name)
	file, _ := os.Open("hello.pdf")
	fmt.Println(name.Name())
	_, err = io.Copy(name, file)
	fmt.Println(err)
	fmt.Println()
}

// fmt.Println("Your password: ")
// bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
// fmt.Println(err)
// password := string(bytePassword)
// fmt.Println() // it's necessary to add a new line after user's input
// fmt.Printf("Your password has leaked, it is '%s'", password)
// os.Chdir("pizza")
// img := gocv.IMRead("image0.jpg", 0)
// if img.Empty() {
// 	println("error")
// 	return
// }
// //var img1 *gocv.Mat
// //gocv.CvtColor(img, img1, gocv.ColorBGRToGray)
// //var img2 gocv.Mat
// gocv.BitwiseNot(img, &img)
// //var img3 gocv.Mat
// gocv.Threshold(img, &img, 0, 255, gocv.ThresholdBinary|gocv.ThresholdOtsu)

// win := gocv.NewWindow("Hello")
// win.IMShow(img)
// win.WaitKey(0)
// println("Done")
//handlers.Deskew()
