import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceb = new aws.appmesh.VirtualRouter("serviceb", {
    meshName: aws_appmesh_mesh_simple.id,
    spec: {
        listener: {
            portMapping: {
                port: 8080,
                protocol: "http",
            },
        },
    },
});

