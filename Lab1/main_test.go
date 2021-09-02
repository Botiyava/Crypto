package main

import "testing"

func TestMain_encrypt(t *testing.T){
	tests := []struct{
		name string
		inputPlainText string
		inputKey string
		want string
	}{
		{"correct input",
			"HELLOWORLD",
			"KEY",
			"DANZQCWNNH",
			},
	}
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			got := encrypt(test.inputPlainText, test.inputKey)
			if got != test.want{
				t.Errorf("encrypt: got: %v, want: %v", got, test.want)
			}
		})
	}
}
