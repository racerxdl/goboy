package audio

import "math"

const (
	SampleRate = 48000
	twoPi      = 2 * math.Pi
	perSample  = 1 / float64(SampleRate)
)

type Channel struct {
	frequency          float64
	time               float64
	amplitude          float64
	duration           int
	envelopeTime       int
	envelopeSteps      int
	envelopeStepsInit  int
	envelopeSamples    int
	envelopeIncreasing bool
	sweepTime          float64
	sweepStepLen       byte
	sweepSteps         byte
	sweepStep          byte
	sweepIncrease      bool
}

var sweepTimes = map[byte]float64{
	1: 7.8 / 1000,
	2: 15.6 / 1000,
	3: 23.4 / 1000,
	4: 31.3 / 1000,
	5: 39.1 / 1000,
	6: 46.9 / 1000,
	7: 54.7 / 1000,
}

func (c *Channel) Reset(duration int) {
	c.amplitude = 1
	c.envelopeTime = 0
	c.sweepTime = 0
	c.sweepStep = 0
	c.duration = duration
}

func (c *Channel) IsActive() bool {
	return (c.duration == -1 || c.duration > 0) && c.envelopeStepsInit > 0
}

func (c *Channel) Cycle() {
	if c.envelopeSamples > 0 {
		c.envelopeTime += 1
		if c.envelopeSteps > 0 && c.envelopeTime >= c.envelopeSamples {
			c.envelopeTime -= c.envelopeSamples
			c.envelopeSteps--
			if c.envelopeSteps == 0 {
				c.amplitude = 0
			} else if c.envelopeIncreasing {
				c.amplitude = 1 - float64(c.envelopeSteps)/float64(c.envelopeStepsInit)
			} else {
				c.amplitude = float64(c.envelopeSteps) / float64(c.envelopeStepsInit)
			}
		}
	}
	if c.sweepStep < c.sweepSteps {
		t := sweepTimes[c.sweepStepLen]
		c.sweepTime += perSample
		if c.sweepTime > t {
			c.sweepTime -= t
			c.sweepStep += 1

			if c.sweepIncrease {
				c.frequency += c.frequency / math.Pow(2, float64(c.sweepStep))
			} else {
				c.frequency -= c.frequency / math.Pow(2, float64(c.sweepStep))
			}
		}
	}
}
