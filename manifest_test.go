package composer

import (
	"reflect"
	"testing"
)

func TestStringOrStrings_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		l       StringOrStrings
		want    []byte
		wantErr bool
	}{
		{"one license", StringOrStrings{"MIT"}, []byte("[\"MIT\"]"), false},
		{"two licenses", StringOrStrings{"LGPL-2.1-only", "GPL-3.0-or-later"}, []byte("[\"LGPL-2.1-only\",\"GPL-3.0-or-later\"]"), false},
		{"one long license", StringOrStrings{"(LGPL-2.1-only or GPL-3.0-or-later)"}, []byte("[\"(LGPL-2.1-only or GPL-3.0-or-later)\"]"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestStringOrStrings_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		l       StringOrStrings
		args    args
		wantErr bool
	}{
		{"license as a string", StringOrStrings{"test string"}, args{data: []byte("\"test string\"")}, false},
		{"one license as array", StringOrStrings{"(LGPL-2.1-only or GPL-3.0-or-later)"}, args{data: []byte("[\"(LGPL-2.1-only or GPL-3.0-or-later)\"]")}, false},
		{"two licenses as array", StringOrStrings{"LGPL-2.1-only", "GPL-3.0-or-later"}, args{data: []byte("[\"LGPL-2.1-only\", \"GPL-3.0-or-later\"]")}, false},
		{"error on non-string", StringOrStrings{""}, args{data: []byte("12345")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValueOrMap_MarshalJSON(t *testing.T) {
	type fields struct {
		Value    string
		Patterns map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{"single value", fields{Value: "one"}, []byte("\"one\""), false},
		{"another value", fields{Value: "auto"}, []byte("\"auto\""), false},
		{"multiple values", fields{Patterns: map[string]string{"one": "two", "three": "four"}}, []byte("{\"one\":\"two\",\"three\":\"four\"}"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pi := &ValueOrMap{
				Value: tt.fields.Value,
				Map:   tt.fields.Patterns,
			}
			got, err := pi.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValueOrMap_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		pi      ValueOrMap
		args    args
		wantErr bool
	}{
		{"single value", ValueOrMap{Value: "one"}, args{data: []byte("\"one\"")}, false},
		{"another value", ValueOrMap{Value: "auto"}, args{data: []byte("\"auto\"")}, false},
		{"multiple values", ValueOrMap{Map: map[string]string{"one": "two", "three": "four"}}, args{data: []byte("{\"one\":\"two\",\"three\":\"four\"}")}, false},
		{"error", ValueOrMap{}, args{data: []byte("true")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pi ValueOrMap
			if err := pi.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.pi, pi) {
				t.Errorf("UnmarshalJSON() %v not equal to %v", tt.pi, pi)
			}
		})
	}
}

func TestBoolOrString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolOrString
		want    []byte
		wantErr bool
	}{
		{"store", BoolOrString{Bool: true}, []byte("true"), false},
		{"do not store", BoolOrString{Bool: false}, []byte("false"), false},
		{"prompt", BoolOrString{String: "prompt"}, []byte("\"prompt\""), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolOrString_UnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		s       BoolOrString
		args    args
		wantErr bool
	}{
		{"true", BoolOrString{Bool: true}, args{[]byte("true")}, false},
		{"True", BoolOrString{Bool: true}, args{[]byte("True")}, false},
		{"1", BoolOrString{Bool: true}, args{[]byte("1")}, false},
		{"\"1\"", BoolOrString{Bool: true}, args{[]byte("\"1\"")}, false},
		{"false", BoolOrString{Bool: false}, args{[]byte("false")}, false},
		{"False", BoolOrString{Bool: false}, args{[]byte("False")}, false},
		{"0", BoolOrString{Bool: false}, args{[]byte("0")}, false},
		{"\"0\"", BoolOrString{Bool: false}, args{[]byte("\"0\"")}, false},
		{"prompt", BoolOrString{String: "prompt"}, args{[]byte("\"prompt\"")}, false},
		{"error", BoolOrString{}, args{[]byte("[\"error\"]")}, true},
	}
	for _, tt := range tests {
		var s BoolOrString
		t.Run(tt.name, func(t *testing.T) {
			if err := s.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.s, s) {
				t.Errorf("UnmarshalJSON() %v not equal to %v", tt.s, s)
			}
		})
	}
}

func TestIntString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		c       IntString
		want    []byte
		wantErr bool
	}{
		{"string 300", IntString("300"), []byte("\"300\""), false},
		{"string 300MiB", IntString("300MiB"), []byte("\"300MiB\""), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntString_UnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		c       IntString
		args    args
		wantErr bool
	}{
		{"int 300", IntString("300"), args{[]byte("300")}, false},
		{"string 300", IntString("300"), args{[]byte("\"300\"")}, false},
		{"string 300MiB", IntString("300MiB"), args{[]byte("\"300MiB\"")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c IntString
			if err := c.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.c, c) {
				t.Errorf("UnmarshalJSON() %v not equal to %v", tt.c, c)
			}
		})
	}
}

func TestBool_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		b       Bool
		want    []byte
		wantErr bool
	}{
		{"true", Bool(true), []byte("true"), false},
		{"false", Bool(false), []byte("false"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_UnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		b       Bool
		args    args
		wantErr bool
	}{
		{"true", Bool(true), args{[]byte("true")}, false},
		{"True", Bool(true), args{[]byte("True")}, false},
		{"1", Bool(true), args{[]byte("1")}, false},
		{"\"1\"", Bool(true), args{[]byte("\"1\"")}, false},
		{"false", Bool(false), args{[]byte("false")}, false},
		{"False", Bool(false), args{[]byte("False")}, false},
		{"0", Bool(false), args{[]byte("0")}, false},
		{"\"0\"", Bool(false), args{[]byte("\"0\"")}, false},
		{"error", Bool(false), args{[]byte("\"error\"")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b Bool
			if err := b.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.b, b) {
				t.Errorf("UnmarshalJSON() %v not equal to %v", tt.b, b)
			}
		})
	}
}

func TestPsr_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		p       Psr
		want    []byte
		wantErr bool
	}{
		{"\"UniqueGlobalClass\":\"\"", Psr(map[string][]string{"UniqueGlobalClass": {""}}), []byte("{\"UniqueGlobalClass\":[\"\"]}"), false},
		{"\"\":\"src/\"", Psr(map[string][]string{"": {"src/"}}), []byte("{\"\":[\"src/\"]}"), false},
		{"\"Monolog\\\\\":[\"src/\",\"lib/\"]", Psr(map[string][]string{"Monolog\\": {"src/", "lib/"}}), []byte("{\"Monolog\\\\\":[\"src/\",\"lib/\"]}"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPsr_UnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		p       Psr
		args    args
		wantErr bool
	}{
		{"\"UniqueGlobalClass\":\"\"", Psr(map[string][]string{"UniqueGlobalClass": {""}}), args{[]byte("{\"UniqueGlobalClass\":[\"\"]}")}, false},
		{"\"\":\"src/\"", Psr(map[string][]string{"": {"src/"}}), args{[]byte("{\"\":\"src/\"}")}, false},
		{"\"\":\"src/\"", Psr(map[string][]string{"": {"src/"}}), args{[]byte("{\"\":\"src/\"}")}, false},
		{"\"Monolog\\\\\":\"src/\",\"Vendor_Namespace_\":[\"abc/\",\"lib/\"]", Psr(map[string][]string{"Monolog\\": {"src/"}, "Vendor_Namespace_": {"abc/", "lib/"}}), args{[]byte("{\"Monolog\\\\\":\"src/\",\"Vendor_Namespace_\":[\"abc/\",\"lib/\"]}")}, false},
		{"\"Monolog\\\\\":\"src/\",\"Vendor_Namespace_\":\"abc/\"", Psr(map[string][]string{"Monolog\\": {"src/"}, "Vendor_Namespace_": {"abc/"}}), args{[]byte("{\"Monolog\\\\\":\"src/\",\"Vendor_Namespace_\":[\"abc/\"]}")}, false},
		{"\"Monolog\\\\\":true", Psr(map[string][]string{}), args{[]byte("{\"Monolog\\\\\":true}")}, true},
		{"\"String\"", Psr(map[string][]string{}), args{[]byte("\"Monolog\"")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Psr(map[string][]string{})
			if err := p.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.p, p) {
				t.Errorf("UnmarshalJSON() %v not equal to %v", tt.p, p)
			}
		})
	}
}

func TestRepositories_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		p       Repositories
		want    []byte
		wantErr bool
	}{
		{"1 repository", Repositories{Repository{Type: "test"}}, []byte("[{\"type\":\"test\"}]"), false},
		{"nothing", Repositories{}, []byte("[]"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %d, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositories_UnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		p       Repositories
		args    args
		wantErr bool
	}{
		{"Array", Repositories{Repository{Type: "test"}}, args{[]byte("[{\"type\":\"test\"}]")}, false},
		{"Object", Repositories{Repository{Type: "test"}}, args{[]byte("{\"1\":{\"type\":\"test\"}}")}, false},
		{"Error", Repositories{}, args{[]byte("\"Nothing\"")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Repositories{}
			if err := p.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.p, p) {
				t.Errorf("UnmarshalJSON() %v not equal to %v", tt.p, p)
			}
		})
	}
}
