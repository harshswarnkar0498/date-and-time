package main
import (
        "fmt"
        "net/http"
        "time"
        "log"
)


func handler(w http.ResponseWriter,r *http.Request) {
                st := time.Now()
                y "= y.Year()
                fmt.Fprintf(w,"Current time is: %d\n\n",st)
}
func main() {
        http.HandleFunc("/",handler)
        log.fatal(http.ListenAndServe(":8000",nil))
}

