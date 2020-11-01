package generator

func parseTemplateError(err error) {
	if err != nil {
		panic(err)
	}
}

func ioError(err error) {
	if err != nil {
		panic(err)
	}
}

func httpError(err error) {
	if err != nil {
		panic(err)
	}
}
