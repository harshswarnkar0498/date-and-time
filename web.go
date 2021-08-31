package main
import (
        "fmt"
        "net/http"
        "time"
        "log"
)


func handler(w http.ResponseWriter,r *http.Request) {
               
        t:= time.Now()
	y := t.Year()
	mo := t.Month()
	d := t.Day()
	h := t.Hour()
	m := t.Minute()
	

	fmt.Fprintf(w,"Year : %d\n\n",y)
	fmt.Fprintf(w,"Month : %d\n\n",mo)
	fmt.Fprintf(w,"Date : %d\n\n",d)
	fmt.Fprintf(w,"hour : %d\n\n",h)
	fmt.Fprintf(w,"minute : %d\n\n",m)
	
}
func main() {
        http.HandleFunc("/", handler)
        log.Fatal(http.ListenAndServe(":8000",nil))
}

