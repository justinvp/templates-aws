import pulumi
import pulumi_aws as aws

simple = aws.appmesh.Mesh("simple", spec={
    "egressFilter": {
        "type": "ALLOW_ALL",
    },
})

