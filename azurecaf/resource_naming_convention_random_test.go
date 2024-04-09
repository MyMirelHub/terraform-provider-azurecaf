package azurecaf

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCafNamingConventionFull_Random(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRandomConfig,
				Check: resource.ComposeTestCheckFunc(

					testAccCafNamingValidation(
						"azurecaf_naming_convention.random_aks",
						"",
						Resources["gke"].MaxLength,
						"utest"),
					regexMatch("azurecaf_naming_convention.random_aks", regexp.MustCompile(Resources["gke"].ValidationRegExp), 1),
				),
			},
		},
	})
}

const testAccResourceRandomConfig = `

# Azure Kubernetes Service
resource "azurecaf_naming_convention" "random_aks" {
    convention      = "random"
    name            = "TEST-DEV-AKS-RG"
    prefix          = "utest"
    resource_type   = "gke_kubernetes_cluster"
}
`
