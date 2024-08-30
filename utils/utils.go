package utils

import (
	"fmt"
	"strings"
)

// Función para reducir un mapa de errores en un string separado por comas
func ReduceErrorsToString(errors map[string]interface{}) string {
	var sb strings.Builder
	for _, err := range errors {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", err))
	}
	return sb.String()
}
