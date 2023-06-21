// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package rolesanywhere

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	rolesanywhere_sdkv2 "github.com/aws/aws-sdk-go-v2/service/rolesanywhere"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceProfile,
			TypeName: "aws_rolesanywhere_profile",
			Name:     "Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceTrustAnchor,
			TypeName: "aws_rolesanywhere_trust_anchor",
			Name:     "Trust Anchor",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.RolesAnywhere
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*rolesanywhere_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return rolesanywhere_sdkv2.NewFromConfig(cfg, func(o *rolesanywhere_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.EndpointResolver = rolesanywhere_sdkv2.EndpointResolverFromURL(endpoint)
		}
	}), nil
}

var ServicePackage = &servicePackage{}
