package grpc

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToGRPCTime конвертирвет
func ToGRPCTime(t time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{Nanos: int32(t.Nanosecond()), Seconds: t.Unix()}
}

// FromGRPCTime конвертирует
func FromGRPCTime(t *timestamppb.Timestamp) time.Time {
	return time.Unix(t.GetSeconds(), int64(t.GetNanos()))
}
