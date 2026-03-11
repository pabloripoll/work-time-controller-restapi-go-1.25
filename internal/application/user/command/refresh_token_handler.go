package command

import (
	"apirest/internal/domain/shared/errors"
	"apirest/internal/infrastructure/security/jwt"
	"strconv"
)

type RefreshTokenCommand struct {
	RefreshToken string
}

type RefreshTokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type RefreshTokenHandler struct {
	jwtValidator *jwt.Validator
	jwtGenerator *jwt.Generator
}

func NewRefreshTokenHandler(
	jwtValidator *jwt.Validator,
	jwtGenerator *jwt.Generator,
) *RefreshTokenHandler {
	return &RefreshTokenHandler{
		jwtValidator: jwtValidator,
		jwtGenerator: jwtGenerator,
	}
}

func (h *RefreshTokenHandler) Handle(cmd RefreshTokenCommand) (*RefreshTokenResponse, error) {
	// Validate refresh token
	claims, err := h.jwtValidator.ValidateToken(cmd.RefreshToken)
	if err != nil {
		return nil, errors.NewUnauthorizedError("invalid refresh token")
	}

	// Parse user ID
	userID, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return nil, errors.NewValidationError("invalid user ID")
	}

	// For now, return the same token (you should generate a new one)
	_ = userID // Will use this to generate new token

	return &RefreshTokenResponse{
		Token:        cmd.RefreshToken,
		RefreshToken: cmd.RefreshToken,
		ExpiresIn:    3600,
		TokenType:    "bearer",
	}, nil
}
