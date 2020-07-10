import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccelerator = new aws.globalaccelerator.Accelerator("example", {
    attributes: {
        flowLogsEnabled: true,
        flowLogsS3Bucket: "example-bucket",
        flowLogsS3Prefix: "flow-logs/",
    },
    enabled: true,
    ipAddressType: "IPV4",
});
const exampleListener = new aws.globalaccelerator.Listener("example", {
    acceleratorArn: exampleAccelerator.id,
    clientAffinity: "SOURCE_IP",
    portRanges: [{
        fromPort: 80,
        toPort: 80,
    }],
    protocol: "TCP",
});

