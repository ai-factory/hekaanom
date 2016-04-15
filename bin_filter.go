package hekaanom

import (
	"errors"
	"time"

	"github.com/mozilla-services/heka/pipeline"
)

type Binner interface {
	pipeline.HasConfigStruct
	pipeline.Plugin
	Connect(in <-chan AnomalousSpan, out chan<- Bin)
}

type BinConfig struct {
	BinWidth int64 `toml:"bin_width"`
}

type BinFilter struct {
	bins Bins
	*BinConfig
}

func (f *BinFilter) ConfigStruct() interface{} {
	return &BinConfig{}
}

func (f *BinFilter) Init(config interface{}) error {
	f.BinConfig = config.(*BinConfig)
	if f.BinConfig.BinWidth <= 0 {
		return errors.New("'bin_width' setting must be greater than zero.")
	}
	f.bins = Bins{}
	return nil
}

func (f *BinFilter) Connect(in <-chan AnomalousSpan, out chan<- Bin) {
	binWidth := time.Duration(f.BinConfig.BinWidth) * time.Second
	for span := range in {
		for _, binTime := range f.spanToBins(span) {
			bin, ok := f.bins[binTime]
			if !ok {
				bin = &Bin{
					Start: binTime,
					End:   binTime.Add(binWidth),
				}
				f.bins[binTime] = bin
			}
			bin.Count += 1
			bin.Entries = append(bin.Entries, span.Series)
			out <- *bin
		}
	}
}

func (f *BinFilter) spanToBins(span AnomalousSpan) []time.Time {
	binWidth := time.Duration(f.BinConfig.BinWidth) * time.Second
	bins := []time.Time{}
	now := span.Start.Truncate(binWidth)
	// Spans are inclusive on both ends, so we want to include where equal
	for !now.After(span.End) {
		bins = append(bins, now)
		now = now.Add(binWidth)
	}
	return bins
}
