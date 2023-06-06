package golog

type structureLogger interface {
	info(any)
	infow(string, ...any)

	warn(any)
	warnw(string, ...any)

	error(any)
	errorw(string, ...any)
}

type unStructureLogger interface {
	raw([]byte)
}
