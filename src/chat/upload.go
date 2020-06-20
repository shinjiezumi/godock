package chat

import (
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func UploaderHandler(w http.ResponseWriter, req *http.Request) {
	userID := req.FormValue("user_id")
	file, header, err := req.FormFile("avatarFile")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	filename := filepath.Join("chat/avatars", userID+filepath.Ext(header.Filename))
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, "成功")
}
