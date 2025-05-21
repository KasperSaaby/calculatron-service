package validate

import "errors"

type ValidatorFunc func() error

func (v ValidatorFunc) Validate() error {
	return v()
}

type Validator interface {
	Validate() error
}

func All(validations ...Validator) error {
	var errs []error
	for _, v := range validations {
		if v == nil {
			continue
		}

		err := v.Validate()
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}
