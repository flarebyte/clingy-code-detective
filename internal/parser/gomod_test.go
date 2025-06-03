package parser

import (
	"reflect"
	"testing"
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
		// TODO: Add test cases.
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
