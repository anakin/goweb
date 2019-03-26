package session

import (
	"github.com/astaxie/beego/session"
)

type Sess struct {
	M *session.Manager
}

func Init() *Sess {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	sess, _ := session.NewManager("memory", sessionConfig)
	return &Sess{
		M: sess,
	}
}
