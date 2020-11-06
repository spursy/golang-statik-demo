//go:generate statik -src ./static

package main
import (
	"fmt"
	_ "github.com/golang-statik-demo/statik"
	"github.com/rakyll/statik/fs"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	//r, err := statikFS.Open("/a.txt")
	r, err := statikFS.Open("/sub/b.txt")
	defer func() {
		_ = r.Close()
	}()
	if err != nil {
		log.Fatal(err)
	}
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(contents))

	// Add http server
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	_ = http.ListenAndServe(":8080", nil)
}