package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_geoTrack(t *testing.T) {
	app := &goBlog{
		cfg: &config{
			Db: &configDb{
				File: filepath.Join(t.TempDir(), "test.db"),
			},
			Server: &configServer{},
			Blogs: map[string]*configBlog{
				"en": {
					Lang: "en",
				},
				"de": {
					Lang: "de",
				},
			},
		},
	}

	_ = app.initDatabase(false)
	app.initComponents(false)

	// First test (just with track)

	gpxBytes, _ := os.ReadFile("testdata/test.gpx")

	p := &post{
		Blog: "en",
		Parameters: map[string][]string{
			"gpx": {
				string(gpxBytes),
			},
		},
	}

	resEn, err := app.getTrack(p)
	require.NoError(t, err)

	assert.True(t, resEn.HasPoints)
	assert.NotEmpty(t, resEn.Paths)
	assert.Empty(t, resEn.Points)
	assert.Equal(t, "2.70", resEn.Kilometers)
	assert.Equal(t, "0:42:53", resEn.Hours)

	p.Blog = "de"

	resDe, err := app.getTrack(p)
	require.NoError(t, err)

	assert.True(t, resDe.HasPoints)
	assert.NotEmpty(t, resDe.Paths)
	assert.Empty(t, resDe.Points)
	assert.Equal(t, "2,70", resDe.Kilometers)
	assert.Equal(t, "0:42:53", resDe.Hours)

	// Second file (with track and waypoint)

	gpxBytes, _ = os.ReadFile("testdata/test2.gpx")

	p = &post{
		Blog: "en",
		Parameters: map[string][]string{
			"gpx": {
				string(gpxBytes),
			},
		},
	}

	resEn, err = app.getTrack(p)
	require.NoError(t, err)

	assert.True(t, resEn.HasPoints)
	assert.NotEmpty(t, resEn.Paths)
	assert.NotEmpty(t, resEn.Points)
	assert.Equal(t, "0.08", resEn.Kilometers)
	assert.Equal(t, "0:01:29", resEn.Hours)

	// Third file (just with route)

	gpxBytes, _ = os.ReadFile("testdata/test3.gpx")

	p = &post{
		Blog: "en",
		Parameters: map[string][]string{
			"gpx": {
				string(gpxBytes),
			},
		},
	}

	resEn, err = app.getTrack(p)
	require.NoError(t, err)

	assert.True(t, resEn.HasPoints)
	assert.NotEmpty(t, resEn.Paths)
	assert.Empty(t, resEn.Points)
	assert.Equal(t, "", resEn.Kilometers)
	assert.Equal(t, "", resEn.Hours)

}