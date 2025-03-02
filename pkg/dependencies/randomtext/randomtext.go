package randomtext

import "math/rand"

func GenerateRandomText() string {
	adjIdx := rand.Intn(len(adjs))
	nounIdx := rand.Intn(len(nouns))

	return adjs[adjIdx] + " " + nouns[nounIdx]
}
