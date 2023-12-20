package html

import "testing"

func TestRender(t *testing.T) {
	type args struct {
		path string
		data any
		tpls []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Case 1",
			args: args{
				path: "../../static/web/template",
				data: struct {
					ContentType string
					Title       string
					DDNSGO      struct {
						Status string
					}
				}{
					ContentType: "text/html; charset=GB2312",
					Title:       "DDNS-GO Manager",
					DDNSGO: struct {
						Status string
					}{
						Status: "status",
					},
				},
				tpls: []string{
					"http.tpl", "html.tpl", "index.tpl",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Render(tt.args.path, tt.args.data, tt.args.tpls...)
			println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
