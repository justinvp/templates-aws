import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Connection("example", {
    connectionProperties: {
        JDBC_CONNECTION_URL: pulumi.interpolate`jdbc:mysql://${aws_rds_cluster_example.endpoint}/exampledatabase`,
        PASSWORD: "examplepassword",
        USERNAME: "exampleusername",
    },
    physicalConnectionRequirements: {
        availabilityZone: aws_subnet_example.availabilityZone,
        securityGroupIdLists: [aws_security_group_example.id],
        subnetId: aws_subnet_example.id,
    },
});

