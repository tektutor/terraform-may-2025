package provider

import (
   "context"
   "log"
   "io"
   "os"

   "github.com/docker/docker/client"
   "github.com/docker/docker/api/types/image"

   "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
   "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDockerImage() *schema.Resource {

	return &schema.Resource {
		Description: "This resource will create, read, update and delete dcoker image resources via Terraform.",

		CreateContext: resourceCreateDockerImage,
		ReadContext  : resourceReadDockerImage,
		UpdateContext: resourceUpdateDockerImage,
		DeleteContext: resourceDeleteDockerImage,

		Schema: map[string]*schema.Schema {
                   "image_name": {
			   Description: "Name of the docker image include tag.",
			   Type: schema.TypeString,
			   Required: true,
		   },
		},
	}
}

func resourceCreateDockerImage(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	
    imageName := d.Get("image_name").(string)

    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    reader, err1 := cli.ImagePull(ctx, imageName, image.PullOptions{})
    if err1 != nil {
        panic(err1)
    }

    io.Copy(os.Stdout,reader)

    imageInspect, _, err2 := cli.ImageInspectWithRaw(ctx, imageName)
    if err2 != nil {
        panic(err2)
    }

    d.SetId( imageInspect.ID )

    return nil
}

func resourceReadDockerImage(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
    return nil
}

func resourceUpdateDockerImage(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
    return nil
}

func resourceDeleteDockerImage(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
    imageName := d.Get("image_name").(string)

    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    _, err = cli.ImageRemove ( ctx, imageName, image.RemoveOptions{ Force: true, } )
    if err != nil {
	log.Printf("Unable to delete docker image: %s", imageName )
        panic(err)
    }

    return nil
}


