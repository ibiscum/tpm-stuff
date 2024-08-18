package main

import (
	"io"
	"reflect"
	"testing"

	"github.com/google/go-tpm/legacy/tpm2"
	"github.com/google/go-tpm/tpmutil"
)

func TestNewKey(t *testing.T) {
	// tpm2.KeyFromIndex(index)
}

func Test_defaultEKAuthPolicy(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultEKAuthPolicy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultEKAuthPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPrimary(t *testing.T) {
	type args struct {
		rw     io.ReadWriter
		hier   tpmutil.Handle
		handle tpmutil.Handle
		tmpl   tpm2.Public
	}
	tests := []struct {
		name    string
		args    args
		want    tpmutil.Handle
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPrimary(tt.args.rw, tt.args.hier, tt.args.handle, tt.args.tmpl)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPrimary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPrimary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createKey(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createKey(); (err != nil) != tt.wantErr {
				t.Errorf("createKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthCommand(t *testing.T) {
	type args struct {
		rw       io.ReadWriteCloser
		password []byte
		pcrSel   tpm2.PCRSelection
	}
	tests := []struct {
		name    string
		args    args
		want    tpm2.AuthCommand
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AuthCommand(tt.args.rw, tt.args.password, tt.args.pcrSel)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
