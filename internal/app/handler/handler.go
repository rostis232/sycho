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

func (h *Handler) RegistrationGET(c echo.Context) error {
	return c.Render(http.StatusOK, "register", models.TemplateData{
		Title:  "Регістрація",
		Menu:   models.UnLoggedUserMenu,
		Active: models.Login.URI,
	},
	)
}

func (h *Handler) RegistrationPOST(c echo.Context) error {

	return c.Render(http.StatusOK, "register", models.TemplateData{
		Title:  "Регістрація",
		Menu:   models.UnLoggedUserMenu,
		Active: models.Login.URI,
	},
	)
}

func (h *Handler) LogInGet(c echo.Context) error {
	return c.Render(http.StatusOK, "login", models.TemplateData{
		Title:  "Авторизація",
		Menu:   models.UnLoggedUserMenu,
		Active: models.Login.URI,
	},
	)
}

func (h *Handler) LogInPost(c echo.Context) error {
	uuid, err := h.service.Authorization.CreateAndSaveUUID(c.FormValue("login"), c.FormValue("pass"))
	fmt.Println(uuid, err)
	if err == sql.ErrNoRows {
		return c.Render(http.StatusOK, "login", models.TemplateData{
			Title:  "Авторизація",
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

func (h *Handler) Help(c echo.Context) error {
	uuid, err := sess.GetUUIDFromSession(c)
	if err != nil {
		fmt.Println(err)
		return c.Render(http.StatusOK, "help", models.TemplateData{
			Title:  "Допомога",
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
		Title:  "Sycho",
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
				Title:  "Помилка",
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
			Title:  "Помилка",
			Menu:   models.ReturnMenuByRole(*user.Role),
			Active: models.Main.URI,
		},
		)
	default:
		return c.Render(http.StatusOK, "error", models.TemplateData{
			Title:  "Помилка",
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
		Title:  "Sycho",
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
	return c.Render(http.StatusOK, "profile", models.TemplateData{
		Title:  "Sycho",
		Menu:   models.ReturnMenuByRole(*user.Role),
		Active: models.Profile.URI,
		StringMap: map[string]string{
			"first_name": *user.FirstName,
			"last_name":  *user.LastName,
			"role":       models.GetRoleTitle(*user.Role),
			"email":      *user.Email,
			"phone":      *user.Phone,
		},
	},
	)
}


