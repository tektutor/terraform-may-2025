package provider

import (
	"context"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDockerContainer() *schema.Resource {

   return &schema.Resource {
      Description: "This resource will support Create, Read, Update and Delete docker container resources via Terraform.",

      CreateContext: resourceCreateDockerContainer,
      ReadContext:   resourceReadDockerContainer,
      UpdateContext: resourceUpdateDockerContainer,
      DeleteContext: resourceDeleteDockerContainer,

      //mention the parameters our file resource requires
      Schema: map[string]*schema.Schema {

	      "container_name": {
		  Description: "Name of the container",
		  Type       : schema.TypeString,
		  Required   : true,
	      },
	      "host_name": {
		      Description: "Hostname of the container.",
		      Type       : schema.TypeString,
		      Required   : true,
	      },
	      "image_name": {
		      Description: "Name of the docker image.",
		      Type       : schema.TypeString,
		      Required   : true,
	      },
      },
   }
}

func resourceCreateDockerContainer( ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

    //Retrieve the inputs user provider in the terraform .tf file
    imageName     := d.Get("image_name").(string)
    containerName := d.Get("container_name").(string)
    containerHostname := d.Get("host_name").(string)

    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image: imageName,
        Hostname: containerHostname,
    }, nil, nil, nil, containerName)
    if err != nil {
        panic(err)
    }

    if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
        panic(err)
    }

    d.Set("container_name", containerName)
    d.Set("host_name", containerHostname )
    d.SetId(resp.ID)

    return nil
}

func resourceReadDockerContainer( ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func resourceUpdateDockerContainer( ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
    containerName := d.Get("container_name").(string)
    containerID := d.Id()

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

    return nil
}

func resourceDeleteDockerContainer( ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
    containerName := d.Get("container_name").(string)

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

    return nil
}
