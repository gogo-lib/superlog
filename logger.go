package superlog

type structureLogger interface {
	info(any)
	infow(string, ...interface{})

	warn(any)
	warnw(string, ...interface{})

	error(any)
	errorw(string, ...interface{})
}

type unStructureLogger interface {
	raw([]byte)
}
