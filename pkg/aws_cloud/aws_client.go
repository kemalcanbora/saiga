package aws_cloud

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	configs "saiga/config"
)

func AwsClient() *session.Session {
	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	SharedConfigState: session.SharedConfigEnable,
	//}))
	//return sess

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(configs.Configure().AWSRegion),
		Credentials: credentials.NewStaticCredentials(configs.Configure().AWSAccessKeyId,
													  configs.Configure().AWSSecretAccessKey, ""),
	})
	return sess
}
