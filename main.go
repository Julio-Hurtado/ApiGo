package main

import (
	"api/db"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("http://localhost:" + db.PortHttp)
	log.Fatal(http.ListenAndServe(":"+db.PortHttp, InitRoutes()))

}
