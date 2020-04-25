package main

import "fmt"

// User is a person
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// Group is a group of users
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

// Describer impliments things that describe
type Describer interface {
	describe() string
}

func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
	return desc
}

func (g *Group) describe() string {
	func() bool {
		if len(g.users) > 2 {
			g.spaceAvailable = false
		} else {
			g.spaceAvailable = true
		}
		return g.spaceAvailable
	}()
	desc := fmt.Sprintf("The %s user group has %d users. Newest user:  %s, Accepting New Users:  %t", g.role, len(g.users), g.newestUser.FirstName, g.spaceAvailable)
	return desc
}

// Describeing describes a struct thats passed in
func Describing(d Describer) string {
	return d.describe()
}

func main() {

	u1 := User{
		ID:        1,
		FirstName: "Marilyn",
		LastName:  "Monroe",
		Email:     "marilyn.monroe@gmail.com",
	}
	u2 := User{
		ID:        2,
		FirstName: "Humphrey",
		LastName:  "Bogart",
		Email:     "humphrey.bogart@gmail.com",
	}

	u3 := User{
		ID:        3,
		FirstName: "Humphrey",
		LastName:  "Bogart",
		Email:     "humphrey.bogart@gmail.com",
	}

	g := Group{
		role:           "admin",
		users:          []User{u1, u2, u3},
		newestUser:     u2,
		spaceAvailable: false,
	}

	gThing := Describing(&g)

	fmt.Println(gThing)

}
