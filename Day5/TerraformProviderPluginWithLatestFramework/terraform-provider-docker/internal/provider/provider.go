package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
	_ provider.ProviderWithFunctions = (*dockerProvider)(nil)
)

func New() provider.Provider {
	return &dockerProvider{}
}

type dockerProvider struct{}

func (p *dockerProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "docker"
}

func (p *dockerProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *dockerProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
//		DockerContainerDataSource,
	}
}

func (p *dockerProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		DockerContainerResource,
	}
}

func (p *dockerProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func (p *dockerProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}
