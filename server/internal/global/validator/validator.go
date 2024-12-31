package validator

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"

	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// BindAndValid 绑定参数校验
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors

	err := c.ShouldBind(v)
	if err == nil {
		return true, nil
	}

	transl := c.Value("trans")
	trans, _ := transl.(ut.Translator)
	var veers val.ValidationErrors
	ok := errors.As(err, &veers)
	if !ok {
		return false, errs
	}

	for key, value := range veers.Translate(trans) {
		errs = append(errs, &ValidError{
			Key:     key,
			Message: value,
		})
	}
	return false, errs
}
