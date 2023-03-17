package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type TerraformBackendStackProps struct {
    awscdk.StackProps
    BucketName  string
    TableName   string
}

func NewTerraformBackendStack(scope constructs.Construct, id string, props *TerraformBackendStackProps) awscdk.Stack {
    stack := awscdk.NewStack(scope, &id, &props.StackProps)

    // Create an S3 bucket for the Terraform backend.
    awss3.NewBucket(stack, jsii.String("TerraformBackendBucket"), &awss3.BucketProps{
        BucketName: jsii.String(props.BucketName),
        Versioned:  jsii.Bool(true),
    })

    // Create a DynamoDB table for the Terraform lock.
    awsdynamodb.NewTable(stack, jsii.String("TerraformLockTable"), &awsdynamodb.TableProps{
        TableName:   jsii.String(props.TableName),
        PartitionKey: &awsdynamodb.Attribute{
            Name: jsii.String("LockID"),
            Type: awsdynamodb.AttributeType_STRING,
        },
        BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
    })

    return stack
}

func main() {
    app := awscdk.NewApp(nil)

    accountID := os.Getenv("CDK_AWS_ACCOUNT_ID")
    if accountID == "" {
        fmt.Println("Error: CDK_AWS_ACCOUNT_ID environment variable is not set")
        os.Exit(1)
    }

    region := os.Getenv("CDK_AWS_REGION")
    if region == "" {
        region = "us-west-2"
    }

    bucketName := os.Getenv("CDK_TF_BACKEND_BUCKET_NAME")
    if bucketName == "" {
        bucketName = "default-terraform-backend-bucket"
    }

    tableName := os.Getenv("CDK_TF_LOCK_TABLE_NAME")
    if tableName == "" {
        tableName = "default-terraform-lock-table"
    }

    NewTerraformBackendStack(app, "TerraformBackendStack", &TerraformBackendStackProps{
        StackProps: awscdk.StackProps{
            Env: &awscdk.Environment{
				Account: jsii.String(accountID),
				Region:  jsii.String(region),
			},
        },
        BucketName: bucketName,
        TableName:  tableName,
    })

    app.Synth(nil)
}
