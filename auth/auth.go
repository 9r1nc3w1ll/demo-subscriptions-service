package auth

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func ValidateToken(ctx context.Context, tokenType string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if len(md.Get("authorization")) != 1 {
			return fmt.Errorf("no authoriziation value found in metadata code: %v", codes.Unauthenticated)
		}

		tokenParts := strings.Split(md.Get("authorization")[0], " ")
		if len(tokenParts) < 2 {
			return fmt.Errorf("token has less than 2 parts: %v", codes.Unauthenticated)
		}

		if tokenParts[0] != tokenType {
			return fmt.Errorf("expected %s token, but got %s code: %v", tokenType, tokenParts[0], codes.Unauthenticated)
		}

		// TODO: Validate token and build probaly add to context
		if tokenParts[1] == "VALID_TEST_TOKEN" {
			return nil
		}

		return fmt.Errorf("invalid token. code: %v", codes.Unauthenticated)
	}
	return nil
}
