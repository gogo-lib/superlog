package superlog

import "io"

type unStructureLog struct {
	writer io.Writer
}

func (u unStructureLog) raw(msg []byte) {
	u.writer.Write(msg)
}
