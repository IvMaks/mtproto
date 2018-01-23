package proto

import (
	"log"

	"github.com/slavaromanov/mtproto"
)

func (p *Proto) GetChans() map[int32]string {
	channels, err := p.MTP.GetChannels()
	if err != nil {
		log.Fatal(err)
	}
	var res = map[int32]string{}
	for _, c := range channels {
		if c.Flags == 8296 {
			res[c.Id] = c.Title
		}
	}
	return res
}

func (p *Proto) GetUsers(i, off, lim int32, filterName string) []mtproto.TL {
	channels, err := p.MTP.GetChannels()
	if err != nil {
		log.Fatal(err)
	}

	var ch *mtproto.TL_inputChannel
	for _, c := range channels {
		if c.Id == i {
			ch = &mtproto.TL_inputChannel{
				Channel_id: i, Access_hash: c.Access_hash,
			}
		}
	}
	if ch == nil {
		log.Fatalln("Can't find channel")
	}

	var filter mtproto.TL
	switch filterName {
	case "kicked":
		filter = mtproto.TL_channelParticipantsKicked{}
	case "admins":
		filter = mtproto.TL_channelParticipantsAdmins{}
	case "recent":
		filter = mtproto.TL_channelParticipantsRecent{}
	case "bots":
		filter = mtproto.TL_channelParticipantsBots{}
	}

	users, err := p.MTP.GetMembers(*ch, filter, off, lim)
	if err != nil {
		log.Fatal(err)
	}
	return users
}

func (p *Proto) GetMembers(i, off, lim int32, filterName string) (us []int32) {
	users := p.GetUsers(i, off, lim, filterName)
	for _, u := range users {
		us = append(us, u.(mtproto.TL_user).Id)
	}
	return us
}
