package corr

import (
	"testing"

	"github.com/go-math/support/assert"
)

func TestDecompose(t *testing.T) {
	m := uint32(5)

	C := []float64{
		+1.000000000000000e+00,
		+1.154127058177033e-01,
		+3.134709671301593e-01,
		-2.624605475789556e-01,
		-2.675628902376415e-01,
		+1.154127058177033e-01,
		+1.000000000000000e+00,
		+1.487775664601468e-01,
		+8.251722261621278e-01,
		+6.471493243665016e-01,
		+3.134709671301593e-01,
		+1.487775664601468e-01,
		+1.000000000000000e+00,
		+2.188948568766876e-01,
		-4.746506616202846e-01,
		-2.624605475789556e-01,
		+8.251722261621278e-01,
		+2.188948568766876e-01,
		+1.000000000000000e+00,
		+7.047860759340304e-01,
		-2.675628902376415e-01,
		+6.471493243665017e-01,
		-4.746506616202846e-01,
		+7.047860759340304e-01,
		+1.000000000000000e+00,
	}

	expectedM := []float64{
		-2.809842768614861e-01,
		+8.627295633325639e-01,
		-1.214267186196807e-01,
		+9.269379853101427e-01,
		+8.961281861000945e-01,
		+6.515118812441264e-01,
		+3.964760087163447e-01,
		+8.860907779360867e-01,
		+2.468650308871921e-01,
		-3.127013101588838e-01,
		+6.984395695466264e-01,
		+1.766729460543113e-01,
		-4.356464873918021e-01,
		-2.493887459101609e-01,
		+2.478421563724771e-01,
		-9.023381630769038e-02,
		+2.584336847835331e-01,
		-8.530725532353041e-02,
		-1.023061219107812e-01,
		-1.828307777260400e-01,
		+2.490242639661812e-02,
		-2.243157966027762e-02,
		-5.507717743518731e-02,
		+8.474867991403982e-02,
		-6.572166265854235e-02,
	}

	M, n, _ := Decompose(C, m, 1)

	assert.Equal(n, m, t)
	assert.AlmostEqual(abs(M), abs(expectedM), t)

	M, n, _ = Decompose(C, m, 0.75)

	assert.Equal(n, uint32(2), t)
	assert.AlmostEqual(abs(M), abs(expectedM[:m*2]), t)
}

func abs(data []float64) []float64 {
	for i := range data {
		if data[i] < 0 {
			data[i] *= -1
		}
	}

	return data
}