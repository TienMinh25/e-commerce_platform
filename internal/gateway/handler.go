package gateway

import (
	"github.com/TienMinh25/delivery-system/pkg"
	"go.uber.org/fx"
)

const _USER_ID_KEY = "userID"

type handler struct {
	gatewayService Service
	apiSecretKey   string
}

func NewHandler(
	lifecycle fx.Lifecycle,
	gatewayService Service,
	tracer pkg.DistributedTracer,
	messageQueue pkg.Queue,
) {

}