package routes

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Saiga")
}
