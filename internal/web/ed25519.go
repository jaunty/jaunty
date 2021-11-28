package web

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"io"
	"net/http"

	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

func (s *Server) verifyInteraction(r *http.Request) bool {
	ctx := r.Context()

	dk, err := hex.DecodeString(s.publicKey)
	if err != nil {
		ctxlog.Error(ctx, "unable to decode public key", zap.Error(err))
		return false
	}

	var msg bytes.Buffer

	signature := r.Header.Get("X-Signature-Ed25519")
	if signature == "" {
		ctxlog.Warn(ctx, "discord sent a blank signature")
		return false
	}

	sig, err := hex.DecodeString(signature)
	if err != nil {
		ctxlog.Error(ctx, "error decoding signature", zap.Error(err))
		return false
	}

	timestamp := r.Header.Get("X-Signature-Timestamp")
	if timestamp == "" {
		ctxlog.Warn(ctx, "discord sent a blank timestamp")
		return false
	}

	msg.WriteString(timestamp)

	defer r.Body.Close()
	var bod bytes.Buffer

	defer func() {
		r.Body = io.NopCloser(&bod)
	}()

	if _, err := io.Copy(&msg, io.TeeReader(r.Body, &bod)); err != nil {
		ctxlog.Error(ctx, "error copying request body to buffer", zap.Error(err))
		return false
	}

	return ed25519.Verify(dk, msg.Bytes(), sig)
}
