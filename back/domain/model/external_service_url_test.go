package model

import (
	"reflect"
	"testing"
)

func TestNewExternalServiceUrl(t *testing.T) {
	type args struct {
		id          string
		serviceType int
		url         string
	}
	tests := []struct {
		name    string
		args    args
		want    *ExternalServiceUrl
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id:          "1",
				serviceType: 1,
				url:         "https://github.com",
			},
			want: &ExternalServiceUrl{
				id:          "1",
				serviceType: 1,
				url:         "https://github.com",
			},
			wantErr: false,
		},
		{
			name: "failure: サービスが想定外",
			args: args{
				id:          "2",
				serviceType: 6,
				url:         "https://github.com",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure: URLが不正",
			args: args{
				id:          "3",
				serviceType: 5,
				url:         "hogehoge",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure: urlが空",
			args: args{
				id:          "4",
				serviceType: 5,
				url:         "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewExternalServiceUrl(tt.args.id, tt.args.serviceType, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewExternalServiceUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExternalServiceUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateUrlLogic(t *testing.T) {
	type args struct {
		url         string
		serviceType int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success: github",
			args: args{
				url:         "https://github.com/elca-hub",
				serviceType: ExternalServiceGithub,
			},
			want:    "https://github.com/elca-hub",
			wantErr: false,
		},
		{
			name: "success: x",
			args: args{
				url:         "https://x.com/elca_e1",
				serviceType: ExternalServiceX,
			},
			want:    "https://x.com/elca_e1",
			wantErr: false,
		},
		{
			name: "success: qiita",
			args: args{
				url:         "https://qiita.com/elca",
				serviceType: ExternalServiceQiita,
			},
			want:    "https://qiita.com/elca",
			wantErr: false,
		},
		{
			name: "success: zenn",
			args: args{
				url:         "https://zenn.dev/momiji",
				serviceType: ExternalServiceZenn,
			},
			want:    "https://zenn.dev/momiji",
			wantErr: false,
		},
		{
			name: "success: note",
			args: args{
				url:         "https://note.com/elca_note",
				serviceType: ExternalServiceNote,
			},
			want:    "https://note.com/elca_note",
			wantErr: false,
		},
		{
			name: "failure: URLが想定するサービス外のURL",
			args: args{
				url:         "https://google.com",
				serviceType: ExternalServiceGithub,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "failure: URLが空",
			args: args{
				url:         "",
				serviceType: ExternalServiceGithub,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "failure: serviceTypeが不正",
			args: args{
				url:         "https://github.com",
				serviceType: 6,
			},
			want:    "",
			wantErr: true,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updateUrlLogic(tt.args.url, tt.args.serviceType)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateUrlLogic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("updateUrlLogic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExternalServiceUrl_UpdateUrl(t *testing.T) {
	type fields struct {
		url         string
		serviceType int
		id          string
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				url:         "https://github.com/elca-hub",
				serviceType: ExternalServiceGithub,
				id:          "1",
			},
			args: args{
				url: "https://github.com/elca_e1",
			},
			wantErr: false,
		},
		{
			name: "failure: URLが指定されているserviceTypeと異なる",
			fields: fields{
				url:         "https://github.com/elca-hub",
				serviceType: ExternalServiceGithub,
				id:          "2",
			},
			args: args{
				url: "https://x.com/elca_e1",
			},
			wantErr: true,
		},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExternalServiceUrl{
				url:         tt.fields.url,
				serviceType: tt.fields.serviceType,
				id:          tt.fields.id,
			}
			if err := e.UpdateUrl(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("ExternalServiceUrl.UpdateUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExternalServiceUrl_Url(t *testing.T) {
	type fields struct {
		url         string
		serviceType int
		id          string
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
			e := &ExternalServiceUrl{
				url:         tt.fields.url,
				serviceType: tt.fields.serviceType,
				id:          tt.fields.id,
			}
			if got := e.Url(); got != tt.want {
				t.Errorf("ExternalServiceUrl.Url() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExternalServiceUrl_ServiceType(t *testing.T) {
	type fields struct {
		url         string
		serviceType int
		id          string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExternalServiceUrl{
				url:         tt.fields.url,
				serviceType: tt.fields.serviceType,
				id:          tt.fields.id,
			}
			if got := e.ServiceType(); got != tt.want {
				t.Errorf("ExternalServiceUrl.ServiceType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExternalServiceUrl_ID(t *testing.T) {
	type fields struct {
		url         string
		serviceType int
		id          string
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
			e := &ExternalServiceUrl{
				url:         tt.fields.url,
				serviceType: tt.fields.serviceType,
				id:          tt.fields.id,
			}
			if got := e.ID(); got != tt.want {
				t.Errorf("ExternalServiceUrl.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateServiceType(t *testing.T) {
	type args struct {
		serviceType int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updateServiceType(tt.args.serviceType)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateServiceType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("updateServiceType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExternalServiceUrl_UpdateServiceType(t *testing.T) {
	type fields struct {
		url         string
		serviceType int
		id          string
	}
	type args struct {
		serviceType int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExternalServiceUrl{
				url:         tt.fields.url,
				serviceType: tt.fields.serviceType,
				id:          tt.fields.id,
			}
			if err := e.UpdateServiceType(tt.args.serviceType); (err != nil) != tt.wantErr {
				t.Errorf("ExternalServiceUrl.UpdateServiceType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
