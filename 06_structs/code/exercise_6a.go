package main

import "fmt"

// User is a person
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
	return desc
}

// Group is a group of users
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func describeGroup(g Group) string {
	desc := fmt.Sprintf("The %s user group has %d users. Newest user:  %s, Accepting New Users:  %t", g.role, len(g.users), g.newestUser.FirstName, g.spaceAvailable)
	return desc
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

	g := Group{
		role:       "admin",
		users:      []User{u1, u2},
		newestUser: u2,
		spaceAvailable: func() bool {
			if len(Group.users) > 2 {
				return false
			} else {
				return true
			}
		}(),
	}

	userDescription := describeUser(u1)
	groupDescription := describeGroup(g)

	fmt.Println(userDescription)
	fmt.Println(groupDescription)

}
