package azurecaf

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCafNamingConvention_Passthrough(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePassthroughConfig,
				Check: resource.ComposeTestCheckFunc(

					testAccCafNamingValidation(
						"azurecaf_naming_convention.passthrough_aks",
						"kubedemo",
						8,
						"kube"),
					regexMatch("azurecaf_naming_convention.passthrough_aks", regexp.MustCompile(Resources["gke"].ValidationRegExp), 1),
				),
			},
		},
	})
}

const testAccResourcePassthroughConfig = `

# Azure Kubernetes Services
resource "azurecaf_naming_convention" "passthrough_aks" {
    convention      = "passthrough"
    name            = "kubedemo"
    resource_type   = "gke_kubernetes_cluster"
}
`
