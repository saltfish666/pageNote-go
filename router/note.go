package router

import (
	"net/http"
	"fmt"
	"reflect"
	"strings"
	"gopkg.in/mgo.v2/bson"
	"../mgo"
	"encoding/json"
)

type Note struct {
	Id string
	Domain string
	Path string
	Content string
}
func NoteHandler(w http.ResponseWriter, r *http.Request){
	var authorization []string = r.Header["Authorization"]
	fmt.Println(authorization)                                 // [bearer 7bdde0af0f6adf2ada7f0433089866b6ed190e25]
	fmt.Println(reflect.TypeOf(authorization))                 //[]string

	var token string = strings.Fields(authorization[0])[1]     // 7bdde0af0f6adf2ada7f0433089866b6ed190e25
	fmt.Println(token)

	users := mgo.MySession.DB("PageNote").C("users")
	var user User
	users.Find(bson.M{"token": token}).One(&user)
	if user.Name == "" {
		fmt.Fprintf(w, token + " does not macth user")
	}

	notes := mgo.MySession.DB("PageNote").C("users")
	var thisNotes []Note
	notes.Find(bson.M{"id": user.Id}).All(&thisNotes)

	jsonNotes, _ := json.Marshal(thisNotes)
	fmt.Fprintf(w, string(jsonNotes))
}