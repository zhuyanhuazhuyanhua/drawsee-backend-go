package dto

import (
	"drawsee/pojo/mongo"
)

// AddKnowledgeDTO represents the data transfer object for adding knowledge
type AddKnowledgeDTO struct {
	Name      string                    `bson:"name,omitempty" json:"name"`
	Aliases   []string                  `bson:"aliases,omitempty" json:"aliases"`
	Resources []mongo.KnowledgeResource `bson:"resources,omitempty" json:"resources"`
	ParentId  string                    `bson:"parentId,omitempty" json:"parentId"`
}
