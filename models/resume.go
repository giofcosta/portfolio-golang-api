package models

//Resume defines the model for resume informations
type Resume struct {
	Experiences    []Experiences    `json:"experiences"`
	Education      []Education      `json:"education"`
	Certifications []Certifications `json:"certifications"`
	Readings       []Readings       `json:"readings"`
}

type Experiences struct {
	Title       string `json:"title"`
	Date        string `json:"date"`
	Rule        string `json:"rule"`
	Description string `json:"description"`
}
type Education struct {
	Title string `json:"title"`
	Rule  string `json:"rule"`
	Date  string `json:"date"`
}
type Certifications struct {
	Title       string `json:"title"`
	Rule        string `json:"rule"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
type Readings struct {
	Title       string `json:"title"`
	Rule        string `json:"rule"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

func (r *Resume) GetAll() *Resume {
	return &Resume{
		Experiences:    r.GetExperiences(),
		Education:      r.GetEducation(),
		Certifications: r.GetCertifications(),
		Readings:       r.GetReadings(),
	}
}

func (r *Resume) GetExperiences() []Experiences {
	return []Experiences{
		{
			Title: "After Mercado Livre – Florianópolis, SC - Brazil",
			Date:  "2018 - 2019",
			Rule:  "Senior Software Engineer, Indie Game and Web developer, Freelancer",
			Description: `<p>
				I started to improve my english, hiring a teacher and practicing a
				lot, reaching advanced level. <br />
				I improved my knowledge of React/Vue, nodejs, JSS during this time,
				starting to make my portfolio website. <br />
				I was able to learn Flutter, returning to mobile development, making
				an app. <br />I improved my knowledge of Unity3D skills, starting my
				first indie game. <br />I improved my knowledge of DevOps, creating a
				continuous integrations for my website portfolio in GCP and AWS and
				Github using Docker and Kubernetes.
				<br />
				<br />
				<br />
			  </p>`,
		},
		{
			Title: "Mercado Livre – Florianópolis, SC - Brazil",
			Date:  "2017 - 2018",
			Rule:  "Senior Software Engineer",
			Description: `<p>
				I Participated in a project as a front-end using React, developing
				some components importing some csv information to TMS. <br />
				Development of the new version of the TMS, using Golang in the
				Back-End Company’s Infrastructure and React in the front-end. <br />
				Automated Tests in golang, and React with Jest. <br />
				Deploy in AWS through Fury (their infrastructure wrapper) with
				continuous integration. <br />
				Docker containers and Microservice Architecture in Linux and Mac OS.
				<br />
				Git, Git-flow, GitHub and Jira for Management. <br />
				Member of Committee and Community manager in a project called IT
				Brasil, involving some Senior developers in MeLi Brazil to share
				knowledge and keep the growth of the brazilian community. <br />
				Communication by speaking and writing in Spanish language. <br />
				Slack, Discord and Hangout for remote communication.
			  </p>`,
		},
		{
			Title: "Axado – Florianópolis, SC - Brazil",
			Date:  "2016 - 2017",
			Rule:  "Senior System Analyst",
			Description: `<p>
				Development and maintenance in a tracking system, using Asp.Net
				MVC/.Net Core, C#, WPF and Windows Service, making an integration
				between the site, the desktop services and Zebra and HP printer
				devices (some IoT skills). <br />
				JQuery and AngularJS in Old Applications and Vue.js for a new
				Application. <br />
				All Deploys in Azure in a Distributed Microservices Architecture,
				using also Topics and Queues in Bus Service and Table Storage for the
				Logs. SignalR and AppInsights is used for real time analizes in a
				Dashboard.
				<br />
				SQL Server database with Entity Framework/EFCore, sometimes Dapper for
				persistency. <br />
				Visual Studio Online for management and scrum. <br />
				Git as version control <br />
				The company had buyed by Mercado Livre in 2017
			  </p>`,
		},
		{
			Title: "Black Drake IT Solutions – Juiz de Fora, MG - Brazil",
			Date:  "2014 - 2016",
			Rule:  "Founder/CEO/CTO",
			Description: `<p>
				Administration, Development, Design, Game Development, Project
				Management and Consulting. I made a lot of projects for my clients,
				using technologies like c#, Ap.net MVC, SignalR, Java, Javascript,
				Android, WIndows Phone, PhoneGap, Windows Universal Apps, Xamarin,
				Unity3D. <br />
				Telerik, AngularJS, Ionic, SEO-Friendly and W3C validations, resource
				optimization, performance improvements. <br />
				Microservices and Web Robots (WebCrawler) with Watin and Selenium.
				<br />
				Git as version control and Visual Studio Online for Project Management
				With Scrum. <br />
				SQL Server and Oracle for Database. <br />
				DevOps on Azure.
			  </p>`,
		},
		{
			Title: "Guiando Telecom – Juiz de Fora, MG - Brazil",
			Date:  "2013 - 2014",
			Rule:  "Senior System Analyst",
			Description: `<p>
				Development and maintenance of .Net C# WebForm Applications in Thomson
				Reuters as outsourced. <br />
				Our team resolved hundreds of bugs a week in their applications.
			  </p>`,
		},
		{
			Title: "Nova Consulting – Juiz de Fora, MG - Brazil",
			Date:  "2013",
			Rule:  "Senior System Analyst",
			Description: `<p>
				Development and maintenance of Java applications using Java-ee, JSF,
				Jmms, ejb, JBoss Richfaces. <br />
				Versioning With SVN. <br />
				Oracle and MSSQL with JPA/Hibernate for persistence. <br />
				We had our own data centers. <br />
			  </p>`,
		},
		{
			Title: "Affero – Juiz de Fora, MG - Brazil",
			Date:  "2012 – 2013",
			Rule:  "Senior System Analyst",
			Description: `<p>
				Development and Maintenance of Java applications using JSP, Spring,
				Java-ee, CDI, JSF, Jms, ejb, Jax-rs, Rest-easy. <br />
				We Developed a proprietary reporting server in AWS infrastructure.{" "}
				<br />
				Javascript, Primefaces, jQuery, CSS and HTML in advanced levels.{" "}
				<br />
				SQL Server and JPA/Hibernate for persistence. <br />
				TDD using tools like Arquillian and Mockito. <br />
				Continuous Integrations with jenkins and Sonar. <br />
				SEO-Friendly and W3C validations, resource optimization, performance
				improvements. <br />
				Scrum Agile Methodology. <br />
				All published in AWS in a custom JBoss cluster.
			  </p>`,
		},
		{
			Title: "Delage Consulting – Juiz de Fora, MG - Brazil",
			Date:  "2011 – 2012",
			Rule:  "Senior System Analyst",
			Description: `<p>
				Development and support of a complex Quotation System and start of the
				remade his Warehouse Software using c# and asp.net MVC3. <br />
				Telerik, Javascript, JQuery, CSS and HTML in advanced levels. <br />
				SQL Server with NHibernate for persistence. <b />
				SEO-Friendly and W3C validations, resource optimization, performance
				improvements. <br />
				Scrum and Kanban methodologies. Mercurial and TFS for version control.{" "}
				<br />
				All published on Windows Azure.
			  </p>`,
		},
		{
			Title: "Novaprolink Technology– Juiz de Fora, MG - Brazil",
			Date:  "2008 – 2011",
			Rule:  "Senior Developer - Team Leader",
			Description: `<p>
				Support and development in Classic ASP. <br />
				Development of Company Website, CMS, E-Commerce and Framework in C#
				and .Net WebForms and MVC. <br />
				DevExpress, Telerik, Javascript, JQuery, CSS, HTML in advanced levels.
				<br />
				MySQL and SQL Server for data. <br />
				SVN and TFS for version control. <br />
				SEO-Friendly and W3C validations, resource optimization, performance
				improvements. <br />
				We had our own data centers. <br />
				The company was bought by Thomson Reuters in 2012.
			  </p>`,
		},
	}
}

func (r *Resume) GetEducation() []Education {
	return []Education{
		{
			Title: "Salgado de Oliveira University – Juiz de Fora, MG Brasil",
			Rule:  "System Analyst - Degree - Incomplete",
			Date:  "2008",
		},
		{
			Title: "Windows Azure Course – Belo Horizonte, MG Brasil",
			Rule:  "Microsoft/PUC Minas",
			Date:  "2011",
		},
		{
			Title: "Scrum Practical Course – Juiz de Fora, MG Brasil",
			Rule:  "Bridge Consulting/Delage Consulting",
			Date:  "2011",
		},
		{
			Title: "OOP Course – Juiz de Fora, MG Brasil",
			Rule:  "Novaprolink Technology/Salgado de Oliveira University",
			Date:  "2008",
		},
		{
			Title: "Web Development Course – Juiz de Fora, MG Brasil",
			Rule:  "Salgado de Oliveira University",
			Date:  "2008",
		},
		{
			Title: "Basic Computing Course – Juiz de Fora, MG Brasil",
			Rule:  "SOS computers",
			Date:  "2002",
		},
	}
}

func (r *Resume) GetCertifications() []Certifications {
	return []Certifications{
		{
			Title:       "Salgado de Oliveira University",
			Rule:        "Android Developer Course",
			Description: "Two days course for android development at Salgado de Oliveira University.",
			Date:        "2010",
		},
		{
			Title: "Microsoft Certified",
			Rule:  "MCP - Web Developer",
			Description: `<p>
					MCPD – Design and Developing Web Applications using Microsoft .NET{" "}
					<br />
					MCTS - .NET Framework 4, Data Access <br />
					MCTS - .NET Framework 4, Service Communication Applications <br />
					MCTS - .NET Framework 4, Web Applications
		 		 </p>`,
			Date: "2010",
		},
		{
			Title:       "The Developers Conference - Vue.js - SC/SP",
			Rule:        "Fast and Efficient development with Vue.js.",
			Description: "Fifty minutes lecture about Vue.js for São Paulo and Florianópolis for one the largest technology events held in Brazil",
			Date:        "2017",
		},
		{
			Title:       "Udemy",
			Rule:        "Vue JS 2 - The Complete Guide (incl. Vue Router & Vuex)",
			Description: "By Maximillian Shwarzmuller, Professional web developer and Instructor. A complete course about Vue JS in the web.",
			Date:        "2017",
		},
		{
			Title:       "Udemy",
			Rule:        "React + Redux",
			Description: "By Leonardo Moura Leitão, Software Architect. React, Redux, Webpack, Redux-Form, MongoDB, Express, Node... Several practical exercises and two complete apps from zero!",
			Date:        "2018",
		},
		{
			Title:       "Udemy",
			Rule:        "The Ultimate Game Development Guide",
			Description: "By Jonathan Weinberger, Authorized Unity Instructor. Create in partnership with Unity Technologies, i was able to improve my C# experience in 2D and 3D game with this embracing guide.",
			Date:        "2018",
		},
	}
}

func (r *Resume) GetReadings() []Readings {
	return []Readings{
		{
			Title:       "Head First Java",
			Rule:        "Kathy Sierra, Bert Bates",
			Description: "",
			Date:        "",
		},
		{
			Title:       "Head First Iphone",
			Rule:        "Dan Pilone, Tracey Pilone",
			Description: "",
			Date:        "",
		},
		{
			Title:       "Google Android",
			Rule:        "Ricardo R. Lecheta",
			Description: "",
			Date:        "",
		},
		{
			Title:       "System Architecture for Web With Java Using Design Patterns and Frameworks",
			Rule:        "Evandro Carlos Teruel",
			Description: "",
			Date:        "",
		},
		{
			Title:       "The Head Of Steve Jobs",
			Rule:        "Leander Kahney",
			Description: "",
			Date:        "",
		},
		{
			Title:       "The Accidental Billionaires - The Founding of Facebook",
			Rule:        "Ben Mezrich",
			Description: "",
			Date:        "",
		},
		{
			Title:       "The Art  of War",
			Rule:        "Sun Tsu",
			Description: "",
			Date:        "",
		},
	}
}
