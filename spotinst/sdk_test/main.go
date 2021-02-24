package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
	"log"
)

func main() {
	sess := session.New()

	sess.Config.Credentials.Get()
	svc := ocean.New(sess)

	ctx := context.Background()

	out, err := svc.CloudProviderAWS().ListLaunchSpecs(ctx, &aws.ListLaunchSpecsInput{
		OceanID: spotinst.String("o-be42e329"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to list groups: %v", err)
	}

	// Output all groups, if any.
	if len(out.LaunchSpecs) > 0 {
		for _, launchSpec := range out.LaunchSpecs {
			log.Printf("launchSpec %q: %s",
				spotinst.StringValue(launchSpec.ID),
				stringutil.Stringify(launchSpec))
		}
	}
}
