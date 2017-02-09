package signalize

import "errors"

type IncomePacket struct {
	Identifier byte
}

func (packet IncomePacket) ToBytes() (result []byte) {
	retVal := []byte{0x02, packet.Identifier, 0x03, 0x0D}
	return retVal
}

func ParseIncomePacket(raw []byte) (result IncomePacket, err error) {
	if len(raw) != 4 {
		return nil, errors.New("Invalid Packet Length")
	}

	return IncomePacket {
		Identifier: raw[2],
	}, nil
}