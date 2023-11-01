package Vbilibili

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synctv-org/synctv/internal/db"
	dbModel "github.com/synctv-org/synctv/internal/model"
	"github.com/synctv-org/synctv/internal/op"
	"github.com/synctv-org/synctv/server/model"
	"github.com/synctv-org/synctv/vendors/bilibili"
)

type MeResp struct {
	IsLogin  bool   `json:"isLogin"`
	Username string `json:"username"`
	Face     string `json:"face"`
	IsVip    bool   `json:"isVip"`
}

func Me(ctx *gin.Context) {
	user := ctx.MustGet("user").(*op.User)
	vendor, err := db.FirstOrCreateVendorByUserIDAndVendor(user.ID, dbModel.StreamingVendorBilibili)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}
	cli := bilibili.NewClient(vendor.Cookies)
	nav, err := cli.UserInfo()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}
	ctx.JSON(http.StatusOK, model.NewApiDataResp(MeResp{
		IsLogin:  nav.Data.IsLogin,
		Username: nav.Data.Uname,
		Face:     nav.Data.Face,
		IsVip:    nav.Data.VipStatus == 1,
	}))
}
