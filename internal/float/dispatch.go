package float

import (
	"context"
	"fmt"
	"strings"
)

func HandleDispatch(ctx context.Context, endpoint string, payload []byte) error {
	if strings.TrimSpace(endpoint) == "" {
		return fmt.Errorf("empty Float endpoint")
	}
	return PostOutbound(ctx, endpoint, payload)
}
