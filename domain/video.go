package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID         string    `valid:"uuid"`
	ResourceID string    `valid:"required"`
	FilePath   string    `valid:"required"`
	CreatedAt  time.Time `valid:"-"`
}

func Init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{}
}

func (v *Video) Validate() error {
	_, err := govalidator.ValidateStruct(v)

	return err
}
