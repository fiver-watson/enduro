// Code generated by goa v3.4.3, DO NOT EDIT.
//
// batch HTTP client CLI support package
//
// Command:
// $ goa-v3.4.3 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package client

import (
	"encoding/json"
	"fmt"

	batch "github.com/artefactual-labs/enduro/internal/api/gen/batch"
)

// BuildSubmitPayload builds the payload for the batch submit endpoint from CLI
// flags.
func BuildSubmitPayload(batchSubmitBody string) (*batch.SubmitPayload, error) {
	var err error
	var body SubmitRequestBody
	{
		err = json.Unmarshal([]byte(batchSubmitBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"path\": \"Dolor sit quia eum excepturi asperiores.\",\n      \"pipeline\": \"Mollitia velit quaerat quae.\"\n   }'")
		}
	}
	v := &batch.SubmitPayload{
		Path:     body.Path,
		Pipeline: body.Pipeline,
	}

	return v, nil
}