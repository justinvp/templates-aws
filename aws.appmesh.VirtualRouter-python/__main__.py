import pulumi
import pulumi_aws as aws

serviceb = aws.appmesh.VirtualRouter("serviceb",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    spec={
        "listener": {
            "portMapping": {
                "port": 8080,
                "protocol": "http",
            },
        },
    })

