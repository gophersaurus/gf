package generate

import (
	"errors"
	"fmt"
)

func Controller(args []string) error {
	if len(args) < 1 {
		return errors.New("expected to receive the controller name as an argument")
	}
	fmt.Println(args[0])
	return nil
}
