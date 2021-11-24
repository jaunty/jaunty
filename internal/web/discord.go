package web

import (
	"context"

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

	ur := s.discord.User()
	u, err := ur.Get(ctx)
	if err != nil {
		return nil, err
	}

	s.cache.Set(sf, u, 0)

	return u, nil
}
