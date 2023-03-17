package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestTerraformBackendStack(t *testing.T) {
	os.Setenv("CDK_AWS_ACCOUNT_ID", "123456789012")
	os.Setenv("CDK_AWS_REGION", "us-west-2")
	os.Setenv("CDK_TF_BACKEND_BUCKET_NAME", "test-terraform-backend-bucket")
	os.Setenv("CDK_TF_LOCK_TABLE_NAME", "test-terraform-lock-table")

	app := awscdk.NewApp(nil)

	stack := NewTerraformBackendStack(app, "TestTerraformBackendStack", &TerraformBackendStackProps{
		StackProps: awscdk.StackProps{
			Env: &awscdk.Environment{
				Account: jsii.String(os.Getenv("CDK_AWS_ACCOUNT_ID")),
				Region:  jsii.String(os.Getenv("CDK_AWS_REGION")),
			},
		},
		BucketName: os.Getenv("CDK_TF_BACKEND_BUCKET_NAME"),
		TableName:  os.Getenv("CDK_TF_LOCK_TABLE_NAME"),
	})

	template := assertions.Template_FromStack(stack, nil)
	jsonStr, _ := json.Marshal(template.ToJSON())
	t.Log(string(jsonStr))

	template.HasResourceProperties(jsii.String("AWS::S3::Bucket"), map[string]interface{}{
		"BucketName": "test-terraform-backend-bucket",
		"VersioningConfiguration": map[string]interface{}{
			"Status": "Enabled",
		},
	})

	template.HasResourceProperties(jsii.String("AWS::DynamoDB::Table"), map[string]interface{}{
		"TableName":   "test-terraform-lock-table",
		"BillingMode": "PAY_PER_REQUEST",
		"AttributeDefinitions": []map[string]interface{}{
			{
				"AttributeName": "LockID",
				"AttributeType": "S",
			},
		},
	})
}
