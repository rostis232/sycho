package models

const (
	Undefined int = iota
	Admin
	Psychologist
)

// Table users
type User struct {
	UserId *int `db:"user_id"`
	Login *string `db:"login"`
	Pass *string `db:"pass"`
	FirstName *string `db:"first_name"`
	LastName *string `db:"last_name"`
	Email *string `db:"email"`
	Phone *string `db:"phone"`
	OrgId *int `db:"org_id"`
	Role *int `db:"role"`
}

//Table organizations
type Organisation struct{
	OrgId *int `db:"org_id"`
	Title *string `db:"title"`
	Code *string `db:"code"`
}

//Table projects
type Project struct {
	PrjId *int `db:"prj_id"`
	ShortTitle *string `db:"short_title"`
	FullTitle *string `db:"full_title"`
	Code *string `db:"code"`
}

//Table beneficiaries
type Beneficiary struct {
	BnfId *int `db:"bnf_id"`
	FirstName *string `db:"first_name"`
	MiddleName *string `db:"middle_name"`
	LastName *string `db:"last_name"`
	Phone *int `db:"phone"`
	BirthDay *string `db:"birthday"`
	PrjId *int `db:"prj_id"`
	OrgId *int `db:"org_id"`
	UserId *int `db:"user_id"`
	Done *bool `db:"done"`
}

//Table activities Додати опис

type Activity struct {
	ActId *int `db:"act_id"`
	Time *string `db:"time"`
	Req *int `db:"req"`
	Desc *string `db:"description"`
	BnfId *int `db:"bnf_id"`
	UserId *int `db:"user_id"`
}

//Table reqs
type Reqs struct {
	ReqId *int `db:"req_id"`
	Title *string `db:"title"`
}

func GetRoleTitle(role int) string {
	switch role {
	case 0:
		return "Невизначено"
	case 1:
		return "Адміністратор"
	case 2:
		return "Психолог"
	case 3:
		return "Meal-спеціаліст"
	default:
		return "Невизначено"
	}
}