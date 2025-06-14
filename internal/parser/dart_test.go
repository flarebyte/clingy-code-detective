package parser

import (
	"reflect"
	"testing"
)

const (
	yamlOnlyProd = `
dependencies:
  http: ^0.13.3
  path: ^1.8.0
`
	yamlOnlyDev = `
dev_dependencies:
  test: ^1.16.0
`
	yamlProdAndDev = `
dependencies:
  http: ^0.13.3
dev_dependencies:
  test: ^1.16.0
`
	yamlComplexVersion = `
dependencies:
  flutter:
    sdk: flutter
`
	yamlInvalid = `dependencies: [`
)

func Test_dartParser_Parse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []Dependency
		wantErr bool
	}{
		{
			name:  "only prod deps",
			input: yamlOnlyProd,
			want: []Dependency{
				{Name: "http", Version: "^0.13.3", Category: "prod"},
				{Name: "path", Version: "^1.8.0", Category: "prod"},
			},
		},
		{
			name:  "only dev deps",
			input: yamlOnlyDev,
			want: []Dependency{
				{Name: "test", Version: "^1.16.0", Category: "dev"},
			},
		},
		{
			name:  "prod and dev deps",
			input: yamlProdAndDev,
			want: []Dependency{
				{Name: "http", Version: "^0.13.3", Category: "prod"},
				{Name: "test", Version: "^1.16.0", Category: "dev"},
			},
		},
		{
			name:  "complex dependency version",
			input: yamlComplexVersion,
			want: []Dependency{
				{Name: "flutter", Version: "", Category: "prod"},
			},
		},
		{
			name:    "invalid yaml",
			input:   yamlInvalid,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := dartParser{}
			got, err := p.Parse([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
