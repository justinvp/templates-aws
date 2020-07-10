import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceb1 = new aws.appmesh.VirtualNode("serviceb1", {
    meshName: aws_appmesh_mesh_simple.id,
    spec: {
        backends: [{
            virtualService: {
                virtualServiceName: "servicea.simpleapp.local",
            },
        }],
        listener: {
            portMapping: {
                port: 8080,
                protocol: "http",
            },
        },
        serviceDiscovery: {
            dns: {
                hostname: "serviceb.simpleapp.local",
            },
        },
    },
});

