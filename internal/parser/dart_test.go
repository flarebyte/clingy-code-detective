package parser

import (
	"reflect"
	"testing"
)

func Test_dartParser_Parse(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		p       dartParser
		args    args
		want    []Dependency
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := dartParser{}
			got, err := p.Parse(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("dartParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dartParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
