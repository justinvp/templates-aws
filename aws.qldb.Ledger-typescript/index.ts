import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sample_ledger = new aws.qldb.Ledger("sample-ledger", {});

