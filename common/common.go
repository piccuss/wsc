package common

//Must panic when err is not nil
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

//Must2 panic when 2nd param is not nil, otherwise return 1st value
func Must2(value interface{}, err error) interface{} {
	Must(err)
	return value
}

//Error2 return error from 2nd param as error
func Error2(v interface{}, err error) error {
	return err
}
