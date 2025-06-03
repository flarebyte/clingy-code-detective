package parser

import (
	"reflect"
	"testing"
)

const (
	goModEmpty = ``

	goModSingleRequire = `
module example.com/my/module

go 1.21

require github.com/stretchr/testify v1.7.0
`

	goModMultipleWithIndirect = `
module example.com/my/module

go 1.20

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/pkg/errors v0.9.1 // indirect
)
`

	goModWithComments = `
module example.com/foo

require (
	// Web framework
	github.com/labstack/echo/v4 v4.9.0

	// Utilities
	golang.org/x/sys v0.15.0 // indirect
)
`
)

func Test_goModParser_Parse(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		p       goModParser
		args    args
		want    []Dependency
		wantErr bool
	}{
		{
			name: "Empty go.mod",
			args: args{
				content: []byte(goModEmpty),
			},
			want:    []Dependency{},
			wantErr: false,
		},
		{
			name: "Single dependency",
			args: args{
				content: []byte(goModSingleRequire),
			},
			want: []Dependency{
				{
					Name:     "github.com/stretchr/testify",
					Version:  "v1.7.0",
					Category: "prod",
				},
			},
			wantErr: false,
		},
		{
			name: "Multiple dependencies with indirect",
			args: args{
				content: []byte(goModMultipleWithIndirect),
			},
			want: []Dependency{
				{
					Name:     "github.com/gin-gonic/gin",
					Version:  "v1.8.1",
					Category: "prod",
				},
				{
					Name:     "github.com/pkg/errors",
					Version:  "v0.9.1",
					Category: "dev",
				},
			},
			wantErr: false,
		},
		{
			name: "Require block with comments and whitespace",
			args: args{
				content: []byte(goModWithComments),
			},
			want: []Dependency{
				{
					Name:     "github.com/labstack/echo/v4",
					Version:  "v4.9.0",
					Category: "prod",
				},
				{
					Name:     "golang.org/x/sys",
					Version:  "v0.15.0",
					Category: "dev",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := goModParser{}
			got, err := p.Parse(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("goModParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goModParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
