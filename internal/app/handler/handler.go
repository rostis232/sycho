package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rostis232/psycho/internal/app/service"
	"github.com/rostis232/psycho/internal/app/sess"
	"github.com/rostis232/psycho/internal/models"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Home is a handler for main page
func (h *Handler) Home(c echo.Context) error {
	uuid, err := sess.GetUUIDFromSession(c)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	user, err := h.service.Authorization.GetUserByUUID(uuid)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	return c.Render(http.StatusOK, "homepage", models.TemplateData{
		Title:  "Caritas PSS",
		Menu:   models.ReturnMenuByRole(*user.Role),
		Active: models.Main.URI,
	},
	)
}

func (h *Handler) Organisations(c echo.Context) error {
	uuid, err := sess.GetUUIDFromSession(c)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	user, err := h.service.Authorization.GetUserByUUID(uuid)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	if *user.Role != models.Admin {
		return c.Render(http.StatusOK, "error", models.TemplateData{
			Title:  "Помилка - Caritas PSS",
			Menu:   models.ReturnMenuByRole(*user.Role),
			Active: models.Main.URI,
		},
		)
	}
	orgs, err := h.service.GetAllOrganisations()
	if err != nil {
		return c.Render(http.StatusOK, "error", models.TemplateData{
			Title:  "Помилка - Caritas PSS",
			Menu:   models.ReturnMenuByRole(*user.Role),
			Active: models.Main.URI,
		},
		)
	}
	data := make(map[string]interface{})
	data["orgs"] = orgs
	return c.Render(http.StatusOK, "organisations", models.TemplateData{
		Title:  "Організації - Caritas PSS",
		Menu:   models.ReturnMenuByRole(*user.Role),
		Active: models.Organisations.URI,
		Data:   data,
	},
	)
}

func (h *Handler) OrganisationsPost(c echo.Context) error {
	uuid, err := sess.GetUUIDFromSession(c)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	user, err := h.service.Authorization.GetUserByUUID(uuid)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	if *user.Role != models.Admin {
		return c.Render(http.StatusOK, "error", models.TemplateData{
			Title:  "Помилка - Caritas PSS",
			Menu:   models.ReturnMenuByRole(*user.Role),
			Active: models.Main.URI,
		},
		)
	}
	_, err = h.service.AddOrganisation(c.FormValue("title"), c.FormValue("code"))
	if err != nil {
		return err
	}

	orgs, err := h.service.GetAllOrganisations()
	if err != nil {
		return c.Render(http.StatusOK, "error", models.TemplateData{
			Title:  "Помилка - Caritas PSS",
			Menu:   models.ReturnMenuByRole(*user.Role),
			Active: models.Main.URI,
		},
		)
	}
	data := make(map[string]interface{})
	data["orgs"] = orgs
	return c.Render(http.StatusOK, "organisations", models.TemplateData{
		Title:  "Організації - Caritas PSS",
		Menu:   models.ReturnMenuByRole(*user.Role),
		Active: models.Organisations.URI,
		StringMap: map[string]string{
			"success": "Організацію успішно додано",
		},
		Data: data,
	},
	)
}

func (h *Handler) Projects(c echo.Context) error {
	uuid, err := sess.GetUUIDFromSession(c)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	user, err := h.service.Authorization.GetUserByUUID(uuid)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	if *user.Role != models.Admin {
		return c.Render(http.StatusOK, "error", models.TemplateData{
			Title:  "Помилка Caritas PSS",
			Menu:   models.ReturnMenuByRole(*user.Role),
			Active: models.Main.URI,
		},
		)
	}
	prjs, err := h.service.GetAllProjects()
	if err != nil {
		return c.Render(http.StatusOK, "error", models.TemplateData{
			Title:  "Помилка Caritas PSS",
			Menu:   models.ReturnMenuByRole(*user.Role),
			Active: models.Main.URI,
		},
		)
	}
	data := make(map[string]interface{})
	data["prjs"] = prjs
	return c.Render(http.StatusOK, "projects", models.TemplateData{
		Title:  "Caritas PSS",
		Menu:   models.ReturnMenuByRole(*user.Role),
		Active: models.Projects.URI,
		Data:   data,
	},
	)
}

func (h *Handler) LogInGet(c echo.Context) error {
	return c.Render(http.StatusOK, "login", models.TemplateData{
		Title:  "Авторизація - Caritas PSS",
		Menu:   models.UnLoggedUserMenu,
		Active: models.Login.URI,
	},
	)
}

func (h *Handler) AddOrganisation(c echo.Context) error {
	uuid, err := sess.GetUUIDFromSession(c)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	user, err := h.service.Authorization.GetUserByUUID(uuid)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	if *user.Role != models.Admin {
		return c.Render(http.StatusOK, "error", models.TemplateData{
			Title:  "Помилка - Caritas PSS",
			Menu:   models.ReturnMenuByRole(*user.Role),
			Active: models.Main.URI,
		},
		)
	}

	return c.Render(http.StatusOK, "add_org", models.TemplateData{
		Title:  "Додавання організації - Caritas PSS",
		Menu:   models.ReturnMenuByRole(*user.Role),
		Active: models.Organisations.URI,
	},
	)
}

func (h *Handler) LogInPost(c echo.Context) error {
	uuid, err := h.service.Authorization.CreateAndSaveUUID(c.FormValue("login"), c.FormValue("pass"))
	fmt.Println(uuid, err)
	if err == sql.ErrNoRows {
		return c.Render(http.StatusOK, "login", models.TemplateData{
			Title:  "Авторизація - Caritas PSS",
			Menu:   models.UnLoggedUserMenu,
			Active: models.Login.URI,
			StringMap: map[string]string{
				"login": c.FormValue("login"),
				"error": "Помилка авторизації. Перевірте правильність введених даних",
			},
		},
		)
	}
	if err != sql.ErrNoRows && err != nil {
		return err
	}
	err = sess.SaveUUIDToSession(c, uuid)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "loginok", models.TemplateData{
		Title:  "Авторизація - Caritas PSS",
		Menu:   models.UnLoggedUserMenu,
		Active: models.Login.URI,
	},
	)
}

func (h *Handler) Help(c echo.Context) error {
	uuid, err := sess.GetUUIDFromSession(c)
	if err != nil {
		fmt.Println(err)
		return c.Render(http.StatusOK, "help", models.TemplateData{
			Title:  "Допомога - Caritas PSS",
			Menu:   models.ReturnMenuByRole(models.Undefined),
			Active: models.Help.URI,
		},
		)
	}
	user, err := h.service.Authorization.GetUserByUUID(uuid)
	if err != nil {
		fmt.Println(err)
	}
	return c.Render(http.StatusOK, "help", models.TemplateData{
		Title:  "Допомога - Caritas PSS",
		Menu:   models.ReturnMenuByRole(*user.Role),
		Active: models.Help.URI,
	},
	)
}

func (h *Handler) Instructions(c echo.Context) error {
	uuid, err := sess.GetUUIDFromSession(c)
	if err != nil {
		fmt.Println(err)
		return c.Render(http.StatusOK, "instrustions", models.TemplateData{
			Title:  "Інструкції - Caritas PSS",
			Menu:   models.ReturnMenuByRole(models.Undefined),
			Active: models.Instructions.URI,
		},
		)
	}
	user, err := h.service.Authorization.GetUserByUUID(uuid)
	if err != nil {
		fmt.Println(err)
	}
	return c.Render(http.StatusOK, "instrustions", models.TemplateData{
		Title:  "Інструкції - Caritas PSS",
		Menu:   models.ReturnMenuByRole(*user.Role),
		Active: models.Instructions.URI,
	},
	)
}

func (h *Handler) Clients(c echo.Context) error {
	uuid, err := sess.GetUUIDFromSession(c)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	user, err := h.service.Authorization.GetUserByUUID(uuid)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	beneficiaries, err := h.service.GetAllBeneficiariesByUserID(*user.UserId)
	if err != nil {
		return err
	}
	data := make(map[string]interface{})
	data["bens"] = beneficiaries
	return c.Render(http.StatusOK, "clients", models.TemplateData{
		Title:  "Caritas PSS",
		Menu:   models.ReturnMenuByRole(*user.Role),
		Active: models.Clients.URI,
		Data:   data,
	},
	)
}

func (h *Handler) BeneficiaryPage(c echo.Context) error {
	uuid, err := sess.GetUUIDFromSession(c)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	user, err := h.service.Authorization.GetUserByUUID(uuid)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	switch *user.Role {
	case models.Psychologist:
		
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("err")
			return c.Render(http.StatusOK, "error", models.TemplateData{
				Title:  "Помилка - Caritas PSS",
				Menu:   models.ReturnMenuByRole(*user.Role),
				Active: models.Main.URI,
			},
			)
		}
		//Дістаємо бенефіціара з ід у якого співпадає userID та orgid
		beneficiary, err := h.service.GetBeneficiaryByID(id)
		if err != nil {
			return err
		}
		acts, err := h.service.GetActivitiesByBnfID(id)
		if err != nil {
			return err
		}
		data := make(map[string]interface{})
		data["ben"] = beneficiary
		data["acts"] = acts

		return c.Render(http.StatusOK, "benpage", models.TemplateData{
			Title:  *beneficiary.FirstName + " " + *beneficiary.LastName +" - Caritas PSS",
			Menu:   models.ReturnMenuByRole(*user.Role),
			Active: models.Clients.URI,
			Data: data,
			},
		)
	case models.Admin:
		//TODO: Прописати функціонал
		return c.Render(http.StatusOK, "error", models.TemplateData{
			Title:  "Помилка - Caritas PSS",
			Menu:   models.ReturnMenuByRole(*user.Role),
			Active: models.Main.URI,
		},
		)
	default:
		return c.Render(http.StatusOK, "error", models.TemplateData{
			Title:  "Помилка - Caritas PSS",
			Menu:   models.ReturnMenuByRole(*user.Role),
			Active: models.Main.URI,
		},
		)
	}
}

func (h *Handler) Journal(c echo.Context) error {
	uuid, err := sess.GetUUIDFromSession(c)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	user, err := h.service.Authorization.GetUserByUUID(uuid)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	return c.Render(http.StatusOK, "journal", models.TemplateData{
		Title:  "Caritas PSS",
		Menu:   models.ReturnMenuByRole(*user.Role),
		Active: models.Journal.URI,
	},
	)
}

func (h *Handler) Profile(c echo.Context) error {
	uuid, err := sess.GetUUIDFromSession(c)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	user, err := h.service.Authorization.GetUserByUUID(uuid)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	org, err := h.service.GetUsersOrganisation(*user.OrgId)
	if err != nil {
		fmt.Println(err)
	}
	return c.Render(http.StatusOK, "profile", models.TemplateData{
		Title:  "Caritas PSS",
		Menu:   models.ReturnMenuByRole(*user.Role),
		Active: models.Profile.URI,
		StringMap: map[string]string{
			"first_name": *user.FirstName,
			"last_name":  *user.LastName,
			"role":       models.GetRoleTitle(*user.Role),
			"email":      *user.Email,
			"phone":      *user.Phone,
			"org_title":  *org.Title,
			"org_code":   *org.Code,
		},
	},
	)
}

func (h *Handler) Logout(c echo.Context) error {
	uuid, err := sess.GetUUIDFromSession(c)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
	}
	err = h.service.Authorization.DeleteUUID(uuid)
	if err != nil {
		fmt.Println(err)
	}
	err = sess.DeleteUUIDFromSession(c, uuid)
	if err != nil {
		fmt.Println(err)
	}

	return c.Redirect(http.StatusTemporaryRedirect, models.Login.URI)
}
