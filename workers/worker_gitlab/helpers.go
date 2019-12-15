package worker_gitlab

import (
	"strings"

	"github.com/ossn/fixme_backend/models"
)

func split(r rune) bool {
	return r == ' ' || r == ':' || r == '.' || r == ','
}

// Searches if a label matches some known labels and updates the model
func searchForMatchingLabels(label *string, model *models.Issue) bool {
	switch strings.ToLower(*label) {
	case "help_wanted", "help wanted", "good first issue", "easyfix", "easy":
		model.ExperienceNeeded = "easy"
		return true
	case "moderate":
		model.ExperienceNeeded = "moderate"
		return true
	case "senior":
		model.ExperienceNeeded = "senior"
		return true
	case "enhancement":
		model.Type = "enhancement"
		return true
	case "bug", "bugfix":
		model.Type = "bugfix"
		return true
	}
	return false
}

var technologiesMap map[string]string

func create_technologies_map() {

	technologiesMap = map[string]string{
		"react.js": "React",
		"reactjs": "React",
		"react": "React",
		"node.js": "Node",
		"nodejs": "Node",
		"node": "Node",
		"vue.js": "Vue",
		"vuejs": "Vue",
		"vue": "Vue",
		"express.js": "Express",
		"expressjs": "Express",
		"express": "Express",
		"spring.js": "Spring",
		"springjs": "Spring",
		"spring": "Spring",
		"angular.js": "Angular",
		"angularjs": "Angular",
		"angular": "Angular",
		"redux.js": "Redux",
		"reduxjs": "Redux",
		"redux": "Redux",
		"asp.net": "ASP.NET",
		"django": "Django",
		"flask": "Flask",
		"laravel": "Laravel",
		"ruby on rails": "Rails",
		"rails": "Rails",
		"jquery": "jQuery",
		"drupal": "Drupal",
		".net": ".NET",
		".net core": ".NET Core",
		"pandas": "pandas",
		"unity 3d": "Unity",
		"unity": "Unity",
		"react native": "React Native",
		"tensorflow": "TensorFlow",
		"ansible": "Ansible",
		"cordova": "Apache Cordova",
		"apache cordova": "Apache Cordova",
		"xamarin": "Xamarin",
		"apache spark": "Apache Spark",
		"hadoop": "Apache Hadoop",
		"apache hadoop": "Apache Hadoop",
		"unreal engine": "Unreal Engine",
		"flutter": "Flutter",
		"pytorch": "PyTorch",
		"torch": "PyTorch",
		"puppet": "Puppet",
		"chef": "Chef",
		"cryengine": "CryEngine",
	}
}


func stringToWords(str string) []string{
	return strings.Fields(str)
}

// Remove duplicate strings from an array
func cleanupArray(s []string) (r []string) {
	seen := make(map[string]bool, len(s))
	seen[""] = true
	for _, str := range s {
		if _, exists := seen[str]; !exists {
			seen[str] = true
			r = append(r, str)
		}
	}
	return
}


func searchForMatchingTechnologies(words []string) []string {

	myTechnologies := []string{}

	for i := 0; i < len(words); i++ {
    if word := strings.ToLower(words[i]); technologiesMap[word] != "" {
			myTechnologies = append(myTechnologies, technologiesMap[word])
		}
	}

	for i := 0; i < len(words)-1; i++ {
		if word := strings.ToLower(words[i]) + " " + strings.ToLower(words[i+1]); technologiesMap[word] != "" {
			myTechnologies = append(myTechnologies, technologiesMap[word])
		}
	}

	for i := 0; i < len(words)-2; i++ {
		if word := strings.ToLower(words[i]) + " " + strings.ToLower(words[i+1]) + " " + strings.ToLower(words[i+2]); technologiesMap[word] != "" {
			myTechnologies = append(myTechnologies, technologiesMap[word])
		}
	}

	return myTechnologies
}