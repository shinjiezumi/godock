package chat

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合、GetAvatarURLはErrNoAvatarURLを返すべきです")
	}

	testUrl := "http://url-to-avatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}

	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("値が存在する場合、GetAvatarURLはエラーを返すべきではありません")
	} else {
		if url != testUrl {
			t.Error("GetAvatarURLは正しいURLを返すべきです")
		}
	}
}
