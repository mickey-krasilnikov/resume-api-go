package memory

import (
	"time"

	"github.com/mickey-krasilnikov/resume-api-go/internal/resume/models"
)

var ResumeMock = models.Resume{
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
	Skills: []models.SkillGroup{
		{
			Name: "Programming Languages",
			SubGroups: []models.SkillGroup{
				{
					Name: "Backend",
					Skills: []models.Skill{
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
					Skills: []models.Skill{
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
					Skills: []models.Skill{
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
			SubGroups: []models.SkillGroup{
				{
					Name: "SQL",
					Skills: []models.Skill{
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
					Skills: []models.Skill{
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
			Skills: []models.Skill{
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
			Skills: []models.Skill{
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
			Skills: []models.Skill{
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
			Skills: []models.Skill{
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
			Skills: []models.Skill{
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
			Skills: []models.Skill{
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
			SubGroups: []models.SkillGroup{
				{
					Name: "Compute",
					SubGroups: []models.SkillGroup{
						{
							Name: "IaaS",
							Skills: []models.Skill{
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
							Skills: []models.Skill{
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
							Skills: []models.Skill{
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
					SubGroups: []models.SkillGroup{
						{
							Name: "RDS",
							Skills: []models.Skill{
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
							Skills: []models.Skill{
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
							Skills: []models.Skill{
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
							Skills: []models.Skill{
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
							Skills: []models.Skill{
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
					Skills: []models.Skill{
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
					Skills: []models.Skill{
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
					Skills: []models.Skill{
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
					SubGroups: []models.SkillGroup{
						{
							Name: "Virtual Network",
							Skills: []models.Skill{
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
							Skills: []models.Skill{
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
							Skills: []models.Skill{
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
							Skills: []models.Skill{
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
							Skills: []models.Skill{
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
							Skills: []models.Skill{
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
					Skills: []models.Skill{
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
	Experience: []models.Experience{},
	Certifications: []models.Certification{
		{
			ID:              0,
			Name:            "Azure Solutions Architect Expert",
			Issuer:          "Microsoft",
			IssueDate:       time.Date(2020, time.June, 10, 0, 0, 0, 0, time.UTC),
			ExpirationDate:  time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC),
			VerificationURL: "https://www.credly.com/badges/42f9be70-0b40-4bb0-8256-bd9d573e36af/public_url",
		},
		{
			ID:              1,
			Name:            "Professional Cloud Architect",
			Issuer:          "Google",
			IssueDate:       time.Date(2020, time.October, 27, 0, 0, 0, 0, time.UTC),
			ExpirationDate:  time.Date(2022, time.October, 27, 0, 0, 0, 0, time.UTC),
			VerificationURL: "https://www.credential.net/786961ad-0225-4bb0-883a-3c2feceb5174",
		},
		{
			ID:              2,
			Name:            "AWS Certified Solutions Architect Associate",
			Issuer:          "Amazon",
			IssueDate:       time.Date(2021, time.April, 28, 0, 0, 0, 0, time.UTC),
			ExpirationDate:  time.Date(2024, time.April, 28, 0, 0, 0, 0, time.UTC),
			VerificationURL: "https://www.credly.com/badges/dc95ff2a-12d8-4816-91b2-a9292ae5df85/public_url",
		},
		{
			ID:              3,
			Name:            "DevOps Engineer Expert",
			Issuer:          "Microsoft",
			IssueDate:       time.Date(2020, time.December, 10, 0, 0, 0, 0, time.UTC),
			ExpirationDate:  time.Date(2022, time.December, 10, 0, 0, 0, 0, time.UTC),
			VerificationURL: "https://www.credly.com/badges/2d29613e-7c7e-481d-aba3-842b409fecd7/public_url",
		},
	},
	Education: []models.Education{
		{
			ID:           0,
			Name:         "Penza State University",
			Degree:       "Bachelor's degree",
			FieldOfStudy: "Design and technology of radio electronic devices",
			StartDate:    time.Date(2006, time.September, 1, 0, 0, 0, 0, time.UTC),
			EndDate:      time.Date(2013, time.June, 19, 0, 0, 0, 0, time.UTC),
			URL:          "https://international.pnzgu.ru/",
		},
	},
}
