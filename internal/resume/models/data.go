package models

import (
	"time"
)

var ResumeMock = Resume{
	FirstName: "Mikhail",
	LastName:  "Krasilnikov",
	Title:     "Lead Software Engineer",
	Contacts: map[string]string{
		"Phone":      "+1(408)480-3600",
		"Email":      "mickey.krasilnikov@gmail.com",
		"LinkedIn":   "https://www.linkedin.com/in/mickeykrasilnikov/",
		"Credly":     "https://www.credly.com/users/mickey-krasilnikov/badges",
		"Accredible": "https://www.credential.net/profile/mikhailkrasilnikov328966/wallet",
		"Location":   "Sunnyvale, CA, 94085"},
	Summary: []string{
		"12+ years of professional programming experience with consistently increasing responsibilities in the design and development of commercial applications on different platforms and environments",
		"Certified Cloud Solution Architect (Azure, GCP, AWS)",
		"Experienced as a Team Leader",
		"Involved in full Software Development Life Cycle (SDLC)",
	},
	Skills: []SkillGroup{
		{
			Name: "Programming Languages",
			SubGroups: []SkillGroup{
				{
					Name: "Backend",
					Skills: []Skill{
						{
							Name:           "C#",
							AdditionalInfo: ".Net Core, .Net",
						},
						{
							Name:           "GoLang",
							AdditionalInfo: "Gin",
						},
						{
							Name:           "Javascript",
							AdditionalInfo: "Node.js",
						},
						{
							Name:           "Python",
							AdditionalInfo: "FastAPI",
						},
					},
				},
				{
					Name: "Frontend",
					Skills: []Skill{
						{
							Name:           "TypeScript",
							AdditionalInfo: "Angular, Vue.js",
						},
						{
							Name: "Javascript",
						},
						{
							Name: "HTML",
						},
						{
							Name: "CSS",
						},
					},
				},
				{
					Name: "Database",
					Skills: []Skill{
						{
							Name:           "T-SQL",
							AdditionalInfo: "MS SQL",
						},
						{
							Name:           "SQL/PSM",
							AdditionalInfo: "MySQL",
						},
						{
							Name:           "PL-SQL",
							AdditionalInfo: "PostgreSQL, Oracle",
						},
					},
				},
			},
		},
		{
			Name: "Databases",
			SubGroups: []SkillGroup{
				{
					Name: "SQL",
					Skills: []Skill{
						{
							Name: "MS SQL",
						},
						{
							Name: "MySQL",
						},
						{
							Name: "PostgreSQL",
						},
						{
							Name: "Oracle",
						},
					},
				},
				{
					Name: "NoSQL",
					Skills: []Skill{
						{
							Name: "MongoDB",
						},
						{
							Name: "Redis",
						},
						{
							Name: "CosmosDB",
						},
						{
							Name: "DynamoDB",
						},
						{
							Name: "BigTable",
						},
					},
				},
			},
		},
		{
			Name: "Container Services",
			Skills: []Skill{
				{
					Name: "Docker",
				},
				{
					Name: "Docker Swarm",
				},
				{
					Name: "Kubernetes",
				},
			},
		},
		{
			Name: "CI/CD",
			Skills: []Skill{
				{
					Name: "Azure DevOps",
				},
				{
					Name: "GitHub Actions",
				},
				{
					Name: "AWS CodePipeline",
				},
				{
					Name: "Octopus",
				},
				{
					Name: "Jenkins",
				},
				{
					Name: "TeamCity",
				},
			},
		},
		{
			Name: "Bug Tracking",
			Skills: []Skill{
				{
					Name: "JIRA",
				},
				{
					Name: "Azure DevOps",
				},
			},
		},
		{
			Name: "Source Control",
			Skills: []Skill{
				{
					Name: "GIT",
				},
				{
					Name: "SVN",
				},
			},
		},
		{
			Name: "OS",
			Skills: []Skill{
				{
					Name: "Windows",
				},
				{
					Name: "MacOS",
				},
				{
					Name: "Ubuntu",
				},
				{
					Name: "Debian",
				},
			},
		},
		{
			Name: "Clouds",
			Skills: []Skill{
				{
					Name: "Azure",
				},
				{
					Name: "GCP",
				},
				{
					Name: "AWS",
				},
			},
			SubGroups: []SkillGroup{
				{
					Name: "Compute",
					SubGroups: []SkillGroup{
						{
							Name: "IaaS",
							Skills: []Skill{
								{
									Name: "Azure VM",
								},
								{
									Name: "Google Compute Engine",
								},
								{
									Name: "AWS EC2",
								},
								{
									Name: "Azure Container Instance",
								},
								{
									Name: "AWS ECS",
								},
							},
						},
						{
							Name: "PaaS",
							Skills: []Skill{
								{
									Name: "Azure App Service",
								},
								{
									Name: "Google App Engine",
								},
								{
									Name: "AWS Beanstalk",
								},
								{
									Name: "Azure AKS",
								},
								{
									Name: "Google Kubernetes Service",
								},
								{
									Name: "AWS EKS",
								},
							},
						},
						{
							Name: "Serverless",
							Skills: []Skill{
								{
									Name: "Azure Functions",
								},
								{
									Name: "Azure Logic Apps",
								},
								{
									Name: "Google Cloud Functions",
								},
								{
									Name: "Google Workflows",
								},
								{
									Name: "AWS Lambda",
								},
								{
									Name: "AWS Step Functions",
								},
								{
									Name: "AWS Fargate",
								},
							},
						},
					},
				},
				{
					Name: "Storage",
					SubGroups: []SkillGroup{
						{
							Name: "RDS",
							Skills: []Skill{
								{
									Name: "Azure SQL",
								},
								{
									Name: "Google Cloud SQL",
								},
								{
									Name: "AWS RDS",
								},
							},
						},
						{
							Name: "NoSQL",
							Skills: []Skill{
								{
									Name: "Azure CosmosDB",
								},
								{
									Name: "Azure Table Storage",
								},
								{
									Name: "Google BigTable",
								},
								{
									Name: "AWS DynamoDB",
								},
							},
						},
						{
							Name: "Object",
							Skills: []Skill{
								{
									Name: "Azure Blob Storage",
								},
								{
									Name: "Google Cloud Storage",
								},
								{
									Name: "AWS S3",
								},
							},
						},
						{
							Name: "File",
							Skills: []Skill{
								{
									Name: "Azure File Storage",
								},
								{
									Name: "Google Filestore",
								},
								{
									Name: "AWS EFS",
								},
							},
						},
						{
							Name: "Archive",
							Skills: []Skill{
								{
									Name: "Azure Archive Storage",
								},
								{
									Name: "Google Archive Storage",
								},
								{
									Name: "Amazon S3 Glacier",
								},
							},
						},
					},
				},
				{
					Name: "Messaging",
					Skills: []Skill{
						{
							Name: "Azure Service Bus",
						},
						{
							Name: "Azure Queue Storage",
						},
						{
							Name: "Azure Notification Hub",
						},
						{
							Name: "Google Pub/Sub",
						},
						{
							Name: "AWS SQS",
						},
						{
							Name: "AWS SNS",
						},
					},
				},
				{
					Name: "Secret Management",
					Skills: []Skill{
						{
							Name: "Azure KeyVault",
						},
						{
							Name: "Google Secret Manager",
						},
						{
							Name: "Google Cloud KMS",
						},
						{
							Name: "AWS Secrets Manager",
						},
						{
							Name: "AWS KMS",
						},
					},
				},
				{
					Name: "Monitoring and Logging",
					Skills: []Skill{
						{
							Name: "Azure Monitor",
						},
						{
							Name: "Google Cloud Logging",
						},
						{
							Name: "Google Cloud Monitoring",
						},
						{
							Name: "AWS CloudWatch",
						},
					},
				},
				{
					Name: "Networking",
					SubGroups: []SkillGroup{
						{
							Name: "Virtual Network",
							Skills: []Skill{
								{
									Name: "Azure VNET",
								},
								{
									Name: "Google VPC",
								},
								{
									Name: "AWS VPC",
								},
							},
						},
						{
							Name: "Load Balancing",
							Skills: []Skill{
								{
									Name: "Azure Load Balancer",
								},
								{
									Name: "Azure Traffic Manager",
								},
								{
									Name: "Google Cloud Load Balancing",
								},
								{
									Name: "AWS ELB",
								},
							},
						},
						{
							Name: "Firewall",
							Skills: []Skill{
								{
									Name: "Azure WAF",
								},
								{
									Name: "Google Cloud Firewalls",
								},
								{
									Name: "AWS WAF",
								},
							},
						},
						{
							Name: "DNS",
							Skills: []Skill{
								{
									Name: "Azure DNS",
								},
								{
									Name: "Google Cloud DNS",
								},
								{
									Name: "AWS Route 53",
								},
							},
						},
						{
							Name: "CDN",
							Skills: []Skill{
								{
									Name: "Azure CDN",
								},
								{
									Name: "Google Cloud CDN",
								},
								{
									Name: "AWS CloudFront",
								},
							},
						},
						{
							Name: "Gateways",
							Skills: []Skill{
								{
									Name: "Azure API Gateway",
								},
								{
									Name: "Azure API Management",
								},
								{
									Name: "Azure Frontdoor",
								},
								{
									Name: "Google API Management",
								},
								{
									Name: "AWS API Gateway",
								},
							},
						},
					},
				},
				{
					Name: "DevOps",
					Skills: []Skill{
						{
							Name:           "Azure DevOps",
							AdditionalInfo: "Boards, Pipelines, Repos, TestPlans, Artifacts",
						},
						{
							Name:           "GCP DevOps",
							AdditionalInfo: "Cloud Build, Artifact Registry",
						},
						{
							Name:           "AWS DevOps",
							AdditionalInfo: "CodePipeline, CodeBuild, CodeDeploy, CodeStar",
						},
					},
				},
			},
		},
	},
	Experience: []Experience{
		{
			Title:     "Lead Software Engineer",
			Company:   "EPAM Systems",
			StartDate: time.Date(2019, time.May, 1, 0, 0, 0, 0, time.UTC),
			Projects: []Project{
				{
					Client:    "Broadridge",
					StartDate: time.Date(2021, time.November, 1, 0, 0, 0, 0, time.UTC),
					Roles: []string{
						"Key Developer",
						"Team Lead",
					},
					Envirnment: []string{
						".NET Core 6.0",
						"Vue.js",
						"Python 3.7",
						"AWS",
						"AWS Lambda",
						"AWS S3",
						"AWS Beanstalk",
						"AWS EC2",
						"AWS Code Commit",
						"AWS Code Build",
						"AWS Code Deploy",
						"AWS Code Pipeline",
						"Docker",
						"AWS RDS",
						"GIT",
					},
				},
				{
					Client:    "ElipsLife",
					StartDate: time.Date(2019, time.May, 1, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2021, time.November, 1, 0, 0, 0, 0, time.UTC),
					Roles: []string{
						"Key Developer",
						"Team Lead",
					},
					Envirnment: []string{
						".NET Core 6.0",
						"Angular 13",
						"Go Lang",
						"Azure",
						"Azure DevOps",
						"Octopus",
						"OKTA",
						"Docker",
						"Cosmos DB",
						"Redis",
						"MS SQL",
						"GIT",
					},
					TaskPerformed: []string{
						"Designed microservice-based hybrid-cloud solution architecture for multi-role self-service portal for managing insurance claims, underwriting, and broker reports.",
						"Designed and developed microservices following a RESTful approach and applying best practices. • Covered code with unit and contract tests and maintained code coverage on 80%.",
						"Configured CI/CD pipelines and established automated quality gates on different stages.",
						"Mentored new joiners and supported them during the ramp-up period.",
						"Performed systematic knowledge transfer to the client support team, including the creation of visual documentation. • Introduced and ensured the development team follows the very best engineering practices (with respect to various limitations on the client side).",
					},
				},
			},
		},
		{
			Title:     "Front Office Risk System Developer (AVP) (Senior Software Engineer)",
			Company:   "Credit Suisse",
			StartDate: time.Date(2015, time.September, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC),
			Projects: []Project{
				{
					Roles: []string{
						"Key Developer",
						"Tech Lead",
					},
					Envirnment: []string{
						".NET Core 2.1",
						".NET 4.6",
						"WPF",
						"WCF",
						"EF",
						"MS SQL",
						"MS SSAS",
						"TeamCity",
						"Jira",
						"Confluence",
						"Splunk",
						"GIT",
					},
					TaskPerformed: []string{
						"Maintained Risk Management System for Credit Suisse Front Office and developed new modules. • Improved performance and decreased memory usage in CS in-house distributed cache.",
						"Splatted monolith server-side architecture to microservices.",
						"Analyzed and arranged requirements with traders, BAU team and QA.",
						"Supported worldwide users as a part of the SL3 team. • Administrated CI/CD on TeamCity.",
						"Documented design decisions and usage examples.",
						"Mentored new joiners.",
					},
				},
			},
		},
		{
			Title:     "Senior Software Engineer",
			Company:   "SINTEGRO SOFT",
			StartDate: time.Date(2014, time.September, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2015, time.September, 1, 0, 0, 0, 0, time.UTC),
			Projects: []Project{
				{
					Client: "Yandex",
					Roles: []string{
						"Key Developer",
					},
					Envirnment: []string{
						".NET 4.5",
						"ASP.NET MVC 5",
						"ASP.NET Web API",
						"WPF",
						"MS SQL",
						"EF",
						"TFS",
						"WIX",
						"GIT",
					},
					TaskPerformed: []string{
						"Participated in the development of the system that helps organize corporate bonuses, such as meals, mobile phones, fitness, transportation, etc.",
						"Developed user personal page that contains account balance and transaction history. ",
						"Developed desktop administration systems for corporate bonus managers.",
						"Documented design decisions, component APIs and usage examples.",
						"Participated in the software integration process.",
					},
				},
			},
		},
		{
			Title:     "Software Engineer",
			Company:   "Gollard",
			StartDate: time.Date(2013, time.July, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2014, time.August, 1, 0, 0, 0, 0, time.UTC),
			Projects: []Project{
				{
					Roles: []string{
						"Key Developer",
					},
					Envirnment: []string{
						".NET 4.5",
						"WPF",
						"WCF",
						"ASP.NET Web API",
						"Esri ArcGIS",
						"MS SQL",
						"MySQL",
						"GitLab",
						"WIX",
						"GIT",
					},
					TaskPerformed: []string{
						"Participated in the development of desktop GIS application for the power distribution company dealing with electrical grids, stations and facilities on the map.",
						"Implemented a feature that allows monitoring the service transport on the map in real-time",
						"Introduced the possibility to view video streams from the cameras marked on the map.",
						"Participated in the development of the solution to maintain security and safety of the buildings and areas using cameras and different types of sensors.",
						"Implemented an alerting system that works based on the biometric data from the cameras or events from the sensors.",
					},
				},
			},
		},
		{
			Title:     "Software Engineer",
			Company:   "SGC",
			StartDate: time.Date(2012, time.August, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2013, time.July, 1, 0, 0, 0, 0, time.UTC),
			Projects: []Project{
				{
					Roles: []string{
						"Key Developer",
					},
					Envirnment: []string{
						".NET 4.0", "WPF", "SOAP", "MS SQL", "WIX", "SVN",
					},
					TaskPerformed: []string{
						"Developed software for composing estimates to determine the cost of construction.",
						"Integrated with other well-known systems available on the market for composing estimates.",
					},
				},
			},
		},
		{
			Title:     "Software Engineer",
			Company:   "Research and Production Center 'Start'",
			StartDate: time.Date(2010, time.September, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2012, time.July, 1, 0, 0, 0, 0, time.UTC),
			Projects: []Project{
				{
					Roles: []string{
						"Key Developer",
					},
					Envirnment: []string{
						".NET 3.5",
						"WinForms",
						"SOAP",
						"MS SQL",
						"SVN",
					},
					TaskPerformed: []string{
						"Developed software for designing micro boards and creating a mask for interacting with third-party equipment for printing this kind of product.",
					},
				},
			},
		},
	},
	Certifications: []Certification{
		{
			Name:            "Azure Solutions Architect Expert",
			Issuer:          "Microsoft",
			IssueDate:       time.Date(2020, time.June, 10, 0, 0, 0, 0, time.UTC),
			ExpirationDate:  time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC),
			VerificationURL: "https://www.credly.com/badges/42f9be70-0b40-4bb0-8256-bd9d573e36af/public_url",
		},
		{
			Name:            "Professional Cloud Architect",
			Issuer:          "Google",
			IssueDate:       time.Date(2020, time.October, 27, 0, 0, 0, 0, time.UTC),
			ExpirationDate:  time.Date(2022, time.October, 27, 0, 0, 0, 0, time.UTC),
			VerificationURL: "https://www.credential.net/786961ad-0225-4bb0-883a-3c2feceb5174",
		},
		{
			Name:            "AWS Certified Solutions Architect Associate",
			Issuer:          "Amazon",
			IssueDate:       time.Date(2021, time.April, 28, 0, 0, 0, 0, time.UTC),
			ExpirationDate:  time.Date(2024, time.April, 28, 0, 0, 0, 0, time.UTC),
			VerificationURL: "https://www.credly.com/badges/dc95ff2a-12d8-4816-91b2-a9292ae5df85/public_url",
		},
		{
			Name:            "DevOps Engineer Expert",
			Issuer:          "Microsoft",
			IssueDate:       time.Date(2020, time.December, 10, 0, 0, 0, 0, time.UTC),
			ExpirationDate:  time.Date(2022, time.December, 10, 0, 0, 0, 0, time.UTC),
			VerificationURL: "https://www.credly.com/badges/2d29613e-7c7e-481d-aba3-842b409fecd7/public_url",
		},
	},
	Education: []Education{
		{
			Name:         "Penza State University",
			Degree:       "Bachelor's degree",
			FieldOfStudy: "Design and technology of radio electronic devices",
			StartDate:    time.Date(2006, time.September, 1, 0, 0, 0, 0, time.UTC),
			EndDate:      time.Date(2013, time.June, 19, 0, 0, 0, 0, time.UTC),
			URL:          "https://international.pnzgu.ru/",
		},
	},
}
