package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/jung-kurt/gofpdf"
	"github.com/otiai10/gosseract"
)

//Create ...
//Initialize tesseract engine on image when in pdf extracted folder
func (person *User) Create() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Times", "", 16)

	client := gosseract.NewClient()
	lang := []string{"eng"}
	client.Languages = lang
	//pdf.SetLineWidth(0.1)
	i := 0
	for i != person.Length+1 {
		client.SetImage("image" + strconv.Itoa(i) + ".jpg")
		text, err := client.Text()
		if err == nil {
			pdf.AddPage()
			pdf.MultiCell(0, 5, string(text), "", "", false)
		} else {
			fmt.Println(err)
		}
		i++
		fmt.Println(i)
	}
	os.Chdir("..") //This is to change to the next directory and delete
	os.RemoveAll(person.Path)
	//pdf.WritePdf("hello.pdf")
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		println(err, "This has to do with creating")
	}
}

func (person *User) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	templates, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}
	if err = templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		panic(err)
	}

	r.ParseMultipartForm(32 << 20)
	upload, _, err := r.FormFile("file")
	defer upload.Close()
	if err != nil {
		//Solve this
	}
}
