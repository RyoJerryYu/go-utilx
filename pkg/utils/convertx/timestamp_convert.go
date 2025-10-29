package convertx

import (
	"database/sql"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimePtr2Pb converts a time.Time pointer to a protobuf timestamp.
func TimePtr2Pb(in *time.Time) *timestamppb.Timestamp {
	if in == nil {
		return nil
	}
	return timestamppb.New(*in)
}

// TimePb2Ptr converts a protobuf timestamp to a time.Time pointer.
func TimePb2Ptr(in *timestamppb.Timestamp) *time.Time {
	if in == nil {
		return nil
	}
	return Ptr(in.AsTime())
}

// TimeSql2Pb converts a sql.NullTime to a protobuf timestamp.
func TimeSql2Pb(in sql.NullTime) *timestamppb.Timestamp {
	if !in.Valid {
		return nil
	}
	return timestamppb.New(in.Time)
}

// TimePb2Sql converts a protobuf timestamp to a sql.NullTime.
func TimePb2Sql(in *timestamppb.Timestamp) sql.NullTime {
	if in == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  in.AsTime(),
		Valid: true,
	}
}

// TimeSql2Ptr converts a sql.NullTime to a time.Time pointer.
func TimeSql2Ptr(in sql.NullTime) *time.Time {
	if !in.Valid {
		return nil
	}
	return Ptr(in.Time)
}

// TimePtr2Sql converts a time.Time pointer to a sql.NullTime.
func TimePtr2Sql(in *time.Time) sql.NullTime {
	if in == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  *in,
		Valid: true,
	}
}

// TimePb2Go converts a protobuf timestamp to a time.Time.
// It will ensure the null timestamp is converted to the zero value of time.Time.
func TimePb2Go(in *timestamppb.Timestamp) time.Time {
	if in == nil {
		return time.Time{}
	}
	return in.AsTime()
}

// TimeGo2Pb converts a time.Time to a protobuf timestamp.
// It will ensure the zero value of time.Time is converted to the null timestamp.
func TimeGo2Pb(in time.Time) *timestamppb.Timestamp {
	if in.IsZero() {
		return nil
	}
	return timestamppb.New(in)
}
