package parser

import (
	"reflect"
	"testing"
)

func Test_pythonParser_Parse(t *testing.T) {
	const multipleDepsExaample = `
	requests==2.25.1
	flask==1.1.2
	numpy
	`
	const exampleWithEmptyLines = `
	# This is a comment
	requests==2.25.1
	
	# Another comment
	flask
	`
	const exampleWithTrailingSpaces = `
	   requests==2.25.1  
		flask==1.1.2
	numpy  
	`

	const exampleWithComment = `
	# just a comment
	
	# another
	   	# indented comment
	`
	tests := []struct {
		name    string
		input   string
		want    []Dependency
		wantErr bool
	}{
		{
			name:  "single dependency with version",
			input: "requests==2.25.1",
			want: []Dependency{
				{Name: "requests", Version: "2.25.1", Category: "prod"},
			},
		},
		{
			name:  "single dependency without version",
			input: "flask",
			want: []Dependency{
				{Name: "flask", Version: "", Category: "prod"},
			},
		},
		{
			name:  "multiple dependencies",
			input: multipleDepsExaample,
			want: []Dependency{
				{Name: "requests", Version: "2.25.1", Category: "prod"},
				{Name: "flask", Version: "1.1.2", Category: "prod"},
				{Name: "numpy", Version: "", Category: "prod"},
			},
		},
		{
			name:  "with comments and empty lines",
			input: exampleWithEmptyLines,
			want: []Dependency{
				{Name: "requests", Version: "2.25.1", Category: "prod"},
				{Name: "flask", Version: "", Category: "prod"},
			},
		},
		{
			name:  "trailing and leading whitespace",
			input: exampleWithTrailingSpaces,
			want: []Dependency{
				{Name: "requests", Version: "2.25.1", Category: "prod"},
				{Name: "flask", Version: "1.1.2", Category: "prod"},
				{Name: "numpy", Version: "", Category: "prod"},
			},
		},
		{
			name:  "only comments and whitespace",
			input: exampleWithComment,
			want:  []Dependency{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pythonParser{}
			got, err := p.Parse([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("pythonParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pythonParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
