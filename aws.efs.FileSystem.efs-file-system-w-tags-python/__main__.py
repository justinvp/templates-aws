import pulumi
import pulumi_aws as aws

foo = aws.efs.FileSystem("foo", tags={
    "Name": "MyProduct",
})

