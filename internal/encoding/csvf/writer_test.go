// Copyright 2020 Eurac Research. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package csvf

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/euracresearch/browser"
	"github.com/google/go-cmp/cmp"
)

func TestWrite(t *testing.T) {
	testCases := map[string]struct {
		in   browser.TimeSeries
		want string
	}{
		"empty": {
			browser.TimeSeries{},
			"",
		},
		"one_station_one_measurement": {
			browser.TimeSeries{
				testMeasurement("a_avg", "s1", "c", nil, 5),
			},
			`station,s1
landuse,me_s1
latitude,3.14159
longitude,2.71828
elevation,1000
parameter,a
depth,
aggregation,avg
unit,c
2020-01-01 00:15:00,0
2020-01-01 00:30:00,1
2020-01-01 00:45:00,2
2020-01-01 01:00:00,3
2020-01-01 01:15:00,4
`,
		},
		"two_station_one_measurement": {
			browser.TimeSeries{
				testMeasurement("a_avg", "s1", "c", nil, 5),
				testMeasurement("a_avg", "s2", "c", nil, 5),
			},
			`station,s1,s2
landuse,me_s1,me_s2
latitude,3.14159,3.14159
longitude,2.71828,2.71828
elevation,1000,1000
parameter,a,a
depth,,
aggregation,avg,avg
unit,c,c
2020-01-01 00:15:00,0,0
2020-01-01 00:30:00,1,1
2020-01-01 00:45:00,2,2
2020-01-01 01:00:00,3,3
2020-01-01 01:15:00,4,4
`,
		},
		"two_with_first_less_points": {
			browser.TimeSeries{
				testMeasurement("a_avg", "s1", "c", nil, 3),
				testMeasurement("a_avg", "s2", "c", nil, 5),
			},
			`station,s1,s2
landuse,me_s1,me_s2
latitude,3.14159,3.14159
longitude,2.71828,2.71828
elevation,1000,1000
parameter,a,a
depth,,
aggregation,avg,avg
unit,c,c
2020-01-01 00:15:00,0,0
2020-01-01 00:30:00,1,1
2020-01-01 00:45:00,2,2
2020-01-01 01:00:00,NaN,3
2020-01-01 01:15:00,NaN,4
`,
		},
		"two_with_last_less_points": {
			browser.TimeSeries{
				testMeasurement("a_avg", "s1", "c", nil, 5),
				testMeasurement("a_avg", "s2", "c", nil, 2),
			},
			`station,s1,s2
landuse,me_s1,me_s2
latitude,3.14159,3.14159
longitude,2.71828,2.71828
elevation,1000,1000
parameter,a,a
depth,,
aggregation,avg,avg
unit,c,c
2020-01-01 00:15:00,0,0
2020-01-01 00:30:00,1,1
2020-01-01 00:45:00,2,NaN
2020-01-01 01:00:00,3,NaN
2020-01-01 01:15:00,4,NaN
`,
		},
		"three_with_middle_less_points": {
			browser.TimeSeries{
				testMeasurement("c_avg", "s1", "c", nil, 5),
				testMeasurement("b_avg", "s4", "b", nil, 3),
				testMeasurement("a_avg", "s5", "a", nil, 4),
			},
			`station,s1,s4,s5
landuse,me_s1,me_s4,me_s5
latitude,3.14159,3.14159,3.14159
longitude,2.71828,2.71828,2.71828
elevation,1000,1000,1000
parameter,c,b,a
depth,,,
aggregation,avg,avg,avg
unit,c,b,a
2020-01-01 00:15:00,0,0,0
2020-01-01 00:30:00,1,1,1
2020-01-01 00:45:00,2,2,2
2020-01-01 01:00:00,3,NaN,3
2020-01-01 01:15:00,4,NaN,NaN
`,
		},
		"depth": {
			browser.TimeSeries{
				testMeasurement("swc_st_05_1_avg", "s1", "c", browser.Int64(5), 5),
				testMeasurement("swc_st_00_avg", "s4", "b", browser.Int64(0), 3),
				testMeasurement("swc_st_05_avg", "s5", "a", browser.Int64(5), 4),
			},
			`station,s1,s4,s5
landuse,me_s1,me_s4,me_s5
latitude,3.14159,3.14159,3.14159
longitude,2.71828,2.71828,2.71828
elevation,1000,1000,1000
parameter,swc_st_1,swc_st,swc_st
depth,5,0,5
aggregation,avg,avg,avg
unit,c,b,a
2020-01-01 00:15:00,0,0,0
2020-01-01 00:30:00,1,1,1
2020-01-01 00:45:00,2,2,2
2020-01-01 01:00:00,3,NaN,3
2020-01-01 01:15:00,4,NaN,NaN
`,
		},
		"gl_issue_116_sorting": {
			browser.TimeSeries{
				testMeasurement("c_05_avg", "s1", "c", browser.Int64(5), 5),
				testMeasurement("a_avg", "s5", "a", nil, 4),
				testMeasurement("c_03_avg", "s1", "c", browser.Int64(3), 5),
				testMeasurement("b_03_avg", "s4", "b", browser.Int64(3), 2),
				testMeasurement("b_avg", "s4", "e", nil, 3),
				testMeasurement("b_avg", "s1", "b", nil, 3),
				testMeasurement("c_avg", "s2", "a", nil, 3),
			},
			`station,s1,s1,s1,s2,s4,s4,s5
landuse,me_s1,me_s1,me_s1,me_s2,me_s4,me_s4,me_s5
latitude,3.14159,3.14159,3.14159,3.14159,3.14159,3.14159,3.14159
longitude,2.71828,2.71828,2.71828,2.71828,2.71828,2.71828,2.71828
elevation,1000,1000,1000,1000,1000,1000,1000
parameter,b,c,c,c,b,b,a
depth,,5,3,,3,,
aggregation,avg,avg,avg,avg,avg,avg,avg
unit,b,c,c,a,b,e,a
2020-01-01 00:15:00,0,0,0,0,0,0,0
2020-01-01 00:30:00,1,1,1,1,1,1,1
2020-01-01 00:45:00,2,2,2,2,NaN,2,2
2020-01-01 01:00:00,NaN,3,3,NaN,NaN,NaN,3
2020-01-01 01:15:00,NaN,4,4,NaN,NaN,NaN,NaN
`,
		},
	}

	for k, tc := range testCases {
		t.Run(k, func(t *testing.T) {
			var buf bytes.Buffer
			w := NewWriter(&buf)
			w.Write(tc.in)

			diff := cmp.Diff(tc.want, buf.String())
			if diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func testMeasurement(label, station, unit string, depth *int64, n int) *browser.Measurement {
	a := strings.Split(label, "_")
	aggr := a[len(a)-1]

	m := &browser.Measurement{
		Label: label,
		Station: &browser.Station{
			Name:      station,
			Landuse:   "me_" + station,
			Elevation: 1000,
			Latitude:  3.14159,
			Longitude: 2.71828,
		},
		Aggregation: aggr,
		Unit:        unit,
		Depth:       depth,
	}

	ts := time.Date(2020, time.January, 1, 0, 0, 0, 0, browser.Location)

	for i := 0; i < n; i++ {
		ts = ts.Add(15 * time.Minute)
		m.Points = append(m.Points, &browser.Point{
			Timestamp: ts,
			Value:     float64(i),
		})
	}

	return m
}
