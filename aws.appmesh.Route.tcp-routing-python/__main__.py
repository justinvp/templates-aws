import pulumi
import pulumi_aws as aws

serviceb = aws.appmesh.Route("serviceb",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    spec={
        "tcpRoute": {
            "action": {
                "weightedTarget": [{
                    "virtualNode": aws_appmesh_virtual_node["serviceb1"]["name"],
                    "weight": 100,
                }],
            },
        },
    },
    virtual_router_name=aws_appmesh_virtual_router["serviceb"]["name"])

