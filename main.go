package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const baseURL = "https://api.modrinth.com/v2"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// Project represents the structure of a Modrinth project
type Project struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Version represents the structure of a Modrinth project version
type Version struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

func SearchProjects() {
	searchURL := fmt.Sprintf("%s/search", baseURL)
	query := url.Values{}
	query.Set("q", "fabric")
	query.Set("limit", "5")
}

// FetchProject fetches the project details from Modrinth API
func FetchProject(projectID string) (*Project, error) {
	url := fmt.Sprintf("%s/project/%s", baseURL, projectID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch project: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var project Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

// FetchProjectsByAuthor fetches all projects by a specific author
func FetchProjectsByAuthor(authorID string) ([]Project, error) {
	url := fmt.Sprintf("%s/user/%s/projects", baseURL, authorID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch projects: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var projects []Project
	err = json.Unmarshal(body, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

// FetchProjectVersions fetches all versions of a specific project
func FetchProjectVersions(projectID string) ([]Version, error) {
	url := fmt.Sprintf("%s/project/%s/version", baseURL, projectID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch project versions: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var versions []Version
	err = json.Unmarshal(body, &versions)
	if err != nil {
		return nil, err
	}

	return versions, nil
}

// FetchProjectDependencies fetches all dependencies of a specific project
func FetchProjectDependencies(projectID string) ([]Project, error) {
	url := fmt.Sprintf("%s/project/%s/dependencies", baseURL, projectID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch project dependencies: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dependencies []Project
	err = json.Unmarshal(body, &dependencies)
	if err != nil {
		return nil, err
	}

	return dependencies, nil
}

func main() {
	projectID := "your_project_id_here"
	project, err := FetchProject(projectID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Project ID: %s\n", project.ID)
	fmt.Printf("Project Slug: %s\n", project.Slug)
	fmt.Printf("Project Title: %s\n", project.Title)
	fmt.Printf("Project Description: %s\n", project.Description)

	// Example usage of FetchProjectsByAuthor
	authorID := "your_author_id_here"
	projects, err := FetchProjectsByAuthor(authorID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Projects by Author %s:\n", authorID)
	for _, p := range projects {
		fmt.Printf("- %s: %s\n", p.ID, p.Title)
	}

	// Example usage of FetchProjectVersions
	versions, err := FetchProjectVersions(projectID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Versions of Project %s:\n", projectID)
	for _, v := range versions {
		fmt.Printf("- %s: %s\n", v.ID, v.Version)
	}

	// Example usage of FetchProjectDependencies
	dependencies, err := FetchProjectDependencies(projectID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Dependencies of Project %s:\n", projectID)
	for _, d := range dependencies {
		fmt.Printf("- %s: %s\n", d.ID, d.Title)
	}
}
