package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "/home/white/nfs/Drama/The.Legend.of.Vox.Machina.S01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args[1] = tt.name
			main()
		})
	}
}
