package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	//"net/http/httputil"
	"fmt"
	"image"
	"image/png"
	"os"
	"bytes"
	"io"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) getHelloWorld(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//w.Header().Set("content-type", "text/plain")
	//_, _ = w.Write([]byte("Hello World!"))
	fmt.Println("Cosa contiene il body: ")
	/*fmt.Println(r.Body)
	dump,err := httputil.DumpRequest(r,true)
	if err != nil{
		fmt.Println("errore")
	}
	//fmt.Printf("%q",dump)*/


	in, _ := io.ReadAll(r.Body)
	img, _, _ := image.Decode(bytes.NewReader(in))
	//save dump to a file
	out, err := os.Create("./Prova.png")

	if err != nil{
		fmt.Println("Errore os create")
		return
	}
	err = png.Encode(out,img)

	if err != nil{
		fmt.Println("Errore png encode")
		return
	}

}
