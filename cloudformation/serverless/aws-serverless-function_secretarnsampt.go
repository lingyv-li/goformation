// Code generated by "go generate". Please don't change this file directly.

package serverless

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Function_SecretArnSAMPT AWS CloudFormation Resource (AWS::Serverless::Function.SecretArnSAMPT)
// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
type Function_SecretArnSAMPT struct {

	// SecretArn AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SecretArn string `json:"SecretArn"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Function_SecretArnSAMPT) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.SecretArnSAMPT"
}
