package handler

import (
	"github.com/shohhei1126/bbs-go/dao"
	"github.com/shohhei1126/bbs-go/http/response"
	"github.com/shohhei1126/bbs-go/log"
	"github.com/shohhei1126/bbs-go/service"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

type Thread struct {
	threadService service.Thread
}

func NewThread(threadService service.Thread) *Thread {
	return &Thread{threadService: threadService}
}

func (t Thread) List(ctx context.Context, r *http.Request) response.Response {
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil {
		return response.BadRequest
	}
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if err != nil {
		return response.BadRequest
	}

	paging := dao.Paging{Limit: uint64(limit), Offset: uint64(offset), OrderBy: "updated_at desc"}
	threads, err := t.threadService.FindThreads(paging)
	if err != nil {
		log.Logger.Error(err)
		return response.ServerError
	}

	return response.Json(http.StatusOK, threads)
}
