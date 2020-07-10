import pulumi
import pulumi_aws as aws

servicea = aws.appmesh.VirtualService("servicea",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    spec={
        "provider": {
            "virtualNode": {
                "virtualNodeName": aws_appmesh_virtual_node["serviceb1"]["name"],
            },
        },
    })

