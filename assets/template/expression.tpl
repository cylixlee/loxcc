{{- /* Go Template */ -}}

{{- /* Lox literal expressions */ -}}
{{- define "boolean" -}} BOOLEAN({{ . }})                      {{- end -}}
{{- define "nil"     -}} NIL                                   {{- end -}}
{{- define "number"  -}} NUMBER({{ . }})                       {{- end -}}
{{- define "strobj"  -}} LRT_NewString("{{ . }}", {{ len . }}) {{- end -}}
{{- define "string"  -}} OBJECT({{ template "strobj" . }})     {{- end -}}
{{- define "ident"   -}} {{ template "mangle" . }}             {{- end -}}
{{- define "fn"      -}}
    OBJECT(LRT_NewFunction({{ template "strobj" .name }}, {{ template "funmangle" .name }}))
{{- end -}}

{{- /* Assignment expression */ -}}
{{- define "assign" -}}
    (({{ .left }}) = ({{ .right }}))
{{- end -}}


{{- /* Binary expression */ -}}
{{- define "binary" -}}
    ({{ .operatorFunc }}({{ .left }}, {{ .right }}))
{{- end -}}


{{- /* Unary expression */ -}}
{{- define "unary" -}}
    ({{ .operatorFunc }}({{ .operand }}))
{{- end -}}