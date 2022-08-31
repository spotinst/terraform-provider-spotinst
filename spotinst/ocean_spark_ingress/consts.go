package ocean_spark_ingress

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

type AnnotationField string

const (
	AnnotationKey   AnnotationField = "key"
	AnnotationValue AnnotationField = "value"
)

const (
	Ingress            commons.FieldName = "ingress"
	ServiceAnnotations commons.FieldName = "service_annotations"
)
