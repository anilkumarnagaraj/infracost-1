package ibm

import (
	"github.com/infracost/infracost/internal/schema"
	"github.com/tidwall/gjson"
)

var DefaultProviderRegion = "us-south"

func GetDefaultRefIDFunc(d *schema.ResourceData) []string {

	defaultRefs := []string{d.Get("id").String()}

	if d.Get("self_link").Exists() {
		defaultRefs = append(defaultRefs, d.Get("self_link").String())
	}

	return defaultRefs
}

func GetSpecialContext(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{}
}

func GetResourceRegion(resourceType string, v gjson.Result) string {
	return ""
}

func ParseTags(resourceType string, v gjson.Result) map[string]string {
	tags := make(map[string]string)
	for k, v := range v.Get("labels").Map() {
		tags[k] = v.String()
	}
	return tags
}
