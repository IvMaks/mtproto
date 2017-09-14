package proto

import (
	"fmt"
	"log"

	"github.com/slavaromanov/mtproto"
)

type Proto struct {
	MTP *mtproto.MTProto
}

func CreateProto(key int32, hash, dst, file string) *Proto {
	var proto = new(mtproto.MTProto)
	appConfig, err := mtproto.NewConfiguration(key, hash, version, device, system, "EN")
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
