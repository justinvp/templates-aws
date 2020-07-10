import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.fsx.WindowsFileSystem("example", {
    kmsKeyId: aws_kms_key_example.arn,
    selfManagedActiveDirectory: {
        dnsIps: [
            "10.0.0.111",
            "10.0.0.222",
        ],
        domainName: "corp.example.com",
        password: "avoid-plaintext-passwords",
        username: "Admin",
    },
    storageCapacity: 300,
    subnetIds: aws_subnet_example.id,
    throughputCapacity: 1024,
});

