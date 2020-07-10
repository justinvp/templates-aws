import pulumi
import pulumi_aws as aws

example = aws.glue.Connection("example", connection_properties={
    "JDBC_CONNECTION_URL": "jdbc:mysql://example.com/exampledatabase",
    "PASSWORD": "examplepassword",
    "USERNAME": "exampleusername",
})

