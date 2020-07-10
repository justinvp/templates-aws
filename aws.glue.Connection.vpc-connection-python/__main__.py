import pulumi
import pulumi_aws as aws

example = aws.glue.Connection("example",
    connection_properties={
        "JDBC_CONNECTION_URL": f"jdbc:mysql://{aws_rds_cluster['example']['endpoint']}/exampledatabase",
        "PASSWORD": "examplepassword",
        "USERNAME": "exampleusername",
    },
    physical_connection_requirements={
        "availability_zone": aws_subnet["example"]["availability_zone"],
        "securityGroupIdList": [aws_security_group["example"]["id"]],
        "subnet_id": aws_subnet["example"]["id"],
    })

