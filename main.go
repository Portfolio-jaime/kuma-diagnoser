package main

import (
	"fmt"
	"log"

	"github.com/Portfolio-jaime/kuma-diagnoser/internal"
)

func main() {
	allowedClusters := []string{
		"nexus-infradev-eks-cluster",
		"nexus-infradev2-eks-cluster",
		"nexus-dev-eks-cluster",
		"nexus-sit-eks-cluster",
		"nexus-uat1-eks-cluster",
		"nexus-pcm-eks-cluster",
		"nexus-prod-eks-cluster",
	}

	if err := internal.ValidateTools(); err != nil {
		log.Fatalf("Error validando herramientas: %v", err)
	}

	if err := internal.ValidateCluster(allowedClusters); err != nil {
		log.Fatalf("Error validando cluster: %v", err)
	}

	fmt.Println("🎉 ¡Todas las validaciones pasaron exitosamente! El entorno está listo para continuar.")
}
