package expandhost

import (
	"reflect"
	"testing"
)

func TestPatternToHosts(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{pattern: "foo[01-03].idc1.bar.com"},
			want:    []string{"foo01.idc1.bar.com", "foo02.idc1.bar.com", "foo03.idc1.bar.com"},
			wantErr: false,
		},
		{
			name: "case2",
			args: args{pattern: "foo[01-03,12].idc[1-3].bar.com"},
			want: []string{
				"foo01.idc1.bar.com",
				"foo01.idc2.bar.com",
				"foo01.idc3.bar.com",
				"foo02.idc1.bar.com",
				"foo02.idc2.bar.com",
				"foo02.idc3.bar.com",
				"foo03.idc1.bar.com",
				"foo03.idc2.bar.com",
				"foo03.idc3.bar.com",
				"foo12.idc1.bar.com",
				"foo12.idc2.bar.com",
				"foo12.idc3.bar.com",
			},
			wantErr: false,
		},
		{
			name: "case3",
			args: args{pattern: "foo[ 01-03, 12].idc[1-3 ].bar.com"},
			want: []string{
				"foo01.idc1.bar.com",
				"foo01.idc2.bar.com",
				"foo01.idc3.bar.com",
				"foo02.idc1.bar.com",
				"foo02.idc2.bar.com",
				"foo02.idc3.bar.com",
				"foo03.idc1.bar.com",
				"foo03.idc2.bar.com",
				"foo03.idc3.bar.com",
				"foo12.idc1.bar.com",
				"foo12.idc2.bar.com",
				"foo12.idc3.bar.com",
			},
			wantErr: false,
		},
		{
			name: "case4",
			args: args{pattern: "foo[01-03,12].[beijing,wuhan].bar.com"},
			want: []string{
				"foo01.beijing.bar.com",
				"foo01.wuhan.bar.com",
				"foo02.beijing.bar.com",
				"foo02.wuhan.bar.com",
				"foo03.beijing.bar.com",
				"foo03.wuhan.bar.com",
				"foo12.beijing.bar.com",
				"foo12.wuhan.bar.com",
			},
			wantErr: false,
		},
		{
			name:    "case5",
			args:    args{pattern: "foo[01-03,12].[a-d].bar.com"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PatternToHosts(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("PatternToHosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PatternToHosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
