package zfClient

import "net/http"

// ZFCookie 维护正方Cookie参数
type ZFCookie struct {
	JSessionID string
	Route      string
}

func (z *ZFCookie) ToCookie() []*http.Cookie {
	cookies := make([]*http.Cookie, 0)
	if z.JSessionID != "" {
		cookies = append(cookies, &http.Cookie{
			Name:  "JSESSIONID",
			Value: z.JSessionID,
		})
	}
	if z.Route != "" {
		cookies = append(cookies, &http.Cookie{
			Name:  "route",
			Value: z.Route,
		})
	}
	return cookies
}

func FromCookie(cookies []*http.Cookie) *ZFCookie {
	t := &ZFCookie{}
	for _, cookie := range cookies {
		if cookie == nil {
			continue
		}
		if cookie.Name == "JSESSIONID" {
			t.JSessionID = cookie.Value
		} else if cookie.Name == "route" {
			t.Route = cookie.Value
		}
	}
	return t
}
