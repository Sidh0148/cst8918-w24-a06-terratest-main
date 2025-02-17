package test

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestTerraform(t *testing.T) {
    // Define the Terraform options. Point to the directory containing the Terraform configuration.
    terraformOptions := &terraform.Options{
        TerraformDir: "../terraform", // Adjust this path as necessary
    }

    // Ensure that Terraform is cleaned up at the end of the test
    defer terraform.Destroy(t, terraformOptions)

    // Initialize and apply Terraform, checking that no errors occur
    initAndApply := terraform.InitAndApply(t, terraformOptions)

    // Check that apply was successful
    assert.Equal(t, initAndApply, 0, "Expected Terraform apply to be successful")
}
