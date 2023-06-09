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
	Role *int `db:"role"`
}

//Table beneficiaries
type Beneficiary struct {
	BnfId *int `db:"bnf_id"`
	FirstName *string `db:"first_name"`
	MiddleName *string `db:"middle_name"`
	LastName *string `db:"last_name"`
	Phone *int `db:"phone"`
	BirthDay *string `db:"birthday"`
	UserId *int `db:"user_id"`
}

//Table activities Додати опис

type Activity struct {
	ActId *int `db:"act_id"`
	Time *string `db:"time"`
	Desc *string `db:"description"`
	BnfId *int `db:"bnf_id"`
	UserId *int `db:"user_id"`
}

func GetRoleTitle(role int) string {
	switch role {
	case 0:
		return "Невизначено"
	case 1:
		return "Адміністратор"
	case 2:
		return "Психолог"
	default:
		return "Невизначено"
	}
}