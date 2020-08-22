package todo

import (
	"context"
	"fmt"
	"github.com/anggakes/gosvelte/src/backend/pkg/events"
)



func SendEmail(ctx context.Context, msg events.Msg) error {

	fmt.Println("send Email")

	return nil
}
