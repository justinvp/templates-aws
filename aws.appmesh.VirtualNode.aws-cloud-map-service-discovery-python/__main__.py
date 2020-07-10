import pulumi
import pulumi_aws as aws

example = aws.servicediscovery.HttpNamespace("example")
serviceb1 = aws.appmesh.VirtualNode("serviceb1",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    spec={
        "backend": [{
            "virtualService": {
                "virtualServiceName": "servicea.simpleapp.local",
            },
        }],
        "listener": {
            "portMapping": {
                "port": 8080,
                "protocol": "http",
            },
        },
        "serviceDiscovery": {
            "awsCloudMap": {
                "attributes": {
                    "stack": "blue",
                },
                "namespaceName": example.name,
                "service_name": "serviceb1",
            },
        },
    })

