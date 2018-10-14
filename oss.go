package barkup

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

// OSS is a `Storer` interface that puts an ExportResult to the specified Aliyun OSS bucket. Don't use your main Aliyun keys for this!! Create sub ac
type OSS struct {
	// oss-[xx]-[xxxxx].aliyuncs.com
	// oss-cn-shenzhen.aliyun.com
	Endpoint string

	// name of bucket
	Bucket string

	// aliyun oss access key
	AccessKey string

	// aliyun oss secret key
	SecretKey string
}

// Store puts an `ExportResult` struct to an aliyun oss bucket within the specified directory
func (x *OSS) Store(result *ExportResult, directory string) *Error {
	if result.Error != nil {
		return result.Error
	}
	// file, err := os.Open(result.Path)
	// if err != nil {
	// 	return makeErr(err, "")
	// }
	// defer file.Close()

	client, err := oss.New(x.Endpoint, x.AccessKey, x.SecretKey)
	if err != nil {
		makeErr(err, "")
	}

	bucket, err := client.Bucket(x.Bucket)
	if err != nil {
		makeErr(err, "")
	}

	err = bucket.PutObjectFromFile(directory+result.Filename(), result.Path)
	return makeErr(err, "")
}
