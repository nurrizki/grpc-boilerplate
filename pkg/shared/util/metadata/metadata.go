package metadata

import (
	"context"
	"encoding/base64"
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	binHdrSuffix = "-bin"
)

// MetaData :
type MetaData struct {
	ClientName  string
	ClientIP    string
	FrontEndIp  string
	Uuid        string
	ActivityId  string
	XTraceId    string
	UserID      string
	BranchCode  string
	UserLevelId string
}

// metadataTextMap extends a metadata.MD to be an opentracing textmap
type MetadataTextMap metadata.MD

// Set is a opentracing.TextMapReader interface that extracts values.
func (m MetadataTextMap) Set(key, val string) {
	// gRPC allows for complex binary values to be written.
	encodedKey, encodedVal := encodeKeyValue(key, val)
	// The metadata object is a multimap, and previous values may exist, but for opentracing headers, we do not append
	// we just override.
	m[encodedKey] = []string{encodedVal}
}

// ForeachKey is a opentracing.TextMapReader interface that extracts values.
func (m MetadataTextMap) ForeachKey(callback func(key, val string) error) error {
	for k, vv := range m {
		for _, v := range vv {
			if err := callback(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// encodeKeyValue encodes key and value qualified for transmission via gRPC.
// note: copy pasted from private values of grpc.metadata
func encodeKeyValue(k, v string) (string, string) {
	k = strings.ToLower(k)
	if strings.HasSuffix(k, binHdrSuffix) {
		val := base64.StdEncoding.EncodeToString([]byte(v))
		v = string(val)
	}
	return k, v
}

func MappingData(ctx context.Context) MetaData {
	var metaData MetaData
	md := ExtractIncoming(ctx)
	metaData.ClientName = md.Get("clientname")
	metaData.FrontEndIp = md.Get("frontendip")
	metaData.ClientIP = md.Get("clientip")
	metaData.Uuid = md.Get("uuid")
	metaData.ActivityId = md.Get("activityid")
	metaData.XTraceId = md.Get("x-trace-id")
	metaData.BranchCode = md.Get("branchcode")
	metaData.UserID = md.Get("userid")
	return metaData
}
