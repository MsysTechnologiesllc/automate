package v1

import (
	"context"

	uuid "github.com/chef/automate/lib/uuid4"
	chef "github.com/chef/go-chef"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/chef/automate/components/infra-proxy-service/service"
)

// ChefConfig is a V1 infra-proxy server
type ChefConfig struct {
	Name    string
	Key     string
	SkipSSL bool
	BaseURL string
}

// NewChefClient is a V1 infra-proxy server
func NewChefClient(config *ChefConfig) (*chef.Client, error) {

	// build a client
	client, err := chef.NewClient(&chef.Config{
		Name:    config.Name,
		Key:     config.Key,
		SkipSSL: config.SkipSSL,
		BaseURL: config.BaseURL,
	})

	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}

	return client, nil
}

func (s *Server) createClient(ctx context.Context, orgID string) (*chef.Client, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	UUID, err := uuid.FromString(orgID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid org id: %s", err.Error())
	}

	org, err := s.service.Storage.GetOrg(ctx, UUID)
	if err != nil {
		return nil, service.ParseStorageError(err, orgID, "org")
	}

	ServerID, err := uuid.FromString(org.ServerId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid server: %s", err.Error())
	}

	server, err := s.service.Storage.GetServer(ctx, ServerID)
	if err != nil {
		return nil, service.ParseStorageError(err, ServerID, "org")
	}

	baseURL := server.IpAddress + "/organizations/" + org.Name + "/"

	client, err := NewChefClient(&ChefConfig{
		Name:    org.AdminUser,
		Key:     org.AdminKey,
		SkipSSL: true,
		BaseURL: baseURL,
	})

	return client, nil
}
