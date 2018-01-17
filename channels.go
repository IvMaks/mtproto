package mtproto

import "fmt"

func (m *MTProto) GetChannels() (channels []TL_channel, err error) {
        if dialogs, err := m.MessagesGetDialogs(true, int32(0), int32(0), TL_inputPeerEmpty{}, int32(100)); err == nil {
		for _, chat := range (*dialogs).(TL_messages_dialogsSlice).Chats {
			switch chat.(type) {
			case TL_channel:
				channels = append(channels, chat.(TL_channel))
			default:
				continue
			}
		}
		return channels, nil
	}
	return nil, err
}

func (m *MTProto) GetMembers(channel, filter TL, off, limit int32) ([]TL, error) {
        var resp = make(chan response, 1)
	m.queueSend <- packetToSend{
		msg: TL_channels_getParticipants{
			Channel: channel,
			Filter:  filter,
			Offset:  off,
			Limit:   limit,
		},
		resp: resp,
	}
	x := <-resp
	if x.err != nil {
		return nil, x.err
	}
	switch x.data.(type) {
	case TL_channels_channelParticipants:
		return x.data.(TL_channels_channelParticipants).Users, nil
	default:
		return nil, fmt.Errorf("Connection error: got: %T", x)
	}
}
