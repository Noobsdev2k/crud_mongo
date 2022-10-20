package controllers
import(
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"github.com/Noobsdev2k/crud_mongo/models"
)
type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	u := models.User{}
	if err := uc.session.DB("crud_mongo").C("user").FindId(oid).One(&u); err != nil{
		w.WriteHeader(404)
		return
	}
	uj, err :=json.Marshal(u)
	if err != nil{
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	uc.session.DB("crud_mongo").C("user").Insert(u)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(w, "%s\n", uj)
}

// func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){

// }
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id :=p.ByName("id")
	if(!bson.IsObjectIdHex(id)){
		w.WriteHeader(404)
		return
	}
	oid := bson.IsObjectIdHex(id) //
	if err := uc.session.DB("crud_mongo").C("user").RemoveId(id); err != nil{
		w.WriteHeader(404)
	}
	w.WriteHeader(http.StatusOK)	
	fmt.Fprintf(w, "User deleted", oid, "\n")
}
