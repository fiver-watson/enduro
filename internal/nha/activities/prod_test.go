package activities

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-logr/logr"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/fs"

	collectionfake "github.com/artefactual-labs/enduro/internal/collection/fake"
	"github.com/artefactual-labs/enduro/internal/nha"
	"github.com/artefactual-labs/enduro/internal/pipeline"
	"github.com/artefactual-labs/enduro/internal/workflow/manager"
)

func TestProdActivity(t *testing.T) {
	t.Parallel()

	tmpdir, err := os.MkdirTemp("", "")
	assert.NilError(t, err)
	defer os.RemoveAll(tmpdir)

	tests := map[string]struct {
		params                UpdateProductionSystemActivityParams
		hookConfig            *map[string]interface{}
		dirOpts               []fs.PathOp
		wantContent           string
		wantNonRetryableError bool
		wantErr               string
	}{
		"Receipt is generated successfully with status 'done'": {
			params: UpdateProductionSystemActivityParams{
				StoredAt:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				PipelineName: "foo-bar-001",
				NameInfo: nha.NameInfo{
					Identifier: "aa1df25d-1477-4085-8be3-a17fed20f843",
					Type:       nha.TransferTypeDPJ,
				},
			},
			dirOpts: []fs.PathOp{
				fs.WithDir("DPJ/journal"),
				fs.WithFile("DPJ/journal/avlxml.xml", "<xml/>"),
				fs.WithDir("metadata"),
				fs.WithFile("metadata/identifiers.json", `[{
					"file": "objects/DPJ/journal/avlxml.xml",
					"identifiers": [{
						"identifierType": "avleveringsidentifikator",
						"identifier": "3.15.578.1.39.120.11.9896.12"
					}]
				}]`),
			},
			wantContent: `{
  "identifier": "aa1df25d-1477-4085-8be3-a17fed20f843",
  "type": "dpj",
  "accepted": true,
  "message": "Package was processed by Archivematica pipeline foo-bar-001",
  "timestamp": "2009-11-10T23:00:00Z",
  "parent": "3.15.578.1.39.120.11.9896.12"
}
`,
		},
		"Receipt does not include parentID in AVLXML SIP": {
			params: UpdateProductionSystemActivityParams{
				StoredAt:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				PipelineName: "foo-bar-001",
				NameInfo: nha.NameInfo{
					Identifier: "3.15.578.1.39.120.11.9896.12",
					Type:       nha.TransferTypeAVLXML,
				},
			},
			dirOpts: []fs.PathOp{
				fs.WithDir("DPJ/journal"),
				fs.WithFile("DPJ/journal/avlxml.xml", "<xml/>"),
			},
			wantContent: `{
  "identifier": "3.15.578.1.39.120.11.9896.12",
  "type": "avlxml",
  "accepted": true,
  "message": "Package was processed by Archivematica pipeline foo-bar-001",
  "timestamp": "2009-11-10T23:00:00Z"
}
`,
		},
		"Missing receiptPath is rejected": {
			params: UpdateProductionSystemActivityParams{
				StoredAt:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				PipelineName: "foo-bar-001",
				NameInfo: nha.NameInfo{
					Identifier: "aa1df25d-1477-4085-8be3-a17fed20f843",
					Type:       nha.TransferTypeDPJ,
				},
			},
			hookConfig:            &map[string]interface{}{},
			wantErr:               "error looking up receiptPath configuration attribute: error accessing \"prod:receiptPath\"",
			wantNonRetryableError: true,
		},
		"Unexistent receiptPath is rejected": {
			params: UpdateProductionSystemActivityParams{
				StoredAt:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				PipelineName: "foo-bar-001",
				NameInfo: nha.NameInfo{
					Identifier: "aa1df25d-1477-4085-8be3-a17fed20f843",
					Type:       nha.TransferTypeDPJ,
				},
			},
			hookConfig: &map[string]interface{}{
				"receiptpath": tmpdir,
			},
			wantErr:               "error creating receipt file",
			wantNonRetryableError: true,
		},
	}

	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			transferDir := fs.NewDir(t, "enduro", tc.dirOpts...)
			defer transferDir.Remove()
			tc.params.FullPath = transferDir.Path()

			tmpdir := fs.NewDir(t, "enduro")
			defer tmpdir.Remove()
			hookConfig := map[string]interface{}{"receiptpath": tmpdir.Path()}
			if tc.hookConfig != nil {
				hookConfig = *tc.hookConfig
			}

			act := createProdActivity(t, hookConfig)

			err := act.Execute(context.Background(), &tc.params)

			testError(t, err, tc.wantErr, tc.wantNonRetryableError)

			// Stop here if we were expecting the activity to fail.
			if tc.wantErr != "" {
				return
			}

			base := fmt.Sprintf("Receipt_%s_%s", tc.params.NameInfo.Identifier, tc.params.StoredAt.Format(rfc3339forFilename))
			assert.Assert(t, fs.Equal(
				tmpdir.Path(),
				fs.Expected(t,
					fs.WithFile(
						base+".mft",
						tc.wantContent,
						fs.WithMode(os.FileMode(0o644))),
				)))
		})
	}
}

func createProdActivity(t *testing.T, hookConfig map[string]interface{}) *UpdateProductionSystemActivity {
	t.Helper()

	ctrl := gomock.NewController(t)

	manager := manager.NewManager(
		logr.Discard(),
		collectionfake.NewMockService(ctrl),
		&pipeline.Registry{},
		map[string]map[string]interface{}{
			"prod": hookConfig,
		},
	)

	return NewUpdateProductionSystemActivity(manager)
}
