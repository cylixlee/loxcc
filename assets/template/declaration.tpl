{{- /* Go Template */ -}}

{{- define "localvar" -}}
    LRT_Value {{ template "mangle" .name }} = {{ .initializer }};
{{- end -}}