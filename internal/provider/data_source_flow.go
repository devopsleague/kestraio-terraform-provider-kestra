package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFlow() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to access information about an existing Kestra Flow",

		ReadContext: dataSourceFlowRead,
		Schema: map[string]*schema.Schema{
			"tenant_id": {
				Description: "The tenant id.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"namespace": {
				Description: "The namespace.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"flow_id": {
				Description: "The flow id.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"revision": {
				Description: "The flow revision.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"content": {
				Description: "The flow content as yaml.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceFlowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*Client)
	var diags diag.Diagnostics

	namespaceId := d.Get("namespace").(string)
	flowId := d.Get("flow_id").(string)
	tenantId := c.TenantId

	r, reqErr := c.request("GET", fmt.Sprintf("%s/flows/%s/%s", apiRoot(tenantId), namespaceId, flowId), nil)
	if reqErr != nil {
		return diag.FromErr(reqErr.Err)
	}

	errs := flowApiToSchema(r.(map[string]interface{}), d, c)
	if errs != nil {
		return errs
	}

	return diags
}
