package main

import "html/template"

var Hello = "Listening at localhost:6969 to close click CTRL+C"

type IndexWelcomeBox struct{
	Title string
	Paragraphs []template.HTML
}



type IndexData struct{
	Welcome IndexWelcomeBox
}

var index_data IndexData=IndexData{
	Welcome: IndexWelcomeBox{
		Title: "Welcome to my page!",
		Paragraphs: [] template.HTML{
		`
I am currently doing my Undergrad at <a href="https://www.kcl.ac.uk"
class="hover:underline hover:text-gray-300">King's College London</a>, where I am doing an industry placement within the University under the <a href="https://www.kcl.ac.uk/research/facilities/e-research" class="hover:underline hover:text-gray-300">e-Research department (under RMID)</a>		
`,
`This website will go over some of the projects I have taken part in, both in and out of my time with e-Research.`,
		},
	},
}
