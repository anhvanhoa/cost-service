syntax = "proto3";

package cost_tracking.v1;

option go_package = "github.com/anhvanhoa/sf-proto/gen/cost_tracking/v1:proto_cost_tracking";

import "google/protobuf/timestamp.proto";
import "buf/validate/validate.proto";
import "common/v1/common.proto";

service CostTrackingService {
  rpc CreateCostTracking(CreateCostTrackingRequest)
      returns (CreateCostTrackingResponse);
  rpc GetCostTracking(GetCostTrackingRequest) returns (GetCostTrackingResponse);
  rpc UpdateCostTracking(UpdateCostTrackingRequest)
      returns (UpdateCostTrackingResponse);
  rpc DeleteCostTracking(DeleteCostTrackingRequest)
      returns (DeleteCostTrackingResponse);
  rpc ListCostTrackings(ListCostTrackingsRequest)
      returns (ListCostTrackingsResponse);
  rpc GetCostTrackingsByPlantingCycle(GetCostTrackingsByPlantingCycleRequest)
      returns (GetCostTrackingsByPlantingCycleResponse);
}

message CostTracking {
  string id = 1;
  string planting_cycle_id = 2;
  string cost_category = 3;
  string cost_type = 4;
  string item_name = 5;
  string description = 6;
  double quantity = 7;
  string unit = 8;
  double unit_cost = 9;
  double total_cost = 10;
  string currency = 11;
  google.protobuf.Timestamp purchase_date = 12;
  string supplier = 13;
  string supplier_contact = 14;
  string invoice_number = 15;
  string payment_method = 16;
  string payment_status = 17;
  google.protobuf.Timestamp payment_due_date = 18;
  double tax_amount = 19;
  double discount_amount = 20;
  int32 warranty_period = 21;
  string notes = 22;
  string receipt_image = 23;
  string created_by = 24;
  google.protobuf.Timestamp created_at = 25;
  google.protobuf.Timestamp updated_at = 26;
}

message CostTrackingFilter {
  string category = 1 [
    (buf.validate.field).string.in = "seed",
    (buf.validate.field).string.in = "fertilizer",
    (buf.validate.field).string.in = "pesticide",
    (buf.validate.field).string.in = "labor",
    (buf.validate.field).string.in = "utilities",
    (buf.validate.field).string.in = "equipment",
    (buf.validate.field).string.in = "packaging",
    (buf.validate.field).string.in = "transportation"
  ];
  string cost_type = 2 [
    (buf.validate.field).string.in = "fixed",
    (buf.validate.field).string.in = "variable",
    (buf.validate.field).string.in = "one_time",
    (buf.validate.field).string.in = "recurring"
  ];
  string payment_status = 3 [
    (buf.validate.field).string.in = "pending",
    (buf.validate.field).string.in = "paid",
    (buf.validate.field).string.in = "overdue",
    (buf.validate.field).string.in = "cancelled"
  ];
  string supplier = 4 [ (buf.validate.field).string = {max_len : 200} ];
  string search = 5 [ (buf.validate.field).string = {max_len : 200} ];
}

message CreateCostTrackingRequest {
  string planting_cycle_id = 1
      [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
  string cost_category = 2 [
    (buf.validate.field).string.in = "seed",
    (buf.validate.field).string.in = "fertilizer",
    (buf.validate.field).string.in = "pesticide",
    (buf.validate.field).string.in = "labor",
    (buf.validate.field).string.in = "utilities",
    (buf.validate.field).string.in = "equipment",
    (buf.validate.field).string.in = "packaging",
    (buf.validate.field).string.in = "transportation"
  ];
  string cost_type = 3 [
    (buf.validate.field).string.in = "fixed",
    (buf.validate.field).string.in = "variable",
    (buf.validate.field).string.in = "one_time",
    (buf.validate.field).string.in = "recurring"
  ];
  string item_name = 4
      [ (buf.validate.field).string = {min_len : 1, max_len : 200} ];
  string description = 5 [ (buf.validate.field).string = {max_len : 1000} ];
  double quantity = 6 [ (buf.validate.field).double = {gte : 0} ];
  string unit = 7 [ (buf.validate.field).string = {min_len : 1, max_len : 20} ];
  double unit_cost = 8 [ (buf.validate.field).double = {gte : 0} ];
  double total_cost = 9 [ (buf.validate.field).double = {gte : 0} ];
  string currency = 10
      [ (buf.validate.field).string =
            {min_len : 3, max_len : 3, pattern : "^[A-Z]{3}$"} ];
  google.protobuf.Timestamp purchase_date = 11;
  string supplier = 12 [ (buf.validate.field).string = {max_len : 200} ];
  string supplier_contact = 13
      [ (buf.validate.field).string = {max_len : 100} ];
  string invoice_number = 14 [ (buf.validate.field).string = {max_len : 100} ];
  string payment_method = 15 [
    (buf.validate.field).string.in = "cash",
    (buf.validate.field).string.in = "bank",
    (buf.validate.field).string.in = "card"
  ];
  string payment_status = 16 [
    (buf.validate.field).string.in = "pending",
    (buf.validate.field).string.in = "paid",
    (buf.validate.field).string.in = "overdue",
    (buf.validate.field).string.in = "cancelled"
  ];
  google.protobuf.Timestamp payment_due_date = 17;
  double tax_amount = 18 [ (buf.validate.field).double = {gte : 0} ];
  double discount_amount = 19 [ (buf.validate.field).double = {gte : 0} ];
  int32 warranty_period = 20
      [ (buf.validate.field).int32 = {gte : 0, lte : 3650} ];
  string notes = 21 [ (buf.validate.field).string = {max_len : 2000} ];
  string receipt_image = 22 [ (buf.validate.field).string = {max_len : 500} ];
  string created_by = 23
      [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
}

message GetCostTrackingRequest {
  string id = 1 [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
}

message UpdateCostTrackingRequest {
  string id = 1 [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
  string planting_cycle_id = 2
      [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
  string cost_category = 3 [
    (buf.validate.field).string.in = "seed",
    (buf.validate.field).string.in = "fertilizer",
    (buf.validate.field).string.in = "pesticide",
    (buf.validate.field).string.in = "labor",
    (buf.validate.field).string.in = "utilities",
    (buf.validate.field).string.in = "equipment",
    (buf.validate.field).string.in = "packaging",
    (buf.validate.field).string.in = "transportation"
  ];
  string cost_type = 4 [
    (buf.validate.field).string.in = "fixed",
    (buf.validate.field).string.in = "variable",
    (buf.validate.field).string.in = "one_time",
    (buf.validate.field).string.in = "recurring"
  ];
  string item_name = 5
      [ (buf.validate.field).string = {min_len : 1, max_len : 200} ];
  string description = 6 [ (buf.validate.field).string = {max_len : 1000} ];
  double quantity = 7 [ (buf.validate.field).double = {gte : 0} ];
  string unit = 8 [ (buf.validate.field).string = {min_len : 1, max_len : 20} ];
  double unit_cost = 9 [ (buf.validate.field).double = {gte : 0} ];
  double total_cost = 10 [ (buf.validate.field).double = {gte : 0} ];
  string currency = 11
      [ (buf.validate.field).string =
            {min_len : 3, max_len : 3, pattern : "^[A-Z]{3}$"} ];
  google.protobuf.Timestamp purchase_date = 12;
  string supplier = 13 [ (buf.validate.field).string = {max_len : 200} ];
  string supplier_contact = 14
      [ (buf.validate.field).string = {max_len : 100} ];
  string invoice_number = 15 [ (buf.validate.field).string = {max_len : 100} ];
  string payment_method = 16 [
    (buf.validate.field).string.in = "cash",
    (buf.validate.field).string.in = "bank",
    (buf.validate.field).string.in = "card"
  ];
  string payment_status = 17 [
    (buf.validate.field).string.in = "pending",
    (buf.validate.field).string.in = "paid",
    (buf.validate.field).string.in = "overdue",
    (buf.validate.field).string.in = "cancelled"
  ];
  google.protobuf.Timestamp payment_due_date = 18;
  double tax_amount = 19 [ (buf.validate.field).double = {gte : 0} ];
  double discount_amount = 20 [ (buf.validate.field).double = {gte : 0} ];
  int32 warranty_period = 21
      [ (buf.validate.field).int32 = {gte : 0, lte : 3650} ];
  string notes = 22 [ (buf.validate.field).string = {max_len : 2000} ];
  string receipt_image = 23 [ (buf.validate.field).string = {max_len : 500} ];
}

message DeleteCostTrackingRequest {
  string id = 1 [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
}

message DeleteCostTrackingResponse { string message = 1; }

message ListCostTrackingsRequest {
  common.PaginationRequest pagination = 1;
  CostTrackingFilter filter = 2;
}

message GetCostTrackingsByPlantingCycleRequest {
  string planting_cycle_id = 1
      [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
  common.PaginationRequest pagination = 2;
  CostTrackingFilter filter = 3;
}

message CreateCostTrackingResponse { CostTracking cost_tracking = 1; }

message GetCostTrackingResponse { CostTracking cost_tracking = 1; }

message UpdateCostTrackingResponse { CostTracking cost_tracking = 1; }

message ListCostTrackingsResponse {
  repeated CostTracking cost_trackings = 1;
  common.PaginationResponse pagination = 2;
}

message GetCostTrackingsByPlantingCycleResponse {
  repeated CostTracking cost_trackings = 1;
  common.PaginationResponse pagination = 2;
}
