{{- define "calendar-chart.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end }}

{{- define "calendar-chart.name" -}}
{{- .Chart.Name -}}
{{- end }}

{{- define "calendar-chart.labels" -}}
app.kubernetes.io/name: {{ include "calendar-chart.name" . }}
helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version | replace "+" "_" }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}