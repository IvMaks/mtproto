package proto

import (
	"fmt"
	"github.com/slavaromanov/mtproto"
	"log"
)

type Proto struct {
  MTP *mtproto.MTProto
}

func CreateProto(key int32, hash, dst, file string) *Proto {
        var proto *mtproto.MTProto
	appConfig, err := mtproto.NewConfiguration(key, hash, "0.0.1", "", "", "ru")
	if err != nil {
		log.Fatalf("Create failed: %s\n", err)
	}
	if proto, err = mtproto.NewMTProto(false, dst, file, *appConfig); err != nil {
		log.Fatalf("Create failed: %s\n", err)
	}
	if err = proto.Connect(); err != nil {
		log.Fatalf("Connect failed: %s\n", err)
	}
        return &Proto{MTP: proto}
}

func (p *Proto) Login(phone string) {
	var code string
	authSentCode, err := p.MTP.AuthSendCode(phone)
	if err != nil {
		log.Fatal(err)
	}
	if !authSentCode.Phone_registered {
		log.Fatal("Cannot sign in: Phone isn't registered")
	}
	fmt.Printf("Enter code: ")
	fmt.Scanf("%s", &code)
	a, err := p.MTP.AuthSignIn(phone, code, authSentCode.Phone_code_hash)
	if err != nil {
		log.Fatal(err)
	}
	userSelf := a.User.(mtproto.TL_user)
	log.Printf("Signed in: Id %d name <%s %s>\n", userSelf.Id, userSelf.First_name, userSelf.Last_name)
}

func (p *Proto) Logout() {
	logout, err := p.MTP.AuthLogOut()
	if err != nil {
		log.Fatal(err)
	}
	if !logout {
		log.Println("Can't logout")
	}
	log.Println("You logged out")
}

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
