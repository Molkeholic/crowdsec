package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Alert holds the schema definition for the Alert entity.
type Alert struct {
	ent.Schema
}

// Fields of the Alert.
func (Alert) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
		field.String("scenario"),
		field.String("bucketId").Default("").Optional(),
		field.String("message").Default("").Optional(),
		field.Int32("eventsCount").Default(0).Optional(),
		field.Time("startedAt").Default(time.Now).Optional(),
		field.Time("stoppedAt").Default(time.Now).Optional(),
		field.String("sourceIp").
			Optional(),
		field.String("sourceRange").
			Optional(),
		field.String("sourceAsNumber").
			Optional(),
		field.String("sourceAsName").
			Optional(),
		field.String("sourceCountry").
			Optional(),
		field.Float32("sourceLatitude").
			Optional(),
		field.Float32("sourceLongitude").
			Optional(),
		field.String("sourceScope").Optional(),
		field.String("sourceValue").Optional(),
		field.Int32("capacity").Optional(),
		field.String("leakSpeed").Optional(),
		field.String("scenarioVersion").Optional(),
		field.String("scenarioHash").Optional(),
		field.Bool("simulated").Default(false),
	}
}

// Edges of the Alert.
func (Alert) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Machine.Type).
			Ref("alerts").
			Unique(),
		edge.To("decisions", Decision.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("events", Event.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("metas", Meta.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (Alert) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
