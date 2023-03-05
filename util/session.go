package util

import (
	"github.com/gorilla/sessions"
	"go_web/config"
	"go_web/domain"
	"log"
	"net/http"
)

/*
对 sessions 库的再封装,实现简单session功能
*/
// 不暴露，保证 session 的单例
type session struct {
	se *sessions.Session
}

// 秘钥，生成唯一 sessionStore
const secretKey = "go-pet-store"

// go web 标准库没有 session，需要自己开发封装或使用第三方的库
var sessionStore = sessions.NewFilesystemStore("", []byte(secretKey))

const sessionName = "session"

// 初始化，通过这个获取唯一 session
func GetSession(r *http.Request) (*session, error) {
	// 设置 fileSystemStore 的最大存储长度，防止溢出
	sessionStore.MaxLength(5 * 4096)
	s, err := sessionStore.Get(r, sessionName)
	if err != nil {
		return nil, err
	}
	return &session{
		s,
	}, nil
}

// 存储和更新，复杂类型存储前需要 gob.Register 进行序列化
func (s *session) Save(key string, val interface{}, w http.ResponseWriter, r *http.Request) error {
	s.se.Values[key] = val
	return s.se.Save(r, w)
}

// 获取值
func (s *session) Get(key string) (result interface{}, ok bool) {
	result, ok = s.se.Values[key]
	return
}

// 删除值
func (s *session) Del(key string, w http.ResponseWriter, r *http.Request) error {
	delete(s.se.Values, key)
	return s.se.Save(r, w)
}

// 从 session 中获取 account
func GetAccountFromSession(r *http.Request) *domain.Account {
	s, err := GetSession(r)
	if err != nil {
		log.Printf("get session error: %v", err.Error())
		return nil
	}
	if s != nil {
		r, ok := s.Get(config.IndexKey)
		if !ok {
			// account 不存在，已登出
			return nil
		}
		a, ok := r.(*domain.Account)
		if !ok {
			log.Print("type assert error *domain.Account")
			return nil
		}
		return a
	}
	log.Print("session get account error: session is nil")
	return nil
}
