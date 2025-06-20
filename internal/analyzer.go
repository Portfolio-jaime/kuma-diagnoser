package internal

import (
	"encoding/json"
	"fmt"
	"strings"
)

var LastDiagnosis *Diagnosis

type Diagnosis struct {
	ControlPlanePods []string `json:"control_plane_pods"`
	DataplaneCount   int      `json:"dataplane_count"`
	Warnings         []string `json:"warnings"`
}

func RunDiagnosis() error {
	controlPlanePods := RunKubectl("get pods -n kuma-system -l app=kuma-control-plane -o name")
	rawDP := RunKumactl("get dataplanes -o json")

	var dpCount int
	if len(rawDP) > 0 {
		var parsed map[string]interface{}
		if err := json.Unmarshal([]byte(strings.Join(rawDP, "")), &parsed); err == nil {
			if items, ok := parsed["items"].([]interface{}); ok {
				dpCount = len(items)
			}
		}
	}

	warnings := RunKubectl(`logs -n kuma-system -l app=kuma-control-plane --tail=500 | grep "initial fetch timed out"`)

	LastDiagnosis = &Diagnosis{
		ControlPlanePods: controlPlanePods,
		DataplaneCount:   dpCount,
		Warnings:         warnings,
	}

	fmt.Println("✅ Diagnóstico completado")
	return nil
}
