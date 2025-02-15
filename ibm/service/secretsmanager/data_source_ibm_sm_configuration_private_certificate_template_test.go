// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmConfigurationPrivateCertificateTemplateDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmConfigurationPrivateCertificateTemplateDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_private_certificate_template.sm_configuration_private_certificate_template", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_private_certificate_template.sm_configuration_private_certificate_template", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_private_certificate_template.sm_configuration_private_certificate_template", "config_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_private_certificate_template.sm_configuration_private_certificate_template", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_private_certificate_template.sm_configuration_private_certificate_template", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_private_certificate_template.sm_configuration_private_certificate_template", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_private_certificate_template.sm_configuration_private_certificate_template", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_private_certificate_template.sm_configuration_private_certificate_template", "certificate_authority"),
				),
			},
		},
	})
}

func testAccCheckIbmSmConfigurationPrivateCertificateTemplateDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_configuration_private_certificate_root_ca" "ibm_sm_configuration_private_certificate_root_ca_instance" {
			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			crl_expiry = "10000h"
			name = "root-ca-terraform-private-cert-datasource-test"
		}
		resource "ibm_sm_configuration_private_certificate_intermediate_ca" "ibm_sm_configuration_private_certificate_intermediate_ca_instance" {
  			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			issuer = ibm_sm_configuration_private_certificate_root_ca.ibm_sm_configuration_private_certificate_root_ca_instance.name
			signing_method = "internal"
			name = "intermediate-ca-terraform-private-cert-datasource-test"
		}
		resource "ibm_sm_configuration_private_certificate_template" "ibm_sm_configuration_private_certificate_template_instance" {
			instance_id   = "%s"
			region        = "%s"
			certificate_authority = ibm_sm_configuration_private_certificate_intermediate_ca.ibm_sm_configuration_private_certificate_intermediate_ca_instance.name
			allow_any_name = true
			name = "template-terraform-private-cert-datasource-test"
		}

		data "ibm_sm_configuration_private_certificate_template" "sm_configuration_private_certificate_template" {
			instance_id   = "%s"
			region        = "%s"
			name = ibm_sm_configuration_private_certificate_template.ibm_sm_configuration_private_certificate_template_instance.name
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
