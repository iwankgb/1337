//chuck helps you to draw a random Chuck Norris joke
package chuck

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

type generator struct {
	scanner  *bufio.Scanner
	messages [129]string
}

func NewGenerator(path string) *generator {
	reader, _ := os.Open(path)
	scanner := bufio.NewScanner(reader)
	generator := &generator{scanner: scanner}
	return generator
}

func (generator *generator) Get() string {
	i := 0
	for generator.scanner.Scan() {
		generator.messages[i] = generator.scanner.Text()
		i = i + 1
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return generator.messages[r.Intn(i)]
}
