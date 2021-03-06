package qrcode

import (
	"testing"

	"github.com/yeqown/go-qrcode/matrix"
)

func TestMask(t *testing.T) {
	qrc, _ := New("baidu.com")
	qrc.initMatrix()

	var stateInitCnt int

	qrc.mat.Iter(matrix.ROW, func(x, y int, s matrix.State) {
		if s == matrix.StateInit {
			stateInitCnt++
		}
	})
	t.Logf("all state init count: %d", stateInitCnt)

	cpyMat := qrc.mat.Copy()
	drawAndSaveToFile("./testdata/mask_origin.jpeg", *cpyMat)

	mask0 := NewMask(cpyMat, Modulo0)
	drawAndSaveToFile("./testdata/modulo0.jpeg", *mask0.mat)

	mask1 := NewMask(cpyMat, Modulo1)
	drawAndSaveToFile("./testdata/modulo1.jpeg", *mask1.mat)

	mask2 := NewMask(cpyMat, Modulo2)
	drawAndSaveToFile("./testdata/modulo2.jpeg", *mask2.mat)

	mask3 := NewMask(cpyMat, Modulo3)
	drawAndSaveToFile("./testdata/modulo3.jpeg", *mask3.mat)

	mask4 := NewMask(cpyMat, Modulo4)
	drawAndSaveToFile("./testdata/modulo4.jpeg", *mask4.mat)

	mask5 := NewMask(cpyMat, Modulo5)
	drawAndSaveToFile("./testdata/modulo5.jpeg", *mask5.mat)

	mask6 := NewMask(cpyMat, Modulo6)
	drawAndSaveToFile("./testdata/modulo6.jpeg", *mask6.mat)

	mask7 := NewMask(cpyMat, Modulo7)
	drawAndSaveToFile("./testdata/modulo7.jpeg", *mask7.mat)
}
