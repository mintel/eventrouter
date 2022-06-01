package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
)

func TestSanitiseObjName(t *testing.T) {
	tests := []struct {
		name  string
		input v1.Event
		want  string
	}{
		{
			name: "Standard ExternalSecret name",
			input: v1.Event{
				InvolvedObject: v1.ObjectReference{
					Name: "image-pull-gitlab",
					Kind: "ExternalSecret",
				},
			},
			want: "image-pull-gitlab",
		},
		{
			name: "Standard Pod name",
			input: v1.Event{
				InvolvedObject: v1.ObjectReference{
					Name: "node-exporter-9mj76",
					Kind: "Pod",
				},
			},
			want: "node-exporter-9mj76",
		},
		{
			name: "CronJob Job name",
			input: v1.Event{
				InvolvedObject: v1.ObjectReference{
					Name: "consul-create-sync-metrics-token-27566595",
					Kind: "Job",
				},
			},
			want: "consul-create-sync-metrics-token",
		},
		{
			name: "CronJob Pod name",
			input: v1.Event{
				InvolvedObject: v1.ObjectReference{
					Name: "consul-create-sync-kubeconfig-27566550--1-78klg",
					Kind: "Pod",
				},
			},
			want: "consul-create-sync-kubeconfig",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := sanitiseObjName(&tt.input)
			assert.Equal(t, tt.want, output)
		})
	}
}
