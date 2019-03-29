package models

type Skill struct {
	Items     []Items     `json:"items"`
	Languages []Languages `json:"languages"`
	Tags      []string    `json:"tags"`
}
type Items struct {
	Title string `json:"title"`
	Stars int    `json:"stars"`
}
type Languages struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (s *Skill) GetAll() *Skill {
	return &Skill{
		Items:     s.GetItems(),
		Languages: s.GetLanguages(),
		Tags:      s.GetTags(),
	}
}

func (s *Skill) GetItems() []Items {
	return []Items{
		{
			Title: "Front-End",
			Stars: 10,
		},
		{
			Title: "Back-End",
			Stars: 10,
		},
		{
			Title: "Desktop",
			Stars: 8,
		},
		{
			Title: "Mobile",
			Stars: 7,
		},
		{
			Title: "DevOps/Infrastructure",
			Stars: 7,
		},
		{
			Title: "Software Architecture",
			Stars: 8,
		},
		{
			Title: "Data Structures",
			Stars: 8,
		},
		{
			Title: "Game Development",
			Stars: 8,
		},
		{
			Title: "Design",
			Stars: 6,
		},
	}
}

func (s *Skill) GetLanguages() []Languages {
	return []Languages{
		{
			Title:       "Portuguese (Native)",
			Description: "Fluent",
		},
		{
			Title:       "English",
			Description: "Advanced",
		},
		{
			Title:       "Spanish",
			Description: "Low Intermediary",
		},
	}
}

func (s *Skill) GetTags() []string {
	return []string{
		"FullStack",
		"Software",
		"Developer",
		"Portuguese",
		"English",
		"Spanish",
		"C++",
		"C#",
		"Java",
		"Go",
		"Javascript",
		"Dart",
		"Front-end",
		"React",
		"Vue",
		"AngularJS",
		"jQuery",
		"Razor",
		"Css",
		"JSS",
		"Less",
		"Sass",
		"Bootstrap",
		"Material-UI",
		"ant Design",
		"MDBootstrap",
		"Ionic",
		"HTML",
		"Telerik",
		"KendoUI",
		"DevExpress",
		"SyncFusion",
		"PrimeFaces",
		"RichFaces",
		"Adobe Flash",
		"Back-end",
		"Asp.Net Web Forms",
		"Asp.Net MVC",
		"Asp.Net Core",
		"JBoss Seam",
		"Java Spring",
		"node.js",
		"Asp Classic",
		"WCF",
		"Desktop",
		"Universal Apps",
		"WPF",
		"VB.Net",
		"Mobile",
		"Flutter",
		"Android",
		"Cordova",
		"Ionic",
		"Xamarin",
		"Windows Phone",
		"Architecture",
		"MVVM",
		"MVC",
		"DDD",
		"TDD",
		"OOP",
		"Design Patterns",
		"Clean Code",
		"Microservices",
		"SOAP",
		"Rest",
		"E-Commerce",
		"UML",
		"DevOps",
		"Infrastructure",
		"Azure DevOps (Visual Studio Online)",
		"Windows",
		"Mac",
		"Linux",
		"Docker",
		"Azure",
		"Aws",
		"Google Cloud Platform",
		"JBoss",
		"Apache Tomcat",
		"IIS",
		"Continuous Integration",
		"GIt",
		"SVN",
		"Mercurial",
		"Source Code",
		"TFS",
		"GitHub",
		"Data Structures",
		"MS Sql Server",
		"Oracle 11g",
		"MySql",
		"Redis",
		"Mongo",
		"Management",
		"Azure DevOps (Visual Studio Online)",
		"Trello",
		"Jira",
		"GitHub",
		"ZenHub",
		"MS Project",
		"Excel",
		"Game Developer",
		"Unity 2D/3D",
		"Maya",
		"3D Studio Max",
		"Design",
		"Photoshop",
		"InDesign",
		"PageMaker",
		"Corel Draw",
		"Illustrator",
		"Personal",
		"Easy to work with",
		"Easy learning",
		"Contributing member",
		"Transparent",
		"Proactive",
		"Easy adaptation",
		"Excellent interpersonal skills",
	}
}
