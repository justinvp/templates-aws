import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceb = new aws.appmesh.Route("serviceb", {
    meshName: aws_appmesh_mesh_simple.id,
    spec: {
        httpRoute: {
            action: {
                weightedTargets: [{
                    virtualNode: aws_appmesh_virtual_node_serviceb.name,
                    weight: 100,
                }],
            },
            match: {
                headers: [{
                    match: {
                        prefix: "123",
                    },
                    name: "clientRequestId",
                }],
                method: "POST",
                prefix: "/",
                scheme: "https",
            },
        },
    },
    virtualRouterName: aws_appmesh_virtual_router_serviceb.name,
});

