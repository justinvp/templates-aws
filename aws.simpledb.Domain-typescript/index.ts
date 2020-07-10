import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const users = new aws.simpledb.Domain("users", {});

