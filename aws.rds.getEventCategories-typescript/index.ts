import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleEventCategories = pulumi.output(aws.rds.getEventCategories({ async: true }));

export const example = exampleEventCategories.eventCategories;

