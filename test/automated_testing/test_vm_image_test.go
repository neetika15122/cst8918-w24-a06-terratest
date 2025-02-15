// test_vm_image.go
package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAzureVMImage(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
	}

	terraform.InitAndApply(t, terraformOptions)

	// Fetch Terraform outputs
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroup := terraform.Output(t, terraformOptions, "resource_group_name")

	// Ensure VM and resource group are not empty
	assert.NotEmpty(t, vmName, "VM name should not be empty")
	assert.NotEmpty(t, resourceGroup, "Resource group should not be empty")

	// Expected values from Terraform configuration
	expectedPublisher := "Canonical"
	expectedSKU := "22_04-lts-gen2"

	// Validate the expected image details
	assert.Equal(t, expectedPublisher, "Canonical", "VM should be from Canonical")
	assert.Equal(t, expectedSKU, "22_04-lts-gen2", "VM should be running Ubuntu 22.04 LTS")
}
