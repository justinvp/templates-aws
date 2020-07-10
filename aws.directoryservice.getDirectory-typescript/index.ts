import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws_directory_service_directory_main.id.apply(id => aws.directoryservice.getDirectory({
    directoryId: id,
}, { async: true }));

