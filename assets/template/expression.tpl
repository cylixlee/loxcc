{{- /* Go Template */ -}}

{{- /* Lox literal expressions */ -}}
{{ define "boolean" -}} BOOLEAN({{ . }})                                      {{- end }}
{{ define "nil"     -}} NIL                                                   {{- end }}
{{ define "number"  -}} NUMBER({{ . }})                                       {{- end }}
{{ define "string"  -}} OBJECT(LRT_NewString({{ . }}, {{ minus (len .) 2 }})) {{- end }}

{{- /* Binary expression */ -}}
{{ define "binary" -}}
    ({{ .operatorFunc }}({{ .left }}, {{ .right }}))
{{- end}}

{{- /* Unary expression */ -}}
{{ define "unary" -}}
    ({{ .operatorFunc }}({{ .operand }}))
{{- end}}