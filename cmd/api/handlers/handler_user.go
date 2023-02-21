package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/base"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

type LoginReq struct {
	Username string `json:"username,required" query:"username,required"`
	Password string `json:"password,required" query:"password,required"`
}

type User struct {
	UserID int64
}

var (
	identityKey = "id"
	UserAuthMw  *jwt.HertzJWTMiddleware
)

type RegisterReq struct {
	Username string `query:"username,required"`
	Password string `query:"password,required"`
}

type UserInfoReq struct {
	UserID int64 `query:"user_id,required"`
}

func Init() {
	var err error
	UserAuthMw, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte("douyin"),
		TokenLookup: "query:token",
		// login process
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var l LoginReq
			if err = c.BindAndValidate(&l); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			req := &user.CheckUserRequest{
				Username: l.Username,
				Password: l.Password,
			}
			userID, err := rpc.CheckUser(context.Background(), req)
			c.Set(identityKey, userID)
			if err != nil {
				return nil, err
			}
			return &User{
				UserID: userID,
			}, nil
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserID,
				}
			}
			return jwt.MapClaims{}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, codeStatusOK int, token string, time time.Time) {
			c.JSON(codeStatusOK, map[string]interface{}{
				"status_code": "0",
				"status_msg":  "success",
				"user_id":     c.GetInt64(identityKey),
				"token":       token,
			})
		},

		// Auth process
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			userID, _ := claims[identityKey].(float64)
			c.Set(identityKey, int64(userID))
			return &User{
				UserID: int64(userID),
			}
		},
	})
	if err != nil {
		panic("JWT Middleware Error: " + err.Error())
	}
}

func Register(ctx context.Context, c *app.RequestContext) {
	var req RegisterReq
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status_code": "-1",
			"status_msg":  err.Error(),
			"user_id":     0,
			"token":       "",
		})
		return
	}
	rpcReq := &user.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	}
	userID, err := rpc.CreateUser(context.Background(), rpcReq)
	if err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status_code": Err.ErrCode,
			"status_msg":  Err.ErrMsg,
			"user_id":     0,
			"token":       "",
		})
		return
	}

	tokenStr, _, err := UserAuthMw.TokenGenerator(&User{UserID: userID})
	if err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status_code": Err.ErrCode,
			"status_msg":  Err.ErrMsg,
			"user_id":     0,
			"token":       "",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"status_code": "0",
		"status_msg":  "success",
		"user_id":     userID,
		"token":       tokenStr,
	})
}

func UserInfo(ctx context.Context, c *app.RequestContext) {
	TokenUserID := c.GetInt64(identityKey)

	var req UserInfoReq
	if err := c.BindAndValidate(&req); err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status_code": Err.ErrCode,
			"status_msg":  Err.ErrMsg,
			"user":        nil,
		})
		return
	}

	rpcReq := &user.MGetUserRequest{
		UserIds: []int64{req.UserID},
		Token:   &base.Token{UserId: &TokenUserID},
	}
	users, err := rpc.MGetUser(ctx, rpcReq)
	if err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status_code": Err.ErrCode,
			"status_msg":  Err.ErrMsg,
			"user":        nil,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status_code": "0",
		"status_msg":  "success",
		"user":        users[0],
	})
	return
}
