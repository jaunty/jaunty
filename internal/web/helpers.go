package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/disaccord/sigil"
	"github.com/disaccord/sigil/embuilder"
	"github.com/jaunty/jaunty/internal/web/templates"
)

func (s *Server) recoverFunc(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.ErrorPage{
		BasePage: s.basePage(r),
		Message:  "Internal Server Error",
	})
}

func followupError(ctx context.Context, msg string, args ...interface{}) *sigil.InteractionResponse {
	msg = fmt.Sprintf(msg, args...)

	e := embuilder.NewEmbed(
		embuilder.Title("Error during Interaction Execution"),
		embuilder.Field("Status Report", msg),
		embuilder.Color(14029073),
	)

	return &sigil.InteractionResponse{
		Type: sigil.InteractionCallbackTypeChannelMessageWithSource,
		Data: &sigil.Message{
			Embeds: []*sigil.Embed{e},
		},
	}
}
