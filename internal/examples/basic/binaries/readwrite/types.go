package readwrite

const (
	maxDataCount = 7
)

type (
	header struct {
		MsgNo     uint16
		DataCount byte
		Dummy     byte
	}

	data struct {
		Id    uint32
		Value [4]byte
	}

	message struct {
		Header header
		Data   [maxDataCount]data
	}
)
