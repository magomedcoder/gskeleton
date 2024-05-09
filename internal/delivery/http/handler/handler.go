package handler

import (
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler/v1"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler/v2"
)

type Handler struct {
	V1 *v1.V1
	V2 *v2.V2
}
