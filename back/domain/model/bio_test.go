package model

import (
	"reflect"
	"testing"
)

func makeContent(num int) string {
	content := ""
	for i := 0; i < num; i++ {
		content += "a"
	}
	return content
}

func TestNewBio(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    *Bio
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				content: makeContent(MaxBioLen),
			},
			want: &Bio{
				content: makeContent(MaxBioLen),
			},
			wantErr: false,
		},
		{
			name: "failure: contentが長すぎる",
			args: args{
				content: makeContent(MaxBioLen + 1),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success: contentが空",
			args: args{
				content: "",
			},
			want: &Bio{
				content: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBio(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBio() = %v, want %v", got, tt.want)
			}
		})
	}
}
