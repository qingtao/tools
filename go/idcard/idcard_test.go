package idcard

import (
	"testing"
)

// func TestMatch(t *testing.T) {
// 	var a = []string{
// 		"34052419800101001X",
// 		"370683198901117657",
// 		"370683198901007657",
// 	}
// 	for _, s := range a {
// 		ss := re.FindStringSubmatch(s)
// 		t.Logf("%+v\n", ss)
// 	}
// }

func Test_Validate(t *testing.T) {
	type args struct {
		s      string
		gender []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s:      "34052419800101001X",
				gender: []int{0},
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				s:      "370683198901117657",
				gender: []int{1},
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				s:      "370683198901117667",
				gender: []int{1},
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				s:      "370683198901007667",
				gender: []int{1},
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				s: "34052419800101001X",
			},
			want: true,
		},
		{
			name: "6",
			args: args{
				s: "abc52419800101001X",
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				s: "身份证号校验", // 不是数字或者Xx
			},
			want: false,
		},
		{
			name: "8",
			args: args{
				s: "34052419800101001",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args.s, tt.args.gender...); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkCode(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{b: []byte("34052419800101001X")},
			want: 2,
		},
		{
			name: "2",
			args: args{b: []byte("34052419800101001x")},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkCode(tt.args.b); got != tt.want {
				t.Errorf("checkCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{b: []byte("34052419800101001X")},
			want: 189,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.b); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{b: []byte("34052419800101001X")},
			want: true,
		},
		{
			name: "2",
			args: args{b: []byte("34052419800100001X")},
			want: false,
		},
		{
			name: "3",
			args: args{b: []byte("34052419800101001x")},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.b); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_Validate(b *testing.B) {
	b.Run("1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Validate("34052419800101001X", 1)
		}
	})

	b.Run("2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Validate("370683198901117657", 1)
		}
	})

	b.Run("3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Validate("370683198901117657")
		}
	})
	b.Run("4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Validate("身份证号校验")
		}
	})
	b.Run("5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Validate("37068319890111657")
		}
	})

	b.Run("6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Validate("abc6*3198901117657")
		}
	})
}

func Benchmark_validate(b *testing.B) {
	b.Run("1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validate([]byte("34052419800101001X"))
		}
	})

	b.Run("2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validate([]byte("370683198901117657"))
		}
	})
}
