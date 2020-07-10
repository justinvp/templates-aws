import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceb = new aws.appmesh.Route("serviceb", {
    meshName: aws_appmesh_mesh_simple.id,
    spec: {
        httpRoute: {
            action: {
                weightedTargets: [
                    {
                        virtualNode: aws_appmesh_virtual_node_serviceb1.name,
                        weight: 90,
                    },
                    {
                        virtualNode: aws_appmesh_virtual_node_serviceb2.name,
                        weight: 10,
                    },
                ],
            },
            match: {
                prefix: "/",
            },
        },
    },
    virtualRouterName: aws_appmesh_virtual_router_serviceb.name,
});

