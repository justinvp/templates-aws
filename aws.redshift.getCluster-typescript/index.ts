import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testCluster = pulumi.output(aws.redshift.getCluster({
    clusterIdentifier: "test-cluster",
}, { async: true }));
const testStream = new aws.kinesis.FirehoseDeliveryStream("test_stream", {
    destination: "redshift",
    redshiftConfiguration: {
        clusterJdbcurl: pulumi.interpolate`jdbc:redshift://${testCluster.endpoint}/${testCluster.databaseName}`,
        copyOptions: "delimiter '|'", // the default delimiter
        dataTableColumns: "test-col",
        dataTableName: "test-table",
        password: "T3stPass",
        roleArn: aws_iam_role_firehose_role.arn,
        username: "testuser",
    },
    s3Configuration: {
        bucketArn: aws_s3_bucket_bucket.arn,
        bufferInterval: 400,
        bufferSize: 10,
        compressionFormat: "GZIP",
        roleArn: aws_iam_role_firehose_role.arn,
    },
});

