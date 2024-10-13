{{- /* Go Template */ -}}


{{- define "funsig" -}}
LRT_Value {{ template "funmangle" .name }}(size_t arity, ...);
{{- end -}}


{{- define "fundef" -}}
LRT_Value {{ template "funmangle" .name }}(size_t arity, ...) {
    if (arity != {{ len .params }}) {
        LRT_Panic("{{ .name }} expected {{ len .params }} arguments");
    }

    va_list varlist;
    va_start(varlist, arity);
    {{ range $index, $param := .params }}
    LRT_Value {{ template "mangle" $param }} = va_arg(varlist, LRT_Value);
    {{ end }}
    va_end(varlist);

    {{ .body }}
}
{{- end -}}


{{- define "localvar" -}}
    LRT_Value {{ template "mangle" .name }} = {{ .initializer }};
{{- end -}}