package bazel

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceBazelTarget() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBazelTargetRead,
		Schema: map[string]*schema.Schema{
			"query": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stdout": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceBazelTargetRead(d *schema.ResourceData, meta interface{}) error {
	var bazelQuery string

	if v, ok := d.GetOk("query"); ok {
		bazelQuery = v.(string)
		if len(bazelQuery) == 0 {
			return nil
		}
	}

	output, err := runBazelQuery(bazelQuery)
	if err != nil {
		return fmt.Errorf("error running bazel query: %v", err)
	}

	err = d.Set("stdout", output)
	if err != nil {
		return errors.New("error setting data output")
	}
	d.SetId(bazelQuery)

	return nil
}

func runBazelQuery(query string) (string, error) {
	var outb, errb bytes.Buffer

	cmd := exec.Command("bazel", "query", query)
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%s, err: %s", errb.String(), err)
	}
	return outb.String(), nil
}
