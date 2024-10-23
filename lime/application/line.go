package application

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
)

type LineUsecase interface {
	// Err: ErrAuthInvalidArg
	VerifySignature(ctx context.Context, lineSignature string, body []byte) error
}

type lineUsecase struct {
	channelSecret string
}

func NewLineUsecase(channelSecret string) LineUsecase {
	return &lineUsecase{
		channelSecret: channelSecret,
	}
}

func (line *lineUsecase) VerifySignature(ctx context.Context, lineSignature string, body []byte) error {
	slog.Info("VerifySignature", "lineSignature", lineSignature, "body", string(body))

	decoded, err := base64.StdEncoding.DecodeString(lineSignature)
	if err != nil {
		return errors.Join(fmt.Errorf("failed decode line signature. err: %w", err), ErrAuthInvalidArg)
	}

	hash := hmac.New(sha256.New, []byte(line.channelSecret))
	if _, err := hash.Write(body); err != nil {
		return errors.Join(fmt.Errorf("failed hash.Write. err: %w", err), ErrAuthInvalidArg)
	}

	if !hmac.Equal(decoded, hash.Sum(nil)) {
		return errors.Join(errors.New("invalid signature"), ErrAuthInvalidArg)
	}
	return nil
}
