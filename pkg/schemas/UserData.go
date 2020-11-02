package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func UserData() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":    StringRequired(),
			"type":    StringOptionalComputed(),
			"content": StringRequired(),
		},
	}
}
