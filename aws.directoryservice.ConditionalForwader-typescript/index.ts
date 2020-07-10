import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.directoryservice.ConditionalForwader("example", {
    directoryId: aws_directory_service_directory_ad.id,
    dnsIps: [
        "8.8.8.8",
        "8.8.4.4",
    ],
    remoteDomainName: "example.com",
});

