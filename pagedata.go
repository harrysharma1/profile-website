package main

import "html/template"

var Hello = "Listening at localhost:6969 to close click CTRL+C"

type IndexWelcomeBox struct {
	Title      string
	Paragraphs []template.HTML
}

type IndexFlake8Box struct {
	Title      string
	Paragraphs []template.HTML
	Lists      []template.HTML
}

type IndexSpackBox struct {
	Title      string
	Paragraphs []template.HTML
	Lists      []template.HTML
}

type GoShellBox struct {
	Title      string
	Paragraphs []template.HTML
}

type GoodReadsBox struct {
	Style   []template.CSS
	Content []template.HTML
}

type IndexData struct {
	Welcome   IndexWelcomeBox
	Flake8    IndexFlake8Box
	Spack     IndexSpackBox
	GoShell   GoShellBox
	GoodReads GoodReadsBox
}

var index_data IndexData = IndexData{
	Welcome: IndexWelcomeBox{
		Title: "Welcome to my page!",
		Paragraphs: []template.HTML{
			`
I am currently doing my Undergrad at <a href="https://www.kcl.ac.uk" target="_blank"
class="hover:underline hover:text-gray-300">King's College London</a>, where I am doing an industry placement within the University under the <a href="https://www.kcl.ac.uk/research/facilities/e-research" target="_blank" class="hover:underline hover:text-gray-300">e-Research department (under RMID)</a>		
`,
			`This website will go over some of the projects I have taken part in, both in and out of my time with e-Research.`,
		},
	},

	Flake8: IndexFlake8Box{
		Title: "Flake8 CLI in different languages",
		Paragraphs: []template.HTML{
			`Here are some CLI implementation calling on basic API endpoints. The main focus was trying to use mainly the languages standard library.`,
			`I did go a bit over the top with python when it came to styling terminal output.`,
			`The best of the three to program in had to be, in my opinion, Golang.`,
		},
		Lists: []template.HTML{
			` <a href="https://github.com/harrysharma1/goflake8api"
							target="_blank"
							class="hover:underline hover:text-gray-300">⚪ Implementation in
							Golang</a>
			`,
			`
			<a href="https://github.com/harrysharma1/rust-flake8api"
							target="_blank"
							class="hover:underline hover:text-gray-300">⚪ Implementation in
							Rust</a>
			`,
			`
			<a href="https://github.com/harrysharma1/flake8api"
							target="_blank"
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
							target="_blank"
							class="hover:underline hover:text-gray-300">⚪
							Vbz-compression</a>`,
			`<a href="https://github.com/spack/spack/pull/41967"
							target="_blank"
							class="hover:underline hover:text-gray-300">⚪ Glow</a>`,
			`<a href="https://github.com/spack/spack/pull/41988"
							target="_blank"
							class="hover:underline hover:text-gray-300">⚪ Dorado</a>`,
			`<a href="https://github.com/spack/spack/pull/44339"
							target="_blank"
						        class="hover:underline hover:text-gray-300">⚪ Metacarpa </a>`,
		},
	},

	GoShell: GoShellBox{
		Title: "Go Shell",
		Paragraphs: []template.HTML{
			`I have been working on writing my own shell implementation using Golang and its standard library`,
			`You can find the implementation <a href="https://github.com/harrysharma1/go-shell" class="hover:underline hover:text-gray-300">here</a>`,
			`The inspiration for this was mainly to improve my skills witha golang and to understand how shelss such as zsh, csh, bash, etc. work`,
		},
	},

	GoodReads: GoodReadsBox{
		Style: []template.CSS{
			`		
			#customize-list{
			float:left;
			margin-left:20px;
			list-style:none;
			}
			`,
			`
			#gr_updates_widget{
				float:left;
				border-radius: 5px;
				background-color:#fff;
				border:solid #82e822 10px;
				-webkit-box-shadow: 0px 0px 4px 1px #70fa07,
				inset 0px 0px 0px 1px #eedb0d;
				-moz-box-shadow: 0px 0px 4px 1px #595959,
				inset 0px 0px 0px 1px #7D730B;
				box-shadow: 0px 0px 4px 1px #595959,
				inset 0px 0px 0px 1px #7D730B;
				padding:15px 0 0px 15px;
				width:400px;
				height:365px;
			}
			`,
			`
			#gr_footer{
				margin-bottom:0px;
				height:30px;
			}
			`,
			`
			#gr_updates_widget p{
				padding:0px;
				margin:0;
				font-size:14px;
			}
			`,
			`
			#gr_footer img{
				width:100px;
				float:left;
			}
			`,
		},
		Content: []template.HTML{
			`
			<div id="gr_updates_widget">
				<iframe sandbox id="the_iframe" src="https://goodreads.com/widgets/user_update_widget?height=250&num_updates=999&user=184799979&width=400" width="398" height="180" frameborder="0"></iframe>
					<div id="gr_footer">
						<a href="https://www.goodreads.com/"><img alt="Goodreads: Book reviews, recommendations, and discussion" src="https://s.gr-assets.com/images/layout/goodreads_logo_140.png" /></a>
					</div>
			</div>
			`,
		},
	},
}

type Timeline struct {
	Date  string
	Title string
	Text  template.HTML
}

type AboutData struct {
	Timeline []Timeline
}

var about_data = AboutData{
	Timeline: []Timeline{
		{
			Date:  "June 2019",
			Title: "Finished GCSE",
			Text: `
			<p> I finished my GCSEs in 2019 and received the following grades:</p>
			<br>
			<ul>
				<li>• Geography: 9</li>
				<li>• History: 8</li>
				<li>• Maths: 8</li>
				<li>• Science: 8/7</li>
				<li>• English Literature: 7</li>
				<li>• English Language: 6</li>
				<li>• Computer Science: 6</li>
				<li>• Business Studies: LM2</li>
			</ul>
			<br>
                        <p><strong>Note to anyone outside of the UK:</strong> General Certificate of Secondary Education Ranking was switched to a 1-9 system a year or two before I took my exams (anything less than a 4 is a fail). </p>
			<p> It was in GCSE Computer Science where I had started using both <strong>Java</strong> and <strong>Python</strong>. </p>
			<p> I learnt the fundementals that were specified by OCR</p>
			<p> You can find the the the specification that OCR set <a href="https://www.ocr.org.uk/images/225975-specification-accredited-gcse-computer-science-j276.pdf" target="_blank" class="hover:underline hover:text-gray-300">here</a> </p>
			<p> It is split into two sections and a 20 Hours Project: </p>
			<br>
			<ol>
				<li>1. <strong>Computer systems:</strong> My favourite topics were: Wired and wireless networks, Network topologies, protocols and layers, System security, System software</li>
				<li>2. <strong>Computational thinking, algorithms and programming:</strong> My favourite topics were: Data representation, Algorithms, Programming techniques</li>
				<li>3. <strong>Programming Project:</strong> My project consisted of a phsysics simulaton, written in Java, that mapped projectile motion. Sadly I had not used any version control system and the computer I programmed it in is bricked so this has been lost to time.
			</ol>
			`,
		},
		{
			Date:  "June 2021",
			Title: "Finished A-Levels",
			Text: `
			<p> I completed my A-levels in Covid so  I won't go too much into detail but my grades were:</p>
			<br>
			<ol>
			<li>1. <strong>Maths:</strong> A </li>
			<li>2. <strong>Computer Science:</strong> B</li>
			<li>3. <strong>Chemistry:</strong> B </li> 
			</ol>
			`,
		},
		{
			Date:  "September 2021",
			Title: "Start University in King's College London",
			Text: `
			<p>Started my first year of university and the first year out of thw Pandemic. </p>
			<p>Adjusting to the workload was difficult at first and I even failed a module or two. </p>
			<p>This became an eye opener for me as I had become complacent with my studies. </p>
			<p>The modules in first year were all mandatory. They consisted of:</p>
			<br>
			<ul>
				<li>• <strong>Elementary Logic with Applications</strong></li>
				<li>• <strong>Introduction to Software Engineering</strong></li>
				<li>• <strong>Database Systems</strong></li>
				<li>• <strong>Data Structures</strong></li>
				<li>• <strong>Foundations of Computing 1</strong></li>
				<li>• <strong>Computer Systems</strong></li>
				<li>• <strong>Programming Practice & Applications</strong></li>
				<li>• <strong>Introduction to Professional Practice</strong></li>
			</ul>
			`,
		},
		{
			Date:  "September 2022",
			Title: "Finish 2nd Year Exams",
			Text: `
			<p> I came into second year with a renewed vigor. </p>
			<p> The modules I selected were: </p>
			<br>
			<ul>
				<li>• <strong>Software Engineering Group Project</strong></li>
				<li>• <strong>Operating Systems & Concurrency</strong></li>
				<li>• <strong>Foundations of Computing 2</strong></li>
				<li>• <strong>Introduction to Artificial Intelligence</strong></li>
				<li>• <strong>Practical Experiences of Programming</strong></li>
				<li>• <strong>Internet Systems</strong></li>
				<li>• <strong>Programming Language Design Paradigms</strong></li>
			</ul>
			<br>
			<p>Along with the modules I needed to apply for placement positions.</p>
			<p>After countless rejection in various stages of the application phase, I managed to secure two offers.</p>
			<p>I took the offer to work with e-Research, which is a sub-department of the Research Management and Innovation Division (RMID) within King's College London.</p>
			`,
		},
		{
			Date:  "September 2023",
			Title: "Start Industry Placement in e-Research",
			Text: `
			<p>My start as an Operations Enginerr within e-Research was hectic and interesting.</p>
			<p>My day to day responisbility included: </p>
			<br>
			<ul>
				<li>• Monitoring HPC health (with <code>htop</code>, <code>slurm</code> and other monitoring software), alonside other services provided by e-Research.</li>
				<li>• Troubleshoot user queries.</li>
				<li>• Adding software to HPC, through Spack. </li>
				<li>• Data migrations using <code> rsync</code>. </li>
				<li>• Helping with in-person training events. </li>
				<li>• Developing HPC related skills during Personal Devlopment Time, alongside a Slackbot and Spack contributions. </li>
			</ul>
			<br>	
			<p>I also got to go to <a href="https://www.scd.stfc.ac.uk/Pages/CIUK2023.aspx">CIUK 2023</a> and <a href="https://archive.fosdem.org/2024/">FOSDEM 2024</a>, both of which were amazing experiences.</p>
			<p>This placement has helped maked meanignful connections within the industry alonside a new found appreciation towards Open Source Software (OSS).</p>
			`,
		},
	},
}

type ContactDetails struct {
	Email   string
	Subject string
	Body    string
}
