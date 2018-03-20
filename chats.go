package mtproto

func (m *MTProto) ChatsCreateChat(title string,users []TL) (*TL,error){
	return m.InvokeSync(TL_messages_createChat{
			Users:users,
			Title:title,
		})
}

func (m *MTProto) ChatsGetAllChats(exceptsIds []int32) (*TL,error){
	return m.InvokeSync(TL_messages_getAllChats{
		Except_ids: exceptsIds,
	})
}
func (m *MTProto) ChatsAddChatUser(chatId int32, user TL, fwdLimit int32) (*TL,error){
	return m.InvokeSync(TL_messages_addChatUser{
		Chat_id:chatId,
		User_id:user,
		Fwd_limit:fwdLimit,
	})
}