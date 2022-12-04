package member

import (
	"context"
	sh "github.com/geb/aweproj/book-store/shared"
	member_dto "github.com/geb/aweproj/book-store/shared/dto/member"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// All godoc
// @Tags member-controller
// @Summary Member
// @Description Put all mandatory parameter
// @Param Accept-Language header string true "EN" default(EN)
// @Accept json
// @Produce json
// @Success 200
// @Router /book-store/member/user/login [post]
func (c *Controller) Login(ec echo.Context) error {

	return sh.Response(ec, "login", nil, nil)

}

// All godoc
// @Tags member-controller
// @Summary Member
// @Description Put all mandatory parameter
// @Param Accept-Language header string true "EN" default(EN)
// @Param UserDto body member.UserReq true "UserReq"
// @Accept json
// @Produce json
// @Success 200
// @Router /book-store/member/user/isUserValid [post]
func (c *Controller) IsUserValid(ec echo.Context) error {
	var (
		ctx     = context.Background()
		request member_dto.UserReq
	)

	if err := ec.Bind(&request); err != nil {
		return sh.Response(ec, "", nil, sh.New(sh.BAD_REQUEST, err))
	}

	if err := ec.Validate(&request); err != nil {
		return sh.Response(ec, "", nil, sh.New(sh.BAD_REQUEST, err))
	}

	token, err := c.InterfacesHolder.MemberViewService.Login(ctx, request.Username, request.Password)

	if err != nil {
		ec.SetCookie(&http.Cookie{
			Name:       "X-Token",
			Value:      "",
			Expires:    time.Now(),
			RawExpires: "",
		})
		return sh.Response(ec, "", nil, sh.New(sh.BAD_REQUEST, err))
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   *token,
		Expires: time.Now().Add(72 * time.Hour),
		Path:    "/",
	}

	ec.SetCookie(cookie)

	ec.Request().AddCookie(cookie)
	return ec.Redirect(http.StatusFound, "/book-store/inventory/landing_page")

}

// All godoc
// @Tags member-controller
// @Summary Member
// @Description Put all mandatory parameter
// @Param Accept-Language header string true "EN" default(EN)
// @Param BookDto body member.UserReq true "UserReq"
// @Accept json
// @Produce json
// @Success 200
// @Router /book-store/member/user/create [post]
func (c *Controller) CreateUser(ec echo.Context) error {
	var (
		ctx     = context.Background()
		request member_dto.UserReq
	)

	if err := ec.Bind(&request); err != nil {
		return sh.Response(ec, "", nil, sh.New(sh.BAD_REQUEST, err))
	}

	if err := ec.Validate(&request); err != nil {
		return sh.Response(ec, "", nil, sh.New(sh.BAD_REQUEST, err))
	}

	err := c.InterfacesHolder.MemberViewService.CreateUser(ctx, request.Username, request.Password)

	if err != nil {
		return sh.Response(ec, "", nil, sh.New(sh.BAD_REQUEST, err))
	}

	return sh.Response(ec, "", nil, nil)

}
