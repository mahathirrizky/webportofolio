package database

import (
	"github.com/mahathirrizky/go-vue/models"
	"gorm.io/gorm"
)

func Seed() {
	contents := []models.Content{
		// HomePage
		{PageName: "home", Section: "hero", Key: "subtitle", Value: "FullStack Developer"},
		{PageName: "home", Section: "hero", Key: "greetings", Value: "Hello I'm"},
		{PageName: "home", Section: "hero", Key: "name", Value: "Person Name"},
		{PageName: "home", Section: "nav", Key: "brand_name", Value: "Person Name"},
		{PageName: "home", Section: "hero", Key: "description", Value: "I excel at crafting elegant digital experiences and I am proficient in various programming languages and technologies."},

		{PageName: "home", Section: "hero", Key: "cv_download_url", Value: "/resume.pdf"}, // Example URL
		{PageName: "home", Section: "hero", Key: "photo_url", Value: "/static/pp.png"},
		{PageName: "home", Section: "social", Key: "items", Value: `[
			{"icon": "mdi-github", "url": "https://github.com/yourusername"},
			{"icon": "mdi-linkedin", "url": "https://linkedin.com/in/yourusername"},
			{"icon": "mdi-twitter", "url": "https://twitter.com/yourusername"}
		]`},
		{PageName: "home", Section: "stats", Key: "num_years_experience", Value: "12"},
		{PageName: "home", Section: "stats", Key: "num_projects_completed", Value: "26"},
		{PageName: "home", Section: "stats", Key: "num_technologies_mastered", Value: "8"},
		{PageName: "home", Section: "stats", Key: "num_code_commits", Value: "200"},

		// ServicePage
		{PageName: "services", Section: "list", Key: "items", Value: `[
			{"num":"01","title":"Web Development","description":"Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos","href":""},
			{"num":"02","title":"UI/UX Design","description":"Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos","href":""},
			{"num":"03","title":"Logo Design","description":"Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos","href":""},
			{"num":"04","title":"SEO","description":"Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos","href":""}
		]`},

		// ResumePage
		{PageName: "resume", Section: "experience", Key: "title", Value: "My Experience"},
		{PageName: "resume", Section: "experience", Key: "description", Value: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos"},
		{PageName: "resume", Section: "experience", Key: "items", Value: `[
			{"company":"Tech Solutions inc","position":"Full Stack Developer","duration":"2022 - Present"},
			{"company":"Web Design Studio","position":"Back-End Developer Freelance","duration":"april 2021"},
			{"company":"Telkomsel","position":"Front-End Developer Freelance","duration":"2020 - 2021"},
			{"company":"Tech Academy","position":"Teaching Assistent","duration":"juni 2020"}
		]`},
		{PageName: "resume", Section: "formaleducation", Key: "title", Value: "My Formal Education"},
		{PageName: "resume", Section: "formaleducation", Key: "description", Value: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos"},
		{PageName: "resume", Section: "formaleducation", Key: "items", Value: `[
			{"institution":"Fakultas Teknik Elektro Unhas","degree":"Magister Teknik Elektro","duration":"2016 - 2018"},
			{"institution":"STMIK BALIKPAPAN","degree":"Sarjana Teknik Informatika","duration":"2011- 2015"}
		]`},
		{PageName: "resume", Section: "informaleducation", Key: "title", Value: "My Informal Education"},
		{PageName: "resume", Section: "informaleducation", Key: "description", Value: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos"},
		{PageName: "resume", Section: "informaleducation", Key: "items", Value: `[
			{"institution":"Online Course Platform","degree":"Full Stack Web Development Bootcamp","duration":"2023"},
			{"institution":"Codecademy","degree":"Front-End Track","duration":"2022"},
			{"institution":"Online Course","degree":"Programing Course","duration":"2020-2021"},
			{"institution":"TEch Institute","degree":"Certivied Web Developer","duration":"2019"}
		]`},
		{PageName: "resume", Section: "skills", Key: "title", Value: "My Skills"},
		{PageName: "resume", Section: "skills", Key: "description", Value: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos"},
		{PageName: "resume", Section: "skills", Key: "items", Value: `[
			{"icon":"/static/html.svg","name":"html 5"},
			{"icon":"/static/css.svg","name":"css"},
			{"icon":"/static/js.svg","name":"javascript"},
			{"icon":"/static/react.svg","name":"react.js"},
			{"icon":"/static/next.svg","name":"next.js"},
			{"icon":"/static/tailwind.svg","name":"tailwind css"}
		]`},
		{PageName: "resume", Section: "about", Key: "title", Value: "About me"},
		{PageName: "resume", Section: "about", Key: "description", Value: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos"},
		{PageName: "resume", Section: "about", Key: "items", Value: `[
			{"fieldName":"Name","fieldValue":"+(62) 812-3456-7890"},
			{"fieldName":"Email","fieldValue":"Mahathirrizky@gmail.com"},
			{"fieldName":"Nationality","fieldValue":"Indonesia"},
			{"fieldName":"Experienxe","fieldValue":"12+ Years"},
			{"fieldName":"Freelance","fieldValue":"Available"},
			{"fieldName":"Languages","fieldValue":"English, Indonesian"}
		]`},

		// WorkPage
		{PageName: "work", Section: "projects", Key: "items", Value: `[
			{"num":"01","category":"Web Development","title":"Project 1","description":"Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos","stack":[{"name":"html 5"},{"name":"css"},{"name":"javascript"},{"name":"next.js"},{"name":"tailwind css"}],"image":"/static/thumb1.png","live":"http://github.com/mahathirrizky","github":""},
			{"num":"02","category":"UI/UX Design","title":"Project 2","description":"Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos","stack":[{"name":"html 5"},{"name":"css"},{"name":"javascript"},{"name":"react.js"},{"name":"next.js"},{"name":"tailwind css"}],"image":"/static/thumb2.png","live":"","github":""},
			{"num":"03","category":"Logo Design","title":"Project 3","description":"Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos","stack":[{"name":"html 5"},{"name":"css"},{"name":"javascript"},{"name":"react.js"},{"name":"next.js"},{"name":"tailwind css"}],"image":"/static/thumb3.png","live":"","github":""}
		]`},

		// ContactPage
		{PageName: "contact", Section: "form", Key: "title", Value: "Let’s work together"},
		{PageName: "contact", Section: "form", Key: "description", Value: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Eligendi vero quaerat dignissimos autem ab."},
		{PageName: "contact", Section: "form", Key: "button_text", Value: "Send message"},
		{PageName: "contact", Section: "info", Key: "items", Value: `[
			{"icon":"faPhone","title":"Phone","description":"+1 234 567 890"},
			{"icon":"faEnvelope","title":"Email","description":"example@example.com"},
			{"icon":"faMapMarkerAlt","title":"Address","description":"1234 Street Name, City Name, Country Name"}
		]`},
	}

	for _, content := range contents {
		// Check if a record with the same unique combination of PageName, Section, and Key already exists
		var existing models.Content
		if DB.Where("page_name = ? AND section = ? AND key = ?", content.PageName, content.Section, content.Key).First(&existing).Error == gorm.ErrRecordNotFound {
			DB.Create(&content)
		}
	}
}
