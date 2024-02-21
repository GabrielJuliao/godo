package main

import (
	"github.com/gabrieljuliao/godo/cmd/env"
	"reflect"
	"testing"
)

func Test_filterGodoArgs(t *testing.T) {
	env.InitEnv()
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{name: "Empty", args: []string{""}, want: []string{""}},
		{name: "Macro only", args: []string{"fake-macro"}, want: []string{"fake-macro"}},
		{name: "Macro+--godo-args, no macro,args", args: []string{"--godo-config-file", "config-test.yaml", "fake-macro"}, want: []string{"fake-macro"}},
		{name: "Macro+args, no --godo-args", args: []string{"fake-macro", "--fake-macro-args", "fake-macro-arg-value"}, want: []string{"fake-macro", "--fake-macro-args", "fake-macro-arg-value"}},
		{name: "Macro+args+--godo-args", args: []string{"--godo-config-editor", "code", "fake-macro", "--fake-macro-args", "fake-macro-arg-value"}, want: []string{"fake-macro", "--fake-macro-args", "fake-macro-arg-value"}},
		{name: "Macro+args+--godo-args", args: []string{"--godo-config-editor-args", "--arg1,value1,arg-2", "fake-macro", "--fake-macro-args", "fake-macro-arg-value"}, want: []string{"fake-macro", "--fake-macro-args", "fake-macro-arg-value"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterGodoArgs(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterGodoArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
