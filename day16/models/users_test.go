package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", "root", "root123456", "127.0.0.1", 3306, "go"))
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	m.Run()
}

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

func TestFindRoleById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    *UserRole
		wantErr bool
	}{
		{
			name: "get role",
			args: args{id: 1},
			want: &UserRole{
				Id:   1,
				Name: "管理员",
				Desc: "管理员",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindRoleById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRoleById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindRoleById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
