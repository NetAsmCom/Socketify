package pipe

// MsgType TODO
type MsgType uint8

// TODO
const (
	InMsgError   MsgType = 0
	InMsgConfig  MsgType = 'c'
	InMsgSendTCP MsgType = 't'
	InMsgSendUDP MsgType = 'u'

	OutMsgConnect    MsgType = 'C'
	OutMsgReceiveTCP MsgType = 'T'
	OutMsgReceiveUDP MsgType = 'U'
	OutMsgDisconnect MsgType = 'D'
)
