package model

import (
	"testing"

	ui "github.com/jsperandio/bfm/app/ui/model"
	"github.com/stretchr/testify/assert"
)

func TestNewLayoutDirFromUI(t *testing.T) {
	type args struct {
		lyt *ui.Layout
	}
	tests := []struct {
		name string
		args args
		want *LayoutDir
	}{
		{
			name: "when layout is given, layout dir is created with correct paths",
			args: args{
				lyt: &ui.Layout{
					FileName: "clean",
					FileExt:  "yaml",
					Structure: map[string]interface{}{
						"dir": map[string]interface{}{
							"app": map[string]interface{}{
								"database": map[string]interface{}{
									"migrations": nil,
									"seeds":      nil,
								},
								"infrastructure": nil,
								"domain":         nil,
								"interfaces":     nil,
								"logs":           nil,
								"usecases":       nil,
							},
						},
					},
				},
			},
			want: &LayoutDir{
				fullpaths: []string{"app/database/migrations", "app/database/seeds", "app/infrastructure", "app/domain", "app/interfaces", "app/logs", "app/usecases"},
			},
		},
		{
			name: "when layout is given, layout dir is created with correct paths 2",
			args: args{
				lyt: &ui.Layout{
					FileName: "hexame",
					FileExt:  "yaml",
					Structure: map[string]interface{}{
						"dir": map[string]interface{}{
							"app": map[string]interface{}{
								"domain": map[string]interface{}{
									"model":      nil,
									"repository": nil,
									"service":    nil,
								},
								"fx": map[string]interface{}{
									"module": map[string]interface{}{
										"provider":   nil,
										"server":     nil,
										"serverless": nil,
										"service":    nil,
									},
								},
								"provider":   nil,
								"server":     nil,
								"serverless": nil,
							},
							"cmd":         nil,
							"config":      nil,
							"deployments": nil,
						},
					},
				},
			},
			want: &LayoutDir{
				fullpaths: []string{"app/domain/model", "app/domain/repository", "app/domain/service", "app/fx/module/provider", "app/fx/module/server", "app/fx/module/serverless", "app/fx/module/service", "app/provider", "app/server", "app/serverless", "cmd", "config", "deployments"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLayoutDirFromUI(tt.args.lyt)

			assert.NotEqual(t, nil, got)
			assert.ElementsMatch(t, got.DirectPaths(), tt.want.DirectPaths())
		})
	}
}
