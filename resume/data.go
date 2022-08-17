package resume

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
							Name:           "Python",
							AdditionalInfo: "FastAPI",
						},
						{
							Name:           "Javascript",
							AdditionalInfo: "Node.js",
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
	Experience:     []Experience{},
	Certifications: []Certification{},
	Education:      []Education{},
}
