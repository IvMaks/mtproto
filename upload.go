package mtproto

func (m *MTProto) UploadGetFile(location TL,limit int32) (*TL, error){
	return m.InvokeSync(TL_upload_getFile{Location:location,Limit:limit})
}
