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

func CheckOrgForNils(org Organisation) Organisation {
	empty := "empty"
	if org.Code == nil {
		org.Code = &empty
	}
	return org
}

func CheckProjectForNil(prj Project) Project {
	empty := "empty"
	if prj.FullTitle == nil {
		prj.FullTitle = &empty
	}
	if prj.Code == nil {
		prj.Code = &empty
	}
	return prj
}