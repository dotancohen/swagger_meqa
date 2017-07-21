// This package handles swagger.json file parsing
package mqswag

import (
	"log"
	"meqa/mqutil"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/loads/fmts"
	"github.com/go-openapi/spec"
)

// string fields in OpenAPI doc
const SWAGGER = "swagger"
const SCHEMES = "schemes"
const CONSUMES = "consumes"
const PRODUCES = "produces"
const INFO = "info"
const HOST = "host"
const BASEPATH = "basePath"
const TAGS = "tags"
const PATHS = "paths"
const SECURITY_DEFINITIONS = "securityDefinitions"
const DEFINITIONS = "definitions"

// Init from a file
func CreateSwaggerFromURL(path string) (*spec.Swagger, error) {
	specDoc, err := loads.Spec(path)
	if err != nil {
		mqutil.Logger.Printf("Can't open the following file: %s", path)
		mqutil.Logger.Println(err.Error())
		return nil, err
	}

	log.Println("Would be serving:", specDoc.Spec().Info.Title)

	return specDoc.Spec(), nil
}

func init() {
	loads.AddLoader(fmts.YAMLMatcher, fmts.YAMLDoc)
}