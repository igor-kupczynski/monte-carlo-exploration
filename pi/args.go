package pi

import (
	"fmt"
	"image"
	"os"

	"github.com/igor-kupczynski/monte-carlo-exploration/montecarlo"
)

// Args represents the parameters of the Pi estimation simulation
type Args struct {
	// How many sample histories do we want to run during the experiment
	Histories int

	// How many darts do we want to throw for each sample
	Rounds int

	// Which image to use as the target
	Image string
}

func (a *Args) String() string {
	return fmt.Sprintf("Pi estimation: %d executions of %d dart throws to %s\n", a.Histories, a.Rounds, a.Image)
}

// Returns new experiment based on the args
func New(args *Args) (montecarlo.Experiment, error) {
	reader, err := os.Open(args.Image)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	x0, y0 := img.Bounds().Min.X, img.Bounds().Min.Y
	x1, y1 := img.Bounds().Max.X, img.Bounds().Max.Y

	hitmap := make([][]bool, y1-y0)

	for y := 0; y < y1-y0; y++ {
		hitmap[y] = make([]bool, x1-x0)
		for x := 0; x < x1-x0; x++ {
			r, g, b, _ := img.At(x+x0, y+y0).RGBA()
			hitmap[y][x] = isBlack(r) && isBlack(g) && isBlack(b)
		}
	}

	states := make([]*state, args.Histories)
	for i := range states {
		states[i] = &state{
			img: hitmap,

			hit:   0,
			total: 0,

			wantRounds: args.Rounds,
		}
	}

	return &experiment{states: states}, nil
}

// isBack returns true for colors #80 and below (so some grays are also black :)).
//
// Note the color is [0, 0xffff]
func isBlack(color uint32) bool {
	return color < 0xffff
}
