package validators

import (
	"fmt"
	"net/mail"
)

type Validator struct {
}

func (v *Validator) ValidateInt(value, min int, max ...int) error {
	if value < min {
		return fmt.Errorf("value contain minimum %d", min)
	}
	if len(max) > 0 && value > max[0] {
		return fmt.Errorf("value contain maximum %d", max)
	}
	return nil
}
func (v *Validator) ValidateFloat(value, min float32, max ...float32) error {
	if value < min {
		return fmt.Errorf("value contain minimum %f", min)
	}
	if len(max) > 0 && value > max[0] {
		return fmt.Errorf("value contain maximum %f", max)
	}
	return nil
}

func (v *Validator) ValidateString(value string, minLength int, maxLength ...int) error {
	n := len(value)
	if n < minLength {
		return fmt.Errorf("must contain minimum %d characters", minLength)
	}
	if len(maxLength) > 0 && n > maxLength[0] {
		return fmt.Errorf("must contain maximum %d characters", maxLength[0])
	}
	return nil
}

func (v *Validator) ValidateEmail(value string) error {
	if err := v.ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}
