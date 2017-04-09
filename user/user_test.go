package user

import (
	"testing"
	//"gopkg.in/mgo.v2"
)

func TestNew(t *testing.T) {
	user := New("john", "password123")
	wanted := "ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f"
	if user.Password != wanted {
		t.Error("Hash is invalid, expected ef...4f, but got:\n", user.Password)
	}
}

func TestAuthentication(t *testing.T) {
	user := New("john", "password123")
	if !user.Authenticate("password123") {
		t.Error("Correct password doesn't authenticate user")
	}

	if user.Authenticate("password") {
		t.Error("Incorrect password authenticates user")
	}
}

/*
func TestDB(t *testing.T) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		t.Error(err)
	}
	defer session.Close()

	c := session.DB("gochat").C("users")
	err = c.Insert(New("john", "password123"))
	if err != nil {
		t.Error(err)
	}
}
*/

func TestFindByUsername(t *testing.T) {
	_, err := FindByUsername("john")
	if err != nil {
		t.Error(err)
	}
}