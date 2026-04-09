package locale

import "testing"

func TestGetLocaleFromFileName(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with standard locale format",
			args: args{fileName: "uk_UA.ftl"},
			want: "uk_UA",
		},
		{
			name: "Test with file name without extension",
			args: args{fileName: "en_US"},
			want: "en_US",
		},
		{
			name: "Test with file name without underscore",
			args: args{fileName: "en.ftl"},
			want: "en_US",
		},
		{
			name: "Test with file name without underscore and no extension",
			args: args{fileName: "/home/vsl/go/src/summit/resapi/lang/en.ftl"},
			want: "en_US",
		},
	}
	for _, tt := range tests {
		if got := GetLocaleFromFileName(tt.args.fileName); got != tt.want {
			t.Errorf("%q. GetLocaleFromFileName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
