package main
import(
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"
	"github.com/Noobsdev2k/crud_mongo/controllers"
)
func main() {
	router := httprouter.New()
	uc := controllers.NewUserController(getSession())
	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)	
	fmt.Println("Starting...")
	http.ListenAndServe("localhost:8080", router)
}
func getSession() *mgo.Session {
	s, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return s
}
