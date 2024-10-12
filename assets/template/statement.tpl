{{- /* Go Template */ -}}


{{- define "exprstmt" -}}
{{ . }};
{{- end -}}


{{- define "if" -}}
if (!LRT_FalsinessOf({{ .condition }})) {{ .then }}
{{ with .else -}}
else {{ . }}
{{ end }}
{{- end -}}


{{- define "print" -}}
    LRT_Print({{ . }});
{{- end -}}


{{- define "while" -}}
while (!LRT_FalsinessOf({{ .condition }})) {{ .body }}
{{- end -}}


{{- define "block" -}}
{
    {{- range . }}
    {{ . }}
    {{- end }}
}
{{- end -}}

