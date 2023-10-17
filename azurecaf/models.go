package azurecaf

import (
	"math/rand"
	"time"
)

const (
	// ConventionCafClassic applies the CAF recommended naming convention
	ConventionCafClassic string = "cafclassic"
	// ConventionCafRandom defines the CAF random naming convention
	ConventionCafRandom string = "cafrandom"
	// ConventionRandom applies a random naming convention based on the max length of the resource
	ConventionRandom string = "random"
	// ConventionPassThrough defines the CAF random naming convention
	ConventionPassThrough string = "passthrough"
)

const (
	alphanum    string = "[^0-9A-Za-z]"
	alphanumh   string = "[^0-9A-Za-z-]"
	alphanumu   string = "[^0-9A-Za-z_]"
	alphanumhu  string = "[^0-9A-Za-z_-]"
	alphanumhup string = "[^0-9A-Za-z_.-]"
	unicode     string = `[^-\w\._\(\)]`
	invappi     string = "[%&\\?/]"     //appinisghts invalid character
	invsqldb    string = "[<>*%&:\\/?]" //sql db invalid character

	//Need to find a way to filter beginning and end of string
	//alphanumstartletter string = "\\A[^a-z][^0-9A-Za-z]"
)

const (
	suffixSeparator string = "-"
)

// ResourceStructure stores the CafPrefix and the MaxLength of an azure resource
type ResourceStructure struct {
	// Resource type name
	ResourceTypeName string `json:"name"`
	// Resource prefix as defined in the Azure Cloud Adoption Framework
	CafPrefix string `json:"slug,omitempty"`
	// MaxLength attribute define the maximum length of the name
	MinLength int `json:"min_length"`
	// MaxLength attribute define the maximum length of the name
	MaxLength int `json:"max_length"`
	// enforce lowercase
	LowerCase bool `json:"lowercase,omitempty"`
	// Regular expression to apply to the resource type
	RegEx string `json:"regex,omitempty"`
	// the Regular expression to validate the generated string
	ValidationRegExp string `json:"validatation_regex,omitempty"`
	// can the resource include dashes
	Dashes bool `json:"dashes"`
}

var (
	alphagenerator = []rune("abcdefghijklmnopqrstuvwxyz")
)

// Generate a random value to add to the resource names
func randSeq(length int, seed *int64) string {
	if length == 0 {
		return ""
	}
	// initialize random seed
	if seed == nil || *seed == 0 {
		value := time.Now().UnixNano()
		seed = &value
	}
	rand.Seed(*seed)
	// generate at least one random character
	b := make([]rune, length)
	for i := range b {
		// We need the random generated string to start with a letter
		b[i] = alphagenerator[rand.Intn(len(alphagenerator)-1)]
	}
	return string(b)
}

// Resources currently supported
var Resources = map[string]ResourceStructure{
	"gke": {"azure kubernetes service", "gke", 1, 63, false, alphanumhu, "^[0-9a-zA-Z][0-9A-Za-z_.-]{0,61}[0-9a-zA-Z]$", true},
	"gen": {"generic", "gen", 1, 24, false, alphanum, "^[0-9a-zA-Z]{1,24}$", true},
}

// ResourcesMapping enforcing new naming convention
var ResourcesMapping = map[string]ResourceStructure{
	"generic":                Resources["gen"],
	"gke_kubernetes_cluster": Resources["gke"],
}
