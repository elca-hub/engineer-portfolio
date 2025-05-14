package model

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestNewBio(t *testing.T) {
	os.Setenv("BIO_RESOURCE_DOMAIN", "http://localhost:9000")

	overflowContent := strings.Repeat("a", MaxBioLen+1)

	overflowImageIds := make([]string, MaxBioImagesLen+1)
	for i := range overflowImageIds {
		overflowImageIds[i] = fmt.Sprintf("![](http://localhost:9000/devport/bio_images/%s.png)", uuid.New().String())
	}

	type args struct {
		userId  string
		id      string
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
				userId: "test",
				id:     "test",
				content: `
				![image](http://localhost:9000/devport/bio_images/test.jpeg)
				# this is title
				## this is subtitle

				I am a test user.

				![test2](http://localhost:9000/devport/bio_images/test2.png)
				`,
			},
			want: &Bio{
				id: "test",
				content: `
				![image](http://localhost:9000/devport/bio_images/test.jpeg)
				# this is title
				## this is subtitle

				I am a test user.

				![test2](http://localhost:9000/devport/bio_images/test2.png)
				`,
				imageIds:   []string{"test.jpeg", "test2.png"},
				objectName: "bio/test/test.md",
			},
			wantErr: false,
		},
		{
			name: "success 同じ画像IDの場合は1つにまとめる",
			args: args{
				userId: "test",
				id:     "test",
				content: `
				![image](http://localhost:9000/devport/bio_images/test.jpeg)
				![image](http://localhost:9000/devport/bio_images/test2.png)
				![image](http://localhost:9000/devport/bio_images/test2.png)
				`,
			},
			want: &Bio{
				id: "test",
				content: `
				![image](http://localhost:9000/devport/bio_images/test.jpeg)
				![image](http://localhost:9000/devport/bio_images/test2.png)
				![image](http://localhost:9000/devport/bio_images/test2.png)
				`,
				imageIds:   []string{"test.jpeg", "test2.png"},
				objectName: "bio/test/test.md",
			},
			wantErr: false,
		},
		{
			name: fmt.Sprintf("fail コンテンツが%d文字以上", MaxBioLen+1),
			args: args{
				userId:  "test",
				id:      "test",
				content: overflowContent,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: fmt.Sprintf("fail 画像IDが%d個以上", MaxBioImagesLen+1),
			args: args{
				userId:  "test",
				id:      "test",
				content: strings.Join(overflowImageIds, "\n"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBio(tt.args.userId, tt.args.id, tt.args.content)
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

func TestBio_ID(t *testing.T) {
	type fields struct {
		id         string
		content    string
		objectName string
		imageIds   []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bio{
				id:         tt.fields.id,
				content:    tt.fields.content,
				objectName: tt.fields.objectName,
				imageIds:   tt.fields.imageIds,
			}
			if got := b.ID(); got != tt.want {
				t.Errorf("Bio.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBio_Content(t *testing.T) {
	type fields struct {
		id         string
		content    string
		objectName string
		imageIds   []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				content: `
				![test](https://localhost:9000/bio_images/test.png)
				# this is title
				## this is subtitle

				I am a test user.

				![test2](https://localhost:9000/bio_images/test2.png)
				`,
			},
			want: `
				![test](https://localhost:9000/bio_images/test.png)
				# this is title
				## this is subtitle

				I am a test user.

				![test2](https://localhost:9000/bio_images/test2.png)
				`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bio{
				id:         tt.fields.id,
				content:    tt.fields.content,
				objectName: tt.fields.objectName,
				imageIds:   tt.fields.imageIds,
			}
			if got := b.Content(); got != tt.want {
				t.Errorf("Bio.Content() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBio_ObjectName(t *testing.T) {
	type fields struct {
		id         string
		content    string
		objectName string
		imageIds   []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bio{
				id:         tt.fields.id,
				content:    tt.fields.content,
				objectName: tt.fields.objectName,
				imageIds:   tt.fields.imageIds,
			}
			if got := b.ObjectName(); got != tt.want {
				t.Errorf("Bio.ObjectName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBio_ImageIds(t *testing.T) {
	type fields struct {
		id         string
		content    string
		objectName string
		imageIds   []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bio{
				id:         tt.fields.id,
				content:    tt.fields.content,
				objectName: tt.fields.objectName,
				imageIds:   tt.fields.imageIds,
			}
			if got := b.ImageIds(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bio.ImageIds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBio_IsFullImage(t *testing.T) {
	type fields struct {
		id         string
		content    string
		objectName string
		imageIds   []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bio{
				id:         tt.fields.id,
				content:    tt.fields.content,
				objectName: tt.fields.objectName,
				imageIds:   tt.fields.imageIds,
			}
			if got := b.IsFullImage(); got != tt.want {
				t.Errorf("Bio.IsFullImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
