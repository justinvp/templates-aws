import pulumi
import pulumi_aws as aws

serviceb = aws.appmesh.Route("serviceb",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    spec={
        "httpRoute": {
            "action": {
                "weightedTarget": [{
                    "virtualNode": aws_appmesh_virtual_node["serviceb"]["name"],
                    "weight": 100,
                }],
            },
            "match": {
                "header": [{
                    "match": {
                        "prefix": "123",
                    },
                    "name": "clientRequestId",
                }],
                "method": "POST",
                "prefix": "/",
                "scheme": "https",
            },
        },
    },
    virtual_router_name=aws_appmesh_virtual_router["serviceb"]["name"])

