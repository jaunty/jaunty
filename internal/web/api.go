package web

import (
	"context"
	"time"

	"github.com/disaccord/sigil"
)

func (s *Server) fetchDiscordUser(ctx context.Context, sf string) (*sigil.User, error) {
	cu, ok := s.cache.Get(sf)
	if ok {
		u, ok := cu.(*sigil.User)
		if ok {
			return u, nil
		}
	}

	u, err := s.discord.GetUser(ctx, sf)
	if err != nil {
		return nil, err
	}

	s.cache.Set(sf, u, 0)

	return u, nil
}

func (s *Server) fetchMojangUUIDByUsername(ctx context.Context, username string) (string, error) {
	uid, ok := s.cache.Get("mojang-" + username)
	if ok {
		id, ok := uid.(string)
		if !ok {
			panic("web: cached Mojang UUID is not a string")
		}

		return id, nil
	}

	uuid, err := s.mojang.FetchUUIDByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	s.cache.Set("mojang-"+username, uid, time.Hour*1)
	return uuid, nil
}

func (s *Server) fetchMojangUsernameByUUID(ctx context.Context, uuid string) (string, error) {
	un, ok := s.cache.Get(uuid)
	if ok {
		username, ok := un.(string)
		if !ok {
			panic("web: cached Mojang username is not a string")
		}

		return username, nil
	}

	uns, err := s.mojang.FetchUsernameByUUID(ctx, uuid)
	if err != nil {
		return "", err
	}

	s.cache.Set(uuid, uns, time.Hour*1)
	return uns, nil
}
