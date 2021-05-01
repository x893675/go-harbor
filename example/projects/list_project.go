package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	goharbor "github.com/x893675/go-harbor"
	"github.com/x893675/go-harbor/schema"
)

const HarborAddress = "https://myharbor.com"

func main() {
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	harborClient, err := goharbor.NewClientWithOpts(goharbor.WithHost(HarborAddress),
		goharbor.WithHTTPClient(c),
		goharbor.WithBasicAuth("admin", "Harbor12345"))
	if err != nil {
		panic(err)
	}
	pr, err := harborClient.ListProjects(context.TODO(), schema.ProjectListOptions{})
	if err != nil {
		panic(err)
	}
	for _, item := range pr {
		log.Printf("%+v", item)
	}

	err = harborClient.CreateProject(context.TODO(), schema.CreateProjectOptions{
		Name:         "test",
		Metadata:     nil,
		CVEAllowlist: nil,
		StorageLimit: nil,
		CountLimit:   nil,
	})
	if err != nil {
		panic(err)
	}

	project, err := harborClient.GetProject(context.TODO(), "test")
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", project)

	exist, err := harborClient.ProjectExist(context.TODO(), "test")
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", exist)

	exist, err = harborClient.ProjectExist(context.TODO(), "aaa")
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", exist)

	prMembers, err := harborClient.ListProjectMembers(context.TODO(), &schema.ProjectMemberListOptions{
		ProjectID: project.ProjectID})
	if err != nil {
		panic(err)
	}
	log.Printf("%v", prMembers)

	exist, err = harborClient.CheckProjectMemberExist(context.TODO(), project.ProjectID, "test")
	if err != nil {
		panic(err)
	}
	log.Println(exist)

	err = harborClient.AddProjectMember(context.TODO(), project.ProjectID, schema.ProjectMember{
		RoleID:      schema.Developer,
		MemberGroup: schema.UserGroup{},
		MemberUser: schema.UserEntity{
			Username: "test",
		},
	})

	prMembers, err = harborClient.ListProjectMembers(context.TODO(), &schema.ProjectMemberListOptions{
		ProjectID:  project.ProjectID,
		EntityName: "test",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("%v", prMembers)

	exist, err = harborClient.CheckProjectMemberExist(context.TODO(), project.ProjectID, "test")
	if err != nil {
		panic(err)
	}
	log.Println(exist)

	err = harborClient.RemoveProjectMember(context.TODO(), project.ProjectID, prMembers[0].ID)
	if err != nil {
		panic(err)
	}

	prMembers, err = harborClient.ListProjectMembers(context.TODO(), &schema.ProjectMemberListOptions{
		ProjectID: project.ProjectID,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("%v", prMembers)

	err = harborClient.DeleteProject(context.TODO(), project.Name)
	if err != nil {
		panic(err)
	}

	exist, err = harborClient.CheckUserExist(context.TODO(), "test")
	if err != nil {
		panic(err)
	}
	log.Println(exist)

	exist, err = harborClient.CheckUserExist(context.TODO(), "aaa")
	if err != nil {
		panic(err)
	}
	log.Println(exist)
}
