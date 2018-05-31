package check

func init() {
	if err := checkEtcd(); err != nil {
		panic(err)
	}

	//if err := checkNSQ(); err != nil {
	//	panic(err)
	//}
}
