// Copyright 2015 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package markdown

import (
	"io/ioutil"
	"testing"

	blackfriday2 "github.com/russross/blackfriday/v2"
)

func BenchmarkRenderSpecNoHTML(b *testing.B) {
	b.StopTimer()
	data, err := ioutil.ReadFile("spec/spec-0.28.txt")
	if err != nil {
		b.Fatal(err)
	}

	md := New(HTML(false), XHTMLOutput(true))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		md.RenderToString(data)
	}
}

func BenchmarkRenderSpec(b *testing.B) {
	b.StopTimer()
	data, err := ioutil.ReadFile("spec/spec-0.28.txt")
	if err != nil {
		b.Fatal(err)
	}

	md := New(HTML(true), XHTMLOutput(true))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		md.RenderToString(data)
	}
}

func BenchmarkRenderSpecBlackFriday2(b *testing.B) {
	b.StopTimer()
	data, err := ioutil.ReadFile("spec/spec-0.28.txt")
	if err != nil {
		panic(err)
	}

	renderer := blackfriday2.NewHTMLRenderer(
		blackfriday2.HTMLRendererParameters{
			Flags: blackfriday2.Smartypants |
				blackfriday2.SmartypantsDashes |
				blackfriday2.SmartypantsLatexDashes |
				blackfriday2.SmartypantsAngledQuotes,
		},
	)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		blackfriday2.Run(data, blackfriday2.WithExtensions(
			blackfriday2.NoIntraEmphasis|
				blackfriday2.Tables|
				blackfriday2.FencedCode|
				blackfriday2.Autolink|
				blackfriday2.Strikethrough),
			blackfriday2.WithRenderer(renderer),
		)
	}
}
