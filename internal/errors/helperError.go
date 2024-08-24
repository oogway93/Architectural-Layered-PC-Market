package errors

import "log"

func PanicError(err error) error {
	if err != nil {
		log.Fatalf(err.Error())
		// panic(err)
	}
	return nil
}