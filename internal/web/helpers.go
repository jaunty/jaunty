package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/disaccord/sigil"
	"github.com/disaccord/sigil/embuilder"
	"github.com/jaunty/jaunty/internal/web/templates"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
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

func respondJSON(ctx context.Context, w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		ctxlog.Error(ctx, "unable to marshal json", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
