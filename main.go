package main
import(
	"fmt"
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"
	"github.com/Noobsdev2k/crud_mongo/controllers"
)
func main() {
	router := httprouter.New()
	uc := controllers.NewUserController(getSession())
	router.GET("/:id", uc.GetUser)
	router.POST("/", uc.CreateUser)
	router.DELETE("/:id", uc.DeleteUser)	
	fmt.Println("Starting...")
	http.ListenAndServe("localhost:8080", router)
}
func getSession() *mgo.Session {
	s, err := mgo.Dial("localhost:8080")
	if err != nil {
		panic(err)
	}
}
