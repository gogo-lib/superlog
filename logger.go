package superlog

type structureLogger interface {
	info(string)
	infow(string, ...interface{})

	warn(string)
	warnw(string, ...interface{})

	error(string)
	errorw(string, ...interface{})
}

type unStructureLogger interface {
	raw(string)
}
