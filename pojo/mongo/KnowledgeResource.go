package mongo

// KnowledgeResource represents a knowledge resource
type KnowledgeResource struct {
	Type  string `bson:"type,omitempty" json:"type"`
	Value string `bson:"value,omitempty" json:"value"`
}
