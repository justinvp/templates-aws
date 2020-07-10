import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const servicea = new aws.appmesh.VirtualService("servicea", {
    meshName: aws_appmesh_mesh_simple.id,
    spec: {
        provider: {
            virtualRouter: {
                virtualRouterName: aws_appmesh_virtual_router_serviceb.name,
            },
        },
    },
});

