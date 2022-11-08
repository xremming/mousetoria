package validate

import (
	"github.com/hashicorp/go-multierror"
)

type Validateable interface {
	Validate() error
}

func ValidateAll(vs ...Validateable) error {
	var err error

	for _, v := range vs {
		if errV := v.Validate(); errV != nil {
			err = multierror.Append(err, errV)
		}
	}

	return err
}
