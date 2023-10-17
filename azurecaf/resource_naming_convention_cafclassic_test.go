package azurecaf

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCafNamingConvention_Classic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCafClassicConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCafNamingValidation(
						"azurecaf_naming_convention.classic_rg",
						"myrg",
						7,
						"rg"),
					regexMatch("azurecaf_naming_convention.classic_rg", regexp.MustCompile(Resources["rg"].ValidationRegExp), 1),
					testAccCafNamingValidation(
						"azurecaf_naming_convention.classic_aks",
						"kubedemo",
						12,
						"gke"),
					regexMatch("azurecaf_naming_convention.classic_aks", regexp.MustCompile(Resources["gke"].ValidationRegExp), 1),
				),
			},
		},
	})
}

const testAccResourceCafClassicConfig = `

# Resource Group
resource "azurecaf_naming_convention" "classic_rg" {
    convention      = "cafclassic"
    name            = "myrg"
    resource_type   = "rg"
}

# Google Kubernetes Service
resource "azurecaf_naming_convention" "classic_aks" {
    convention      = "cafclassic"
    name            = "kubedemo"
    resource_type   = "gke"
}
`
