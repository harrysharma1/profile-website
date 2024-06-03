package main

import "html/template"

var Hello = "Listening at localhost:6969 to close click CTRL+C"

type IndexWelcomeBox struct{
	Title string
	Paragraphs []template.HTML
}

type IndexFlake8Box struct{
	Title string
	Paragraphs []template.HTML
	Lists []template.HTML
}

type IndexSpackBox struct{
	Title string
	Paragraphs []template.HTML
	Lists []template.HTML
}


type IndexData struct{
	Welcome IndexWelcomeBox
	Flake8 IndexFlake8Box
	Spack IndexSpackBox
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

	Flake8: IndexFlake8Box{
		Title: "Flake8 CLI in different languages",
		Paragraphs: [] template.HTML{
			`Here are some CLI implementation calling on basic API endpoints. The main focus was trying to use mainly the languages standard library.`,
			`I did go a bit over the top with python when it came to styling terminal output.`,
			`The best of the three to program in had to be, in my opinion, Golang.`,
		},
		Lists: [] template.HTML{
			` <a href="https://github.com/harrysharma1/goflake8api"
							class="hover:underline hover:text-gray-300">⚪ Implementation in
							Golang</a>
			`,
			`
			<a href="https://github.com/harrysharma1/rust-flake8api"
							class="hover:underline hover:text-gray-300">⚪ Implementation in
							Rust</a>
			`,
			`
			<a href="https://github.com/harrysharma1/flake8api"
							class="hover:underline hover:text-gray-300">⚪ Implementation in
							Python</a>
			`,
		},
	},

	Spack: IndexSpackBox{
		Title: "Spack Open Source Contributions",
		Paragraphs: []template.HTML{
			`Spack is an open source package manager, written in Python, used to build and run software on HPC's.`,
			`I have made some software contributions for Spack. The next goal would be to contribute to the actual Spack build systems.`,
		},
		Lists: []template.HTML{
			`<a href="https://github.com/spack/spack/pull/41714"
							class="hover:underline hover:text-gray-300">⚪
							Vbz-compression</a>`,
			`<a href="https://github.com/spack/spack/pull/41967"
							class="hover:underline hover:text-gray-300">⚪ Glow</a>`,
			`<a href="https://github.com/spack/spack/pull/41988"
							class="hover:underline hover:text-gray-300">⚪ Dorado</a>`,
			`<a href="https://github.com/spack/spack/pull/44339"
						        class="hover:underline hover:text-gray-300">⚪ Metacarpa </a>`,
		},
	},
}
