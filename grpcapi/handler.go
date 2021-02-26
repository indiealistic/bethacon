package grpcapi

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"

	"github.com/indiealistic/bethacon/beaconchain"
	grpcapiproto "github.com/indiealistic/bethacon/proto/bethacon"
	"github.com/indiealistic/bethacon/proto/health"
)

// To make sure Handler implements grpcapiproto.BethaConServiceHandler interface.
var _ grpcapiproto.BethaConServiceServer = &handler{}

// Options serves as the dependency injection container to create a new handler.
type Options struct {
	BeaconChainClient beaconchain.BeaconChain
	Log               *zap.Logger
}

// handler implements grpcapiproto.BethaConServiceHandler interface
type handler struct {
	grpcapiproto.UnimplementedBethaConServiceServer

	beaconChainClient beaconchain.BeaconChain
	log               *zap.Logger
}

// New is the constructor of handler
func New(opts Options) grpcapiproto.BethaConServiceServer {
	return &handler{
		beaconChainClient: opts.BeaconChainClient,
		log:               opts.Log,
	}
}

func (h *handler) Health(context.Context, *empty.Empty) (*health.HealthResponse, error) {
	// TODO: Implement healthcheck
	return &health.HealthResponse{}, nil
}

func (h *handler) Ping(context.Context, *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
