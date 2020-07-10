import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testStream = new aws.kinesis.Stream("test_stream", {
    retentionPeriod: 48,
    shardCount: 1,
    shardLevelMetrics: [
        "IncomingBytes",
        "OutgoingBytes",
    ],
    tags: {
        Environment: "test",
    },
});

