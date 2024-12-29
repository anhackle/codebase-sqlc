package initialize

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/utils/validation"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() {
	// Use for bind tag
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", validation.ValidatePassword)
	}

	// Use for validator package
	validate := validator.New()
	global.Validator = validate
	global.Validator.RegisterValidation("password", validation.ValidatePassword)
}
