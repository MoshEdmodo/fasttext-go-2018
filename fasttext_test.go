package fasttextgo

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	LoadModel("mosh-edmodo-fasttext.model.bin")
	fmt.Println(Predict("parallel lines at infinity"))
	fmt.Println(Predict("Partial differential equations in heat transfer"))

}
