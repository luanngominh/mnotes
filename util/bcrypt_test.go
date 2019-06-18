package util

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid name",
			args: args{
				password: "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	type args struct {
		password       string
		hashedPassword string
	}

	pass := "123"
	hashed := "$2a$14$KsEv9cokGDW8ykftlutPIuEF9W08vcXCMBkrXqPP.zB6idAqwy1j6"

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Password valid",
			args: args{
				password:       pass,
				hashedPassword: hashed,
			},
			wantErr: false,
		},
		{
			name: "Password invalid",
			args: args{
				password:       pass,
				hashedPassword: "adsfjoajdsfojfds",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := VerifyPassword(tt.args.password, tt.args.hashedPassword); (err != nil) != tt.wantErr {
				t.Errorf("VerifyPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
