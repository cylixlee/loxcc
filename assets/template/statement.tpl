{{- /* Go Template */ -}}


{{- define "exprstmt" -}}
{{ . }};
{{- end -}}


{{- define "for" -}}
// generated for-loop start
{
    {{ with .initializer }} {{ . }} {{ end }};
    while (!LRT_FalsinessOf({{- with .condition }} {{ . }} {{ else }} BOOLEAN(true) {{ end -}})) {
        {{ .body }}
        {{ with .incrementer }} {{ . }} {{ end }};
    }
}
// generated for-loop end
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


{{- define "return" -}}
    return {{ . }};
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

