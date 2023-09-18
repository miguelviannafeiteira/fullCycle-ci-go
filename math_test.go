package main

import "testing"

func testSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Resultadoa inváalido: Resultado %d Esperado: %d", total, 30)
	}
}
