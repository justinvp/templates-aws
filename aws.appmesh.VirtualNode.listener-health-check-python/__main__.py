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
            "health_check": {
                "healthyThreshold": 2,
                "intervalMillis": 5000,
                "path": "/ping",
                "protocol": "http",
                "timeoutMillis": 2000,
                "unhealthyThreshold": 2,
            },
            "portMapping": {
                "port": 8080,
                "protocol": "http",
            },
        },
        "serviceDiscovery": {
            "dns": {
                "hostname": "serviceb.simpleapp.local",
            },
        },
    })

