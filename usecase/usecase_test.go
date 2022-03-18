package usecase_test

import (
	"errors"
	"integral_test/mocks"
	"integral_test/repository"
	"integral_test/usecase"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/undefinedlabs/go-mpatch"
)

func TestUsecase_Addition(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocks.NewMockDatabaseInterface(mockCtrl)

	type fields struct {
		Database repository.DatabaseInterface
	}
	type args struct {
		number1 string
		number2 string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
		patch   func()
	}{
		{
			name: "Succesfully Insert to Database",
			fields: fields{
				Database: mockDB,
			},
			args: args{
				number1: "50",
				number2: "20",
			},
			want:    70,
			wantErr: false,
			patch: func() {

				mockDB.
					EXPECT().
					Insert(gomock.Eq(70)).
					Return(nil)

			},
		},
		{
			name: "Failed to convert input to int",
			fields: fields{
				Database: mockDB,
			},
			args: args{
				number1: "50",
				number2: "20",
			},
			want:    0,
			wantErr: true,
			patch: func() {
				var monkeyPatch *mpatch.Patch
				monkeyPatch, err := mpatch.PatchMethod(usecase.ConvertStringToInt, func(s string) (int, error) {
					defer monkeyPatch.Unpatch()
					return 0, errors.New("")
				})
				if err != nil {
					t.Fatal(err)
				}
			},
		},
		{
			name: "Succesfully Insert to Database",
			fields: fields{
				Database: mockDB,
			},
			args: args{
				number1: "50",
				number2: "20",
			},
			want:    0,
			wantErr: true,
			patch: func() {

				mockDB.
					EXPECT().
					Insert(gomock.Eq(70)).
					Return(errors.New(""))

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase.Usecase{
				Database: tt.fields.Database,
			}
			tt.patch()
			got, err := u.Addition(tt.args.number1, tt.args.number2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.Addition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Usecase.Addition() = %v, want %v", got, tt.want)
			}
		})
	}
}
