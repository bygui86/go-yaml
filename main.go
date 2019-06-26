package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

type Sample struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Metadata    Metadata `yaml:"metadata"`
}

type Metadata struct {
	Version int    `yaml:"version"`
	Group   string `yaml:"group"`
	Author  Author `yaml:"author"`
}

type Author struct {
	Name    string `yaml:"name"`
	Surname string `yaml:"surname"`
}

func (s *Sample) String() string {
	return "[ Title: " + s.Title + ", Description: " + s.Description + ", Metadata: " + s.Metadata.String() + " ]"
}

func (m *Metadata) String() string {
	return "[ Version: " + strconv.Itoa(m.Version) + ", Group: " + m.Group + ", Author: " + m.Author.String() + " ]"
}

func (a *Author) String() string {
	return "[ Name: " + a.Name + ", Surname: " + a.Surname + " ]"
}

func main() {
	sample := unmarshalFileToStruct("sample-origin.yaml")
	log.Println("Sample from YAML:", sample)

	tempName := sample.Metadata.Author.Name
	tempSurname := sample.Metadata.Author.Surname

	log.Println("Switching author name and surname...")
	sample.Metadata.Author.Name = tempSurname
	sample.Metadata.Author.Surname = tempName

	marshalFile("sample-new.yaml", sample)
	log.Println("Please check sample-new.yaml to verify switched author name and surname")
}

func marshalFile(filename string, sample *Sample) {
	file, _ := filepath.Rel(".", filename)
	yamlFile, yamlErr := yaml.Marshal(sample)
	if yamlErr != nil {
		log.Panicf("Error marshalling %s to file %s: %s", sample, filename, yamlErr.Error())
	}
	fileErr := ioutil.WriteFile(file, yamlFile, 0644)
	if fileErr != nil {
		log.Panicf("Error writing %s to file %s: %s", yamlFile, filename, fileErr.Error())
	}
}

func unmarshalFileToStruct(filename string) *Sample {
	file, _ := filepath.Rel(".", filename)
	yamlFile, yamlErr := ioutil.ReadFile(file)
	if yamlErr != nil {
		log.Panicf("Error reading file %s: %s", filename, yamlErr.Error())
	}
	var sample Sample
	unmarshalErr := yaml.Unmarshal(yamlFile, &sample)
	if unmarshalErr != nil {
		log.Panicf("Error unmarshalling file %s to struct: %s", filename, unmarshalErr.Error())
	}
	return &sample
}
