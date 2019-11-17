package database

type User struct {
	name string
	password string
	isDoctor bool
}

var databaseUsers = map[string]User{}

func PopulateUsers(){
	databaseUsers["patient1"] = User {
		name: "patient1",
		password: "1234",
		isDoctor: false,
	}
	databaseUsers["patient2"] = User {
		name: "patient2",
		password: "9876",
		isDoctor: false,
	}
	databaseUsers["docteur"] = User {
		name: "docteur",
		password: "abcd",
		isDoctor: true,
	}
}

func CheckUsers(username string, password string) bool {
	if val, ok := databaseUsers[username]; ok {
		if val.password == password {
			return true
		}
	}
	return false
}