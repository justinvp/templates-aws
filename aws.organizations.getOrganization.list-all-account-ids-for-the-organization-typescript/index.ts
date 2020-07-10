import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.organizations.getOrganization({});
export const accountIds = example.then(example => example.accounts.map(__item => __item.id));

