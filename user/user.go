// User defines the data structure and schema for chat users
package user

import (
	"errors"
	"crypto/sha256"
	"encoding/hex"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DBName string ="gochat"
	CollectonName string = "users"
)

var users *mgo.Collection

type User struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string
	Password string
}

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	users = session.DB("gochat").C("users")
}

func New(username, password string) *User {
	hash := Hash(password)
	return &User{
		Id: bson.NewObjectId(),
		Username: username,
		Password: hash,
	}
}

func (user *User) Save() error {
	return users.Insert(user)
}

func (user User) Authenticate(password string) bool {
	hash := Hash(password)
	if hash != user.Password {
		return false
	}
	return true
}

func Hash(password string) string {
	//Create hash and write the password to it
	sum := sha256.New()
	_, err := sum.Write([]byte(password))
	if err != nil {
		panic(err)
	}

	// Hex encode the checksum
	hash := hex.EncodeToString(sum.Sum(nil))

	return hash
}

func FindByUsername(username string) (*User, error) {
	result := &User{}
	query := bson.M{"username": username}
	err := users.Find(query).One(result)
	return result, err
}

func FindById(id string) (*User, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("invalid ObjectId")
	}
	result := &User{}
	err := users.FindId(bson.ObjectIdHex(id)).One(result)
	return result, err
}