package sess

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

//GetUUIDFromSession gets UUID from cookies on the client device
func GetUUIDFromSession(c echo.Context) (string, error) {
	sess, err := session.Get("session", c)
	if err != nil {
		return "", err
	}
	if val, ok := sess.Values["uuid"].(string); ok {
		// Значення поля "username" є рядком.
		// Ви можете використати змінну "val" для отримання значення.
		return val, nil
	} else {
		return "", fmt.Errorf("uuid is absent in the cookies")
	}
}

//SaveUUIDToSession saves UUID to cookiec on the client`s device
func SaveUUIDToSession(c echo.Context, uuid string) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["uuid"] = uuid
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}
	return nil
}

func DeleteUUIDFromSession(c echo.Context, uuid string) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	delete(sess.Values, "uuid")
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}
	return nil
}