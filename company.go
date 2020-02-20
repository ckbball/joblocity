package quik

import (
  "errors"

  "go.mongodb.org/mongo-driver/bson/primitive"
  "golang.org/x/crypto/bcrypt"
)

type User struct {
  Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
  Email        string             `json:"email,omitempty" bson:"email,omitempty"`
  Password     string             `json:"password,omitempty" bson:"password,omitempty"`
  FirstName    string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
  LastName     string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
  JobSearch    int                `json:"job_search,omitempty" bson:"job_search,omitempty"` // 0 - not looking, 1 - open to opportunities, 2 - active
  Profile      Profile            `json:"profile,omitempty" bson:"profile,omitempty"`
  Applications []string           `json:"applications,omitempty" bson:"applications,omitempty"` // holds the id's of applications
  SavedJobs    []string           `json:"saved_jobs,omitempty" bson:"saved_jobs,omitempty"`     // holds the id's of jobs that they want to come back to later
}

type Profile struct {
  Skills      []Skill
  WorkHistory []Work
  Projects    []Project
}

type Skill struct {
  SkillName       string
  ExperienceLevel string // determined by how many projects you have used it on
}

type Work struct {
  CompanyName  string
  BulletPoints []string
  Length       int // in months
}

type Project struct {
  ProjectName string
  Details     string
  TechUsed    []string
  LiveLink    string
  GitLink     string
}

func (u *User) HashPassword(pass string) (string, error) {
  if len(pass) == 0 {
    return "", errors.New("password empty")
  }
  hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
  return string(hash), err
}

func (u *User) CheckPassword(pass string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
  return err == nil
}
