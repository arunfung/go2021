package models

import (
	"reflect"
	"testing"
)

func Test_findUserByName(t *testing.T) {
	type args struct {
		name string
	}
	var tests = []struct {
		name    string
		args    args
		want    *Users
		wantErr bool
	}{
		{
			name: "get name",
			args: args{name: "arun"},
			want: &Users{
				Id:       1,
				Username: "arun",
				Password: "123456",
				Email:    "arun@gmail.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindUserByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("findUserByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findUserByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
