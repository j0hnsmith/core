// Copyright 2023 The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"time"

	"goki.dev/events"
	"goki.dev/grows/images"
)

// AssertRender is a helper function that makes a new window from
// the scene, waits until it is shown, calls [Scene.AssertPixels]
// with the given values, and then closes the window.
// It does not return until all of those steps are completed.
// If a function is passed for the final argument, it is called after the
// scene is shown, right before [Scene.AssertPixels] is called. Also,
// if a function is passed, [Scene.DoNeedsRender] is also called before
// [Scene.AssertPixels].
func (sc *Scene) AssertRender(t images.TestingT, filename string, fun ...func()) {
	showed := make(chan struct{})
	sc.OnShow(func(e events.Event) {
		if len(fun) > 0 {
			fun[0]()
			sc.DoNeedsRender()
		}
		showed <- struct{}{}
	})
	sc.NewWindow().Run()
	<-showed
	time.Sleep(50 * time.Millisecond)
	sc.AssertPixels(t, filename)
	sc.Close()
}

// AssertPixels asserts that [Scene.Pixels] is equivalent
// to the image stored at the given filename in the testdata directory,
// with ".png" added to the filename if there is no extension
// (eg: "button" becomes "testdata/button.png").
// If it is not, it fails the test with an error, but continues its
// execution. If there is no image at the given filename in the testdata
// directory, it creates the image.
func (sc *Scene) AssertPixels(t images.TestingT, filename string) {
	images.Assert(t, sc.Pixels, filename)
}
