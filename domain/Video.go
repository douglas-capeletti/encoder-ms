package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID         string    `valid:"uuid"`
	ResourceID string    `valid:"notnull"`
	FilePath   string    `valid:"notnull"`
	CreatedAt  time.Time `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (video *Video) Validate() error {
	valid, err := govalidator.ValidateStruct(video)
	if !valid {
		return err
	}
	return nil
}

func NewVideo() *Video {
	return &Video{}
}
