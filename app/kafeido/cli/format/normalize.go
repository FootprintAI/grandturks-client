package format

import (
	"fmt"
	"strings"

	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

const protoDataSourcePrefix = "DATASOURCE_"

func NormalizedDataSource(s string) appmodels.DatastreamDataSource {
	if strings.HasPrefix(s, protoDataSourcePrefix) {
		return appmodels.DatastreamDataSource(s)
	}
	return appmodels.DatastreamDataSource(
		fmt.Sprintf("%s%s", protoDataSourcePrefix, strings.ToUpper(s)),
	)
}

func DeNormalizedDataSource(s appmodels.DatastreamDataSource) string {
	return strings.ToLower(string(s)[len(protoDataSourcePrefix):])
}

const protoStreamingInfoProtocoPrefix = "STREAMING_PROTOCOL_"

func NormalizedStreamingInfoProtocol(s string) appmodels.StreamingInfoProtocol {
	if strings.HasPrefix(s, protoStreamingInfoProtocoPrefix) {
		return appmodels.StreamingInfoProtocol(s)
	}
	return appmodels.StreamingInfoProtocol(
		fmt.Sprintf("%s%s", protoStreamingInfoProtocoPrefix, strings.ToUpper(s)),
	)
}

func DeNormalizedStreamingInfoProtocol(s appmodels.StreamingInfoProtocol) string {
	return strings.ToLower(string(s)[len(protoStreamingInfoProtocoPrefix):])
}
