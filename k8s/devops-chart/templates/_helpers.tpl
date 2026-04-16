{{- define "devops-app.envVars" -}}
- name: APP_ENV
  value: {{ .Values.environment | default "production" }}
- name: LOG_LEVEL
  value: {{ .Values.logLevel | default "info" }}
{{- end -}}