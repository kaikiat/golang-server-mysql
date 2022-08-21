package tag_service

import (
	"github.com/kaikiat/golang-server-mysql-template/models"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
}

func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.CreatedBy)
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

func (t *Tag) GetAll() ([]models.Tag, error) {
	tags, err := models.GetTags(t.getMaps())
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}

	return maps
}

func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}
