package azurecaf

import (
	"context"
	"reflect"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func setData(prefixes []string, name string, suffixes []string, cleanInput bool) *schema.ResourceData {
	data := &schema.ResourceData{}
	data.Set("name", name)
	data.Set("prefixes", prefixes)
	data.Set("suffixes", suffixes)
	data.Set("clean_input", cleanInput)
	return data
}

func TestConcatenateParameters_azurerm_public_ip_prefix(t *testing.T) {
	prefixes := []string{"pre"}
	suffixes := []string{"suf"}
	content := []string{"name", "ip"}
	separator := "-"
	expected := "pre-name-ip-suf"
	result := concatenateParameters(separator, prefixes, content, suffixes)
	if result != expected {
		t.Errorf("Expected %s but received %s", expected, result)
	}
}

func TestGetSlug_unknown(t *testing.T) {
	resourceType := "azurerm_does_not_exist"
	convention := ConventionCafClassic
	result := getSlug(resourceType, convention)
	expected := ""
	if result != expected {
		t.Errorf("Expected %s but received %s", expected, result)
	}
}

func TestAccResourceName_CafClassic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{

		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceNameCafClassicConfig,
				Check: resource.ComposeTestCheckFunc(

					testAccCafNamingValidation(
						"azurecaf_name.passthrough",
						"passthrough",
						11,
						""),
					regexMatch("azurecaf_name.passthrough", regexp.MustCompile(ResourceDefinitions["gke_kubernetes_cluster"].ValidationRegExp), 1),
				),
			},
		},
	})
}

func TestComposeName(t *testing.T) {
	namePrecedence := []string{"name", "random", "slug", "suffixes", "prefixes"}
	prefixes := []string{"a", "b"}
	suffixes := []string{"c", "d"}
	name := composeName("-", prefixes, "name", "slug", suffixes, "rd", 21, namePrecedence)
	expected := "a-b-slug-name-rd-c-d"
	if name != expected {
		t.Logf("Fail to generate name expected %s received %s", expected, name)
		t.Fail()
	}
}

func TestComposeNameCutCorrect(t *testing.T) {
	namePrecedence := []string{"name", "slug", "random", "suffixes", "prefixes"}
	prefixes := []string{"a", "b"}
	suffixes := []string{"c", "d"}
	name := composeName("-", prefixes, "name", "slug", suffixes, "rd", 19, namePrecedence)
	expected := "b-slug-name-rd-c-d"
	if name != expected {
		t.Logf("Fail to generate name expected %s received %s", expected, name)
		t.Fail()
	}
}

func TestComposeNameCutMaxLength(t *testing.T) {
	namePrecedence := []string{"name", "slug", "random", "suffixes", "prefixes"}
	prefixes := []string{}
	suffixes := []string{}
	name := composeName("-", prefixes, "aaaaaaaaaa", "bla", suffixes, "", 10, namePrecedence)
	expected := "aaaaaaaaaa"
	if name != expected {
		t.Logf("Fail to generate name expected %s received %s", expected, name)
		t.Fail()
	}
}

func TestComposeNameCutCorrectSuffixes(t *testing.T) {
	namePrecedence := []string{"name", "slug", "random", "suffixes", "prefixes"}
	prefixes := []string{"a", "b"}
	suffixes := []string{"c", "d"}
	name := composeName("-", prefixes, "name", "slug", suffixes, "rd", 15, namePrecedence)
	expected := "slug-name-rd-c"
	if name != expected {
		t.Logf("Fail to generate name expected %s received %s", expected, name)
		t.Fail()
	}
}

func TestComposeEmptyStringArray(t *testing.T) {
	namePrecedence := []string{"name", "slug", "random", "suffixes", "prefixes"}
	prefixes := []string{"", "b"}
	suffixes := []string{"", "d"}
	name := composeName("-", prefixes, "", "", suffixes, "", 15, namePrecedence)
	expected := "b-d"
	if name != expected {
		t.Logf("Fail to generate name expected %s received %s", expected, name)
		t.Fail()
	}
}

func TestGetResourceNameInvalidResourceType(t *testing.T) {
	namePrecedence := []string{"name", "slug", "random", "suffixes", "prefixes"}
	resourceName, err := getResourceName("azurerm_invalid", "-", []string{"a", "b"}, "myrg", nil, "1234", "cafclassic", true, false, true, namePrecedence)
	expected := "a-b-rg-myrg-1234"

	if err == nil {
		t.Logf("Expected a validation error, got nil")
		t.Fail()
	}
	if expected == resourceName {
		t.Logf("valid name received while an error is expected")
		t.Fail()
	}
}

func testResourceNameStateDataV2() map[string]interface{} {
	return map[string]interface{}{}
}

func testResourceNameStateDataV3() map[string]interface{} {
	return map[string]interface{}{
		"use_slug": true,
	}
}

func TestResourceExampleInstanceStateUpgradeV2(t *testing.T) {
	expected := testResourceNameStateDataV3()
	actual, err := resourceNameStateUpgradeV2(context.Background(), testResourceNameStateDataV2(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}

const testAccResourceNameCafClassicConfig = `


resource "azurecaf_name" "passthrough" {
    name            = "passthrough"
	resource_type   = "gke_kubernetes_cluster"
	prefixes        = ["pr1", "pr2"]
	suffixes        = ["su1", "su2"]
	random_seed     = 1
	random_length   = 5
	clean_input     = true
	passthrough     = true
}

`
