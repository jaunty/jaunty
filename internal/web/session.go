package web

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	sessionName             = "jaunty-session-v1"
	sessionDiscordUsername  = "discord-username"
	sessionDiscordSnowflake = "discord-snowflake"
	sessionDiscordAvatar    = "discord-avatar"
)

type session struct {
	s *sessions.Session
}

func (s *Server) getSession(r *http.Request) *session {
	sess, _ := s.store.Get(r, sessionName)
	return &session{s: sess}
}

func (s *Server) destroySession(w http.ResponseWriter, r *http.Request) error {
	sess, _ := s.store.Get(r, sessionName)
	sess.Options.MaxAge = -1
	return sess.Save(r, w)
}

func (s *session) save(w http.ResponseWriter, r *http.Request) error {
	return s.s.Save(r, w)
}

func (s *session) isNew() bool {
	return s.s.IsNew
}

func (s *session) clear() {
	s.s.Values = make(map[interface{}]interface{})
}

func (s *session) setValue(k, v string) {
	s.s.Values[k] = v
}

func (s *session) setUsername(username string) {
	s.setValue(sessionDiscordUsername, username)
}

func (s *session) setSnowflake(username string) {
	s.setValue(sessionDiscordSnowflake, username)
}

func (s *session) setAvatar(avatar string) {
	s.setValue(sessionDiscordAvatar, avatar)
}

func (s *session) getValue(k string) string {
	v, ok := s.s.Values[k].(string)
	if !ok {
		return ""
	}

	return v
}

func (s *session) getUsername() string {
	return s.getValue(sessionDiscordUsername)
}

func (s *session) getAvatar() string {
	return s.getValue(sessionDiscordAvatar)
}

func (s *session) getSnowflake() string {
	return s.getValue(sessionDiscordSnowflake)
}
