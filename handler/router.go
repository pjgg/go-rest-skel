package handler

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"

	cErrors "github.com/pjgg/go-errors/errors"
	"github.com/pjgg/go-rest-skel/utils"

	"github.com/gin-gonic/gin"
)

func SetupPersonRouting(personHandler *PersonHandler) *gin.Engine {
	router := gin.Default()
	router.Use(addUniqueRequestID)
	personHandler.setupPersonAPIRouting(router.Group("/v1/person", checkContentTypeApplicationJSON))

	return router
}

func addUniqueRequestID(c *gin.Context) {
	uniqueID := rand.New(rand.NewSource(time.Now().UnixNano()))
	c.Set(utils.Uid, strconv.Itoa(uniqueID.Int()))
}

func checkContentTypeApplicationJSON(c *gin.Context) {
	var err error
	if !strings.Contains(c.ContentType(), utils.ApplicationJSON) {
		err = cErrors.NewContentTypeError(errors.New("unexpected content type: "+c.ContentType()), utils.Uid)
	}

	if err != nil {
		errorDto := cErrors.ErrorHandlerInstance.Error(err, utils.GetDefaultWebErrorTags(c))
		c.Abort()
		c.JSON(errorDto.Status, err)
	}

}
