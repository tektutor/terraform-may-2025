package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)


var (
	_ resource.Resource = (*dockerContainerResource)(nil)
)

func DockerContainerResource() resource.Resource {
	return &dockerContainerResource{}
}

type dockerContainerResource struct{}

func (n *dockerContainerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "This terraform provides supports Create, Read, Update and Delete operations on Docker containers.",
		Attributes: map[string]schema.Attribute{

			"container_name": schema.StringAttribute{
				Description: "Docker Container name",
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"container_hostname": schema.StringAttribute{
				Description: "Docker Container hostname",
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"container_image": schema.StringAttribute{
				Description: "Docker Image name",
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				Description: "Docker container ID.",
				Computed:    true,
			},
		},
	}
}

func (n *dockerContainerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_container"
}

func (n *dockerContainerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var plan dockerContainerResourceModelV0

    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
	return
    }

    //Retrieve the inputs user provider in the terraform .tf file
    containerImageName := plan.ContainerImageName.ValueString()
    containerName      := plan.ContainerName.ValueString()
    containerHostName  := plan.ContainerHostName.ValueString()

    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    dockerClientResp, err1 := cli.ContainerCreate(ctx, &container.Config{
        Image: containerImageName,
        Hostname: containerHostName,
    }, nil, nil, nil, containerName)
    if err1 != nil {
        panic(err1)
    }

    if err2 := cli.ContainerStart(ctx, dockerClientResp.ID, container.StartOptions{}); err != nil {
        panic(err2)
    }

    plan.ID = types.StringValue(dockerClientResp.ID)
    plan.ContainerName = types.StringValue(containerName)
    plan.ContainerHostName = types.StringValue(containerHostName)
    plan.ContainerImageName =  types.StringValue(containerImageName)

    diags = resp.State.Set(ctx, &plan)
    resp.Diagnostics.Append(diags...)
}

func (n *dockerContainerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state dockerContainerResourceModelV0

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (n *dockerContainerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
    var plan dockerContainerResourceModelV0

    resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

    if resp.Diagnostics.HasError() {
	return
    }

    //Retrieve the inputs user provider in the terraform .tf file
    containerID       := plan.ID.ValueString()
    containerName     := plan.ContainerName.ValueString()

    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    err = cli.ContainerRename ( ctx, containerID, containerName )

    if err != nil {
	log.Printf("Unable to update container : %s", err )
        panic(err)
    }

    plan.ID = types.StringValue(containerID) 
    plan.ContainerName = types.StringValue(containerName)

    resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (n *dockerContainerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    var containerName string 
    req.State.GetAttribute(ctx, path.Root("container_name"), &containerName)

    //Retrieve the inputs user provider in the terraform .tf file
    //containerName     := plan.ContainerName.ValueString()

    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    err = cli.ContainerStop( ctx, containerName, container.StopOptions{} )

    if err != nil {
	log.Printf("Unable to stop container : %s", err )
        panic(err)
    }

    err = cli.ContainerRemove( ctx, containerName, container.RemoveOptions{ RemoveVolumes: true, Force: true } )

    if err != nil {
	log.Printf("Unable to stop container : %s", err )
        panic(err)
    }
}

type dockerContainerResourceModelV0 struct {
	ContainerName       types.String                   `tfsdk:"container_name"`
	ContainerHostName   types.String                   `tfsdk:"container_hostname"`
	ContainerImageName  types.String                   `tfsdk:"container_image"`
	ID                  types.String                   `tfsdk:"id"`
}
