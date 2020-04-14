package logger

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func Logger(prefix string) func(next httprouter.Handle,) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
			log.Printf(
				"%s Method: %s, path: %s",
				prefix,
				request.Method,
				request.URL.Path,
			)
			data := fmt.Sprintf("%s: [%s] %s Method: %s, path: %s\n",
				time.Now().Format("2006-01-02 15:04:05"),
				strings.Split(request.RemoteAddr, ":")[0],
				prefix,
				request.Method,
				request.URL.Path,
				)
			//err := ioutil.WriteFile("logfile.log", []byte(data), 0777)
			file, err := os.OpenFile("logfile.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
			if err != nil {
				log.Panic("Failed to log to file", err)
				panic(err)
			}
			defer func() {
				err := file.Close()
				if err != nil {
					log.Println("Can't close file")
				}
			}()
			_, err = file.Write([]byte(data))
			if err != nil {
				log.Fatalf("Xuyovo vsyo")
			}
			//io.Copy(os.Stdout, []byte(data))
			next(writer, request, pr)
		}
	}
}
