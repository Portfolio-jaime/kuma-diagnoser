package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExportDiagnosis(format, output string) error {
	if LastDiagnosis == nil {
		return fmt.Errorf("no hay diagnóstico ejecutado aún")
	}

	var content string
	switch format {
	case "json":
		data, _ := json.MarshalIndent(LastDiagnosis, "", "  ")
		content = string(data)
	case "md":
		content = exportToMarkdown(LastDiagnosis)
	default:
		return fmt.Errorf("formato no soportado: %s", format)
	}

	if output != "" {
		return os.WriteFile(output, []byte(content), 0644)
	}
	fmt.Println(content)
	return nil
}

func exportToMarkdown(d *Diagnosis) string {
	md := "## Diagnóstico Kuma\n\n"
	md += "### Control Plane Pods:\n"
	for _, pod := range d.ControlPlanePods {
		md += fmt.Sprintf("- %s\n", pod)
	}
	md += fmt.Sprintf("\n### Total Dataplanes: `%d`\n", d.DataplaneCount)
	if len(d.Warnings) > 0 {
		md += "\n### ⚠️ Warnings:\n"
		for _, w := range d.Warnings {
			md += fmt.Sprintf("- `%s`\n", w)
		}
	} else {
		md += "\n✅ Sin warnings\n"
	}
	return md
}
