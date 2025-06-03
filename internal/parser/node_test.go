package parser

import (
	"reflect"
	"testing"
)

func Test_nodeParser_Parse(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		p       nodeParser
		args    args
		want    []Dependency
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := nodeParser{}
			got, err := p.Parse(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("nodeParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nodeParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
