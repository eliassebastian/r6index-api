package ubisoft

import (
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/eliassebastian/r6index-api/pkg/auth"
)

func requestHeaders(req *protocol.Request, auth *auth.UbisoftSession, new bool) {

	si := auth.SessionId
	av := "t=" + auth.Ticket
	ai := UBISOFT_APPID
	exp := auth.Expiration

	if new {
		si = auth.SessionIdNew
		av = "t=" + auth.TicketNew
		ai = UBISOFT_NEWAPPID
		exp = auth.ExpirationNew
		req.Header.Set("User-Agent", UBISOFT_USERAGENT)
	}

	req.Header.Set("Connection", "keep-alive")
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.Header.Add("Accept", "*/*")
	req.Header.Set("Ubi-SessionId", si)
	req.Header.Set("Ubi-AppId", ai)
	req.Header.Set("expiration", exp)
	req.SetAuthSchemeToken("Ubi_v1", av)

}
