{{- /* Go Template */ -}}

{{- define "exprstmt" -}}
{{ . }};
{{- end -}}

{{- define "print" -}}
    LRT_Print({{ . }});
{{- end -}}

{{- define "block" -}}
{
    {{- range . }}
    {{ . }}
    {{- end }}
}
{{- end -}}