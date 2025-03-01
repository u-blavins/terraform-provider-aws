// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package datapipeline

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	datapipeline_sdkv2 "github.com/aws/aws-sdk-go-v2/service/datapipeline"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
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
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourcePipeline,
			TypeName: "aws_datapipeline_pipeline",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  DataSourcePipelineDefinition,
			TypeName: "aws_datapipeline_pipeline_definition",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourcePipeline,
			TypeName: "aws_datapipeline_pipeline",
			Name:     "Pipeline",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Pipeline",
			},
		},
		{
			Factory:  ResourcePipelineDefinition,
			TypeName: "aws_datapipeline_pipeline_definition",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DataPipeline
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*datapipeline_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return datapipeline_sdkv2.NewFromConfig(cfg,
		datapipeline_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
