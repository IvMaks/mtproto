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
func (m *MTProto) ChannelsCreateChannel(broadcast, megagroup bool, title, about string) (*TL, error) {
	return m.InvokeSync(TL_channels_createChannel{
		Broadcast: broadcast,
		Megagroup: megagroup,
		Title:     title,
		About:     about,
	})
}

func (m *MTProto) ChannelsDeleteChannel(channel TL) (*TL, error) {
	return m.InvokeSync(TL_channels_deleteChannel{
		Channel: channel,
	})
}

func (m *MTProto) ChannelsKickFromChannel(channel TL, userId TL, kicked TL) (*TL, error) {
	return m.InvokeSync(TL_channels_kickFromChannel{
		Channel: channel,
		User_id: userId,
		Kicked:  kicked,
	})
}

func (m *MTProto) ChannelsEditAbout(channel TL, about string) (*TL, error) {
	return m.InvokeSync(TL_channels_editAbout{
		Channel: channel,
		About:   about,
	})
}

func (m *MTProto) ChannelsEditTitle(channel TL, title string) (*TL, error) {
	return m.InvokeSync(TL_channels_editTitle{
		Channel: channel,
		Title:   title,
	})
}

func (m *MTProto) ChannelsInviteToChannel(channel TL, users []TL) (*TL, error) {
	return m.InvokeSync(TL_channels_inviteToChannel{
		Channel: channel,
		Users:   users,
	})
}

func (m *MTProto) ChannelsEditAdmin(channel TL, userId TL, role TL) (*TL, error) {
	return m.InvokeSync(TL_channels_editAdmin{
		Channel: channel,
		User_id: userId,
		Role:    role,
	})
}

func (m *MTProto) ChannelsToggleInvites(channel TL, enabled TL) (*TL, error) {
	return m.InvokeSync(TL_channels_toggleInvites{
		Channel: channel,
		Enabled: enabled,
	})
}

func (m *MTProto) ChannelsToggleSignatures(channel TL, enabled TL) (*TL, error) {
	return m.InvokeSync(TL_channels_toggleSignatures{
		Channel: channel,
		Enabled: enabled,
	})
}
