package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/kaikiat/golang-server-mysql-template/pkg/app"
	"github.com/kaikiat/golang-server-mysql-template/pkg/e"
	"github.com/kaikiat/golang-server-mysql-template/service/tag_service"
	"github.com/unknwon/com"
)

type AddTagForm struct {
	Name      string `json:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `json:"created_by" valid:"Required;MaxSize(100)"`
}

// @Summary Add article tag
// @Produce  json
// @Param name body string true "Name"
// @Param created_by body string false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /tags [post]
func AddTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddTagForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := tag_service.Tag{
		Name:      form.Name,
		CreatedBy: form.CreatedBy,
	}

	exists, err := tagService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	err = tagService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /tags [get]
func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	tagService := tag_service.Tag{
		Name: name,
	}
	tags, err := tagService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	count, err := tagService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": tags,
		"total": count,
	})
}

// @Summary Delete article tag
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()

	valid.Min(id, 1, "id").Message("ID must be greater than 0")

	if valid.HasErrors() {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	tagService := tag_service.Tag{ID: id}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	if err := tagService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
