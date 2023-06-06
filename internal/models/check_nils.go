package models

func CheckUserForNils(user User) User {
	empty := "empty"
	if user.Email == nil {
		user.Email = &empty
	}
	if user.Phone == nil {
		user.Phone = &empty
	}
	return user
}
