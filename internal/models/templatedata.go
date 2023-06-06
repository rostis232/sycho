package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	Title     string
	Menu      []MenuItem
	Active    string
	StringMap map[string]string
	// IntMap    map[string]int
	// FloatMap  map[string]float32
	Data      map[string]interface{}
	// CSRFToken string
	// Flash     string
	// Warning   string
	// Error     string
	// Form      *forms.Form
}

type MenuItem struct {
	Title string
	URI   string
}

var (
	Main = MenuItem{
		Title: "Головна",
		URI:   "/",
	}
	Profile = MenuItem{
		Title: "Профіль",
		URI:   "/profile",
	}
	Clients = MenuItem{
		Title: "Клієнти",
		URI:   "/client",
	}
	Journal = MenuItem{
		Title: "Журнал",
		URI:   "/journal",
	}
	Help = MenuItem{
		Title: "Підтримка",
		URI:   "/help",
	}
	Login = MenuItem{
		Title: "Увійти/Зареєструватись",
		URI:   "/login",
	}
	Logout = MenuItem{
		Title: "Вийти",
		URI:   "/logout",
	}
)

var UnLoggedUserMenu = []MenuItem{Main, Help, Login}

var PsychologistMenu = []MenuItem{Main, Clients, Journal, Help, Profile, Logout}

var AdminMenu = []MenuItem{Main, Help, Profile, Logout}

func ReturnMenuByRole(role int) []MenuItem {
	switch role {
	case Undefined:
		return UnLoggedUserMenu
	case Admin:
		return AdminMenu
	case Psychologist:
		return PsychologistMenu
	default:
		return UnLoggedUserMenu
	}
}
