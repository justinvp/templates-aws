import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const fileSystemId = config.get("fileSystemId") || "";

const byId = pulumi.output(aws.efs.getFileSystem({
    fileSystemId: fileSystemId,
}, { async: true }));

