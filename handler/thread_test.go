package handler

import (
	"github.com/golang/mock/gomock"
	"github.com/shohhei1126/bbs-go/dao"
	"github.com/shohhei1126/bbs-go/http/response"
	"github.com/shohhei1126/bbs-go/model"
	"github.com/shohhei1126/bbs-go/service"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"testing"
)

func TestThreadList(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	threadServiceMock := service.NewMockThread(ctl)
	threadServiceMock.
		EXPECT().
		FindThreads(dao.Paging{Limit: 5, Offset: 0, OrderBy: "updated_at desc"}).
		Return(model.ThreadSlice{}, nil).
		Times(1)
	threadHandler := NewThread(threadServiceMock)

	ctx := context.Background()
	{
		r := http.Request{}
		url, err := url.Parse("http://localhost?limit=a&offset=0")
		if err != nil {
			t.Fatal(err)
		}
		r.URL = url
		res := threadHandler.List(ctx, &r)
		assert.Equal(t, response.BadRequest, res, "")
	}

	{
		r := http.Request{}
		url, err := url.Parse("http://localhost?limit=5&offset=0")
		if err != nil {
			t.Fatal(err)
		}
		r.URL = url
		threadHandler.List(ctx, &r)
	}
}
