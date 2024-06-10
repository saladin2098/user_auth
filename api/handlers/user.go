package handlers
import (
	genproto "github.com/Mubinabd/auth_service/genproto"
	"github.com/gin-gonic/gin"
)
// @Router 				/user/register [post]
// @Summary 			REGISTER USER
// @Description		 	This api registers user
// @Tags 				USER
// @Accept 				json
// @Produce 			json
// @Param data 			body genproto.UserCreate true "UserCreate"
// @Success 201 		{object} genproto.User
// @Failure 400 		string Error
func (h *Handler)RegisterUser(c *gin.Context) {
	var req genproto.UserCreate
    if err := c.ShouldBindJSON(&req); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    user, err := h.UserService.RegisterUser(&req)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, user)
}
// @Router 				/user/login [post]
// @Summary 			Login USER
// @Description		 	This api logs  user in
// @Tags 				USER
// @Accept 				json
// @Produce 			json
// @Param data 			body genproto.LoginReq true "LoginReq"
// @Success 201 		{object} genproto.Token
// @Failure 400 		string Error
func (h *Handler)LoginUser(c *gin.Context) {
	var req genproto.LoginReq
    if err := c.ShouldBindJSON(&req); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    token, err := h.UserService.Loginuser(&req)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, token)
}

// @Router 				/user/info [get]
// @Summary 			GET USER
// @Description		 	This api GETS user by username
// @Tags 				USER
// @Accept 				json
// @Produce 			json
// @Security    		BearerAuth
// @Param 			    username path string true "USERNAME"
// @Success 200			{object} genproto.User
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *Handler) GetUserInfo(c *gin.Context) {
	var req genproto.ByUsername
	username := c.Param("username")
	req.Username = username
    if err := c.ShouldBindJSON(&req); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    user, err := h.UserService.GetUserInfo(&req)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, user)
}