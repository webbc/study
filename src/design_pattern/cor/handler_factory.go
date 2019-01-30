package cor

func CreateHandler() Handler {

	sh := &SalesHandler{}
	mh := &ManagerHandler{}
	ch := &CeoHandler{}

	sh.successor = mh
	mh.successor = ch

	return sh
}
