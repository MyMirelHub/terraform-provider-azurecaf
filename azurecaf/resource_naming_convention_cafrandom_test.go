package azurecaf

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCafNamingConventionCaf_Random(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCafRandomConfig,
				Check: resource.ComposeTestCheckFunc(

					testAccCafNamingValidation(
						"azurecaf_naming_convention.gke",
						"kubedemo",
						Resources["gke"].MaxLength,
						"rdmi"),
					regexMatch("azurecaf_naming_convention.gke", regexp.MustCompile(Resources["gke"].ValidationRegExp), 1),
				),
			},
		},
	})
}

const testAccResourceCafRandomConfig = `

# Azure Kubernetes Service
resource "azurecaf_naming_convention" "gke" {
    convention      = "cafrandom"
    name            = "kubedemo"
    prefix          = "rdmi"
    resource_type   = "gke"
}
`
