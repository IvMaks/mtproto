package mtproto

import "errors"

func (m *MTProto) ContactsGetContacts(hash string) (*TL, error) {
	return m.InvokeSync(TL_contacts_getContacts{
		Hash: hash,
	})
}

func (m *MTProto) ContactsGetTopPeers(correspondents, botsPM, botsInline, groups, channels bool, offset, limit, hash int32) (*TL, error) {
	tl, err := m.InvokeSync(TL_contacts_getTopPeers{
		Correspondents: correspondents,
		Bots_pm:        botsPM,
		Bots_inline:    botsInline,
		Groups:         groups,
		Channels:       channels,
		Offset:         offset,
		Limit:          limit,
		Hash:           hash,
	})

	if err != nil {
		return nil, err
	}

	switch (*tl).(type) {
	case TL_contacts_topPeersNotModified:
	case TL_contacts_topPeers:
	default:
		return nil, errors.New("MTProto::ContactsGetTopPeers error: Unknown type")
	}

	return tl, nil
}

func (m *MTProto) AddContactsByPhone(phoneNumbers []string) (*TL, error){
	contacts := make([]TL,0)
	for _,phone := range phoneNumbers{
		contacts = append(contacts,TL_inputPhoneContact{
			Phone:phone,
		})
	}
	return m.InvokeSync(TL_contacts_importContacts{
		Contacts:contacts,
		Replace:TL_boolTrue{},
	})
}
func (m *MTProto) AddContactsByPhones(phoneNumbers []string) (*TL, error){
	contacts := make([]TL,0)
	for p := range phoneNumbers {
		contacts = append(contacts, TL_inputPhoneContact{
			Phone: phoneNumbers[p],
		})
	}
	return m.InvokeSync(TL_contacts_importContacts{
		Contacts:contacts,
		Replace:TL_boolTrue{},
	})
}
