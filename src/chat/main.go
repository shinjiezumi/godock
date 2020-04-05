package chat

import "net/http"

func Main(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
<html>
    <head>
        <title>チャット</title>
    </head>
    <body>
        チャットしましょう!
    </body>
</html>
		`))
}
