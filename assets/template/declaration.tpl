{{- /* Go Template */ -}}


{{- define "funsig" -}}
LRT_Value {{ template "funmangle" .name }}(size_t arity, va_list args);
{{- end -}}


{{- define "fundef" -}}
LRT_Value {{ template "funmangle" .name }}(size_t arity, va_list args) {
    // Arity checking
    if (arity != {{ len .params }}) {
        LRT_Panic("{{ .name }} expected {{ len .params }} arguments");
    }

    // Argument evaluation
    {{ range $index, $param := .params }}
    LRT_Value {{ template "mangle" $param }} = va_arg(args, LRT_Value);
    {{ end }}

    // User logic
    {{ .body }}
}
{{- end -}}


{{- define "localvar" -}}
    LRT_Value {{ template "mangle" .name }} = {{ .initializer }};
{{- end -}}