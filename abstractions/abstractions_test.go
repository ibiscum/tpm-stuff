package main

import (
	"io"
	"testing"

	"github.com/google/go-tpm/tpmutil"
)

func TestAuthSession_StartAuth(t *testing.T) {
	type fields struct {
		rwc            io.ReadWriteCloser
		sessionHandler tpmutil.Handle
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthSession{
				rwc:            tt.fields.rwc,
				sessionHandler: tt.fields.sessionHandler,
			}
			if err := a.StartAuth(); (err != nil) != tt.wantErr {
				t.Errorf("AuthSession.StartAuth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthSession_Flush(t *testing.T) {
	type fields struct {
		rwc            io.ReadWriteCloser
		sessionHandler tpmutil.Handle
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthSession{
				rwc:            tt.fields.rwc,
				sessionHandler: tt.fields.sessionHandler,
			}
			if err := a.Flush(); (err != nil) != tt.wantErr {
				t.Errorf("AuthSession.Flush() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
