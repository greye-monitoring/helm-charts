apiVersion: v1
kind: Pod
metadata:
  name: "{{ include "greye.fullname" . }}-test-monitoring"
  labels:
    {{- include "greye.labels" . | nindent 4 }}
  annotations:
    "helm.sh/hook": test
    "helm.sh/hook-weight": "5"
spec:
  {{- if .Values.serviceAccount.create }}
  serviceAccountName: {{ .Release.Name }}
  {{- end }}
  containers:
    - name: wget
      image: curlimages/curl:7.82.0
      command: ['sh', '-c']
      args:
        - |
          {{- if .Values.monitoring.enabled }}
          curl -s {{ include "greye.fullname" . }}:{{ .Values.service.port }}/metrics || exit 1
          echo "Monitoring endpoint accessible"
          {{- else }}
          echo "Monitoring not enabled, skipping test"
          {{- end }}
  restartPolicy: Never