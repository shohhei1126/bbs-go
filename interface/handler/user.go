package handler

import (
	"github.com/shohhei1126/bbs-go/common/http/response"
	"github.com/shohhei1126/bbs-go/interface/service"
	"goji.io/pat"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
	"database/sql"
	"github.com/shohhei1126/bbs-go/common/log"
)

type User struct {
	userService service.User
}

func NewUser(userService service.User) *User {
	return &User{userService: userService}
}

func (u User) Show(ctx context.Context, r *http.Request) response.Response {
	id, err := strconv.ParseInt(pat.Param(ctx, "id"), 10, 64)
	if err != nil {
		return response.ServerError
	}
	user, err := u.userService.FindById(uint32(id))
	if err == sql.ErrNoRows {
		return response.NotFound
	} else if err != nil {
		log.Logger.Error(err)
		return response.ServerError
	}
	return response.Json(http.StatusOK, user)
}
