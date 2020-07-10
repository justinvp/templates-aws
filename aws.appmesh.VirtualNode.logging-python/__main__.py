import pulumi
import pulumi_aws as aws

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
        "logging": {
            "accessLog": {
                "file": {
                    "path": "/dev/stdout",
                },
            },
        },
        "serviceDiscovery": {
            "dns": {
                "hostname": "serviceb.simpleapp.local",
            },
        },
    })

