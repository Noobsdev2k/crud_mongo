package controllers.User
import(
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)
type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	if(!bson.IsObjectIdHex(id)){
		w.WriterHeader(http.StatusNotFound)
	}
	oid := bson.ObjectHex(id)
	u := models.User{}
	if err := uc.Session.DB("crud_mongo").C("user").FindId(oid).One(&u); err != nil{
		w.WriterHeader(404)
		return
	}
	uj, err :=json.Marshal(u)
	if err != nil{
		fmt.Printf(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Printf(w, "%s\n", uj)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
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
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id :=p.ByName("id")
	if(!bson.IsObjectIdHex(id)){
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	oid := bson.IsObjectIdHex(id) //
	if err := uc.session.DB("crud_mongo").C("user").RemoveId(id); err != nil{
		w.WriteHeader(404)
	}
	w.WriteHeader(http.StatusOK)	
	fmt.Fprintf(w, "User deleted", oid, "\n")
}
