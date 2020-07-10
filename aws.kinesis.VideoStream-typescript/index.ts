import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultVideoStream = new aws.kinesis.VideoStream("default", {
    dataRetentionInHours: 1,
    deviceName: "kinesis-video-device-name",
    mediaType: "video/h264",
    tags: {
        Name: "kinesis-video-stream",
    },
});

