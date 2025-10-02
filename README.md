# Cost Service

Microservice quản lý chi phí trồng trọt và theo dõi tài chính trong hệ thống nông nghiệp, được xây dựng bằng Go và tuân theo nguyên tắc Clean Architecture.

## 🏗️ Kiến trúc

Dự án này tuân theo **Clean Architecture** với sự phân tách rõ ràng các mối quan tâm:

```
├── domain/           # Tầng logic nghiệp vụ
│   ├── entity/       # Các thực thể nghiệp vụ cốt lõi
│   │   └── cost_tracking.go      # Entity theo dõi chi phí
│   ├── repository/   # Giao diện truy cập dữ liệu
│   │   └── cost_tracking_repository.go
│   └── usecase/      # Các trường hợp sử dụng nghiệp vụ
│       └── cost_tracking/       # Use cases theo dõi chi phí
├── infrastructure/   # Các mối quan tâm bên ngoài
│   ├── grpc_service/ # Triển khai API gRPC
│   │   └── cost_tracking/       # gRPC handlers theo dõi chi phí
│   └── repo/         # Triển khai repository cơ sở dữ liệu
├── bootstrap/        # Khởi tạo ứng dụng
└── cmd/             # Điểm vào ứng dụng
```

## 🚀 Tính năng

### Quản lý Chi phí Trồng trọt
- ✅ Tạo, đọc, cập nhật, xóa bản ghi chi phí
- ✅ Liệt kê chi phí với bộ lọc (loại chi phí, nhà cung cấp, trạng thái thanh toán)
- ✅ Theo dõi chi phí theo chu kỳ trồng trọt
- ✅ Quản lý nhiều loại chi phí (hạt giống, phân bón, thuốc trừ sâu, lao động, tiện ích, thiết bị, đóng gói, vận chuyển)
- ✅ Hỗ trợ phân loại chi phí (cố định, biến đổi, một lần, định kỳ)
- ✅ Quản lý thông tin nhà cung cấp và hóa đơn
- ✅ Theo dõi trạng thái thanh toán và ngày đến hạn
- ✅ Tính toán thuế và chiết khấu
- ✅ Lưu trữ hình ảnh biên lai
- ✅ Báo cáo chi phí theo chu kỳ trồng
- ✅ Hỗ trợ phân trang và sắp xếp
- ✅ Xác thực dữ liệu đầu vào

## 🛠️ Công nghệ sử dụng

- **Ngôn ngữ**: Go 1.24.6
- **Cơ sở dữ liệu**: PostgreSQL
- **API**: gRPC
- **Kiến trúc**: Clean Architecture
- **Dependencies**:
  - `github.com/go-pg/pg/v10` - PostgreSQL ORM
  - `google.golang.org/grpc` - Framework gRPC
  - `github.com/spf13/viper` - Quản lý cấu hình
  - `go.uber.org/zap` - Logging có cấu trúc

## 📋 Yêu cầu hệ thống

- Go 1.24.6 trở lên
- PostgreSQL 12 trở lên
- [golang-migrate](https://github.com/golang-migrate/migrate) để quản lý migration cơ sở dữ liệu

## 🚀 Hướng dẫn nhanh

### 1. Clone repository
```bash
git clone <repository-url>
cd cost-service
```

### 2. Cài đặt dependencies
```bash
go mod download
```

### 3. Thiết lập cơ sở dữ liệu
```bash
# Tạo cơ sở dữ liệu
make create-db

# Chạy migrations
make up
```

### 4. Cấu hình ứng dụng
Sao chép và chỉnh sửa file cấu hình:
```bash
cp dev.config.yml config.yml
```

Cập nhật chuỗi kết nối cơ sở dữ liệu trong `config.yml`:
```yaml
node_env: "development"
url_db: "postgres://postgres:123456@localhost:5432/cost_service_db?sslmode=disable"
name_service: "CostService"
port_grpc: 50054
host_grpc: "localhost"
interval_check: "20s"
timeout_check: "15s"
```

### 5. Chạy ứng dụng
```bash
# Build và chạy service chính
make run

# Hoặc chạy client để test
make client
```

## 🗄️ Quản lý Cơ sở dữ liệu

Dự án sử dụng `golang-migrate` để quản lý schema cơ sở dữ liệu:

```bash
# Chạy tất cả migrations đang chờ
make up

# Rollback migration cuối cùng
make down

# Reset cơ sở dữ liệu hoàn toàn
make reset

# Tạo migration mới
make create name=migration_name

# Force migration đến phiên bản cụ thể
make force version=1
```

## 🌱 Dữ liệu mẫu

Dự án bao gồm dữ liệu mẫu để phát triển và kiểm thử:

```bash
# Chèn dữ liệu mẫu vào cơ sở dữ liệu
make seed

# Reset cơ sở dữ liệu và chèn dữ liệu mẫu
make seed-reset

# Chèn dữ liệu mẫu vào cơ sở dữ liệu Docker
make docker-seed
```

### Dữ liệu mẫu bao gồm:

**Bản ghi chi phí mẫu với các loại chi phí đa dạng:**
- **Hạt giống**: Chi phí mua hạt giống các loại rau
- **Phân bón**: Chi phí phân bón hữu cơ và vô cơ
- **Thuốc trừ sâu**: Chi phí thuốc bảo vệ thực vật
- **Lao động**: Chi phí thuê nhân công
- **Tiện ích**: Chi phí điện, nước, gas
- **Thiết bị**: Chi phí mua sắm dụng cụ, máy móc
- **Đóng gói**: Chi phí bao bì, thùng carton
- **Vận chuyển**: Chi phí vận chuyển sản phẩm

Mỗi bản ghi chi phí bao gồm:
- Thông tin cơ bản (tên vật tư, mô tả, số lượng, đơn vị)
- Chi phí (đơn giá, tổng chi phí, thuế, chiết khấu)
- Thông tin nhà cung cấp và hóa đơn
- Trạng thái thanh toán và ngày đến hạn
- Liên kết với chu kỳ trồng trọt

## 📁 Cấu trúc Dự án

```
cost-service/
├── bootstrap/                 # Khởi tạo ứng dụng
│   ├── app.go               # Khởi tạo app
│   └── env.go               # Cấu hình môi trường
├── cmd/                     # Điểm vào ứng dụng
│   ├── main.go             # Điểm vào service chính
│   └── client/             # gRPC client để test
├── domain/                  # Logic nghiệp vụ (Clean Architecture)
│   ├── entity/             # Các thực thể nghiệp vụ cốt lõi
│   │   └── cost_tracking.go     # Entity theo dõi chi phí và DTOs
│   ├── repository/         # Giao diện truy cập dữ liệu
│   │   └── cost_tracking_repository.go
│   └── usecase/            # Các trường hợp sử dụng nghiệp vụ
│       └── cost_tracking/       # Use cases theo dõi chi phí
│           ├── create_cost_tracking_usecase.go
│           ├── get_cost_tracking_usecase.go
│           ├── update_cost_tracking_usecase.go
│           ├── delete_cost_tracking_usecase.go
│           ├── list_cost_tracking_usecase.go
│           ├── get_cost_trackings_by_planting_cycle_usecase.go
│           └── cost_tracking_usecase.go
├── infrastructure/          # Các mối quan tâm bên ngoài
│   ├── grpc_service/       # Triển khai API gRPC
│   │   ├── cost_tracking/        # gRPC handlers theo dõi chi phí
│   │   │   ├── base.go
│   │   │   ├── create.go
│   │   │   ├── get.go
│   │   │   ├── update.go
│   │   │   ├── delete.go
│   │   │   ├── list.go
│   │   │   └── get_by_planting_cycle.go
│   │   └── server.go             # Thiết lập gRPC server
│   └── repo/               # Triển khai cơ sở dữ liệu
│       ├── cost_tracking_repository.go
│       └── repository_factory.go
├── migrations/              # Database migrations
│   ├── 000000_common.up.sql
│   ├── 000001_create_trigger_updated_column.up.sql
│   ├── 000002_create_cost_tracking.up.sql
│   └── seed/                     # Dữ liệu mẫu
├── script/seed/             # Script chèn dữ liệu mẫu
├── doc/                     # Tài liệu
└── logs/                    # Log ứng dụng
```

## 🔧 Các lệnh có sẵn

```bash
# Thao tác cơ sở dữ liệu
make up              # Chạy migrations
make down            # Rollback migration
make reset           # Reset cơ sở dữ liệu
make create-db       # Tạo cơ sở dữ liệu
make drop-db         # Xóa cơ sở dữ liệu

# Ứng dụng
make build           # Build ứng dụng
make run             # Chạy service chính
make client          # Chạy client test
make test            # Chạy tests

# Trợ giúp
make help            # Hiển thị tất cả lệnh có sẵn
```

## 📊 Mô hình Dữ liệu

### Theo dõi Chi phí (Cost Tracking)
- **ID**: Định danh duy nhất (UUID)
- **PlantingCycleID**: ID chu kỳ trồng trọt liên quan
- **CostCategory**: Loại chi phí (seed, fertilizer, pesticide, labor, utilities, equipment, packaging, transportation)
- **CostType**: Hình thức chi phí (fixed, variable, one_time, recurring)
- **ItemName**: Tên vật tư/sản phẩm/dịch vụ
- **Description**: Mô tả chi tiết
- **Quantity**: Số lượng mua/sử dụng
- **Unit**: Đơn vị đo lường (kg, lít, cái, giờ công...)
- **UnitCost**: Chi phí cho một đơn vị
- **TotalCost**: Tổng chi phí (quantity × unit_cost - discount + tax)
- **Currency**: Loại tiền tệ (mặc định VND)
- **PurchaseDate**: Ngày phát sinh chi phí
- **Supplier**: Tên nhà cung cấp
- **SupplierContact**: Thông tin liên hệ nhà cung cấp
- **InvoiceNumber**: Số hóa đơn/chứng từ
- **PaymentMethod**: Phương thức thanh toán (cash, bank, card)
- **PaymentStatus**: Trạng thái thanh toán (pending, paid, overdue, cancelled)
- **PaymentDueDate**: Ngày đến hạn thanh toán
- **TaxAmount**: Số tiền thuế
- **DiscountAmount**: Số tiền chiết khấu/giảm giá
- **WarrantyPeriod**: Thời gian bảo hành (ngày)
- **Notes**: Ghi chú bổ sung
- **ReceiptImage**: Tham chiếu hình ảnh biên lai
- **CreatedBy**: Định danh người tạo
- **Timestamps**: Thời gian tạo/cập nhật

## 🔌 API Endpoints

Service cung cấp các endpoint gRPC:

### Cost Tracking Service
- `CreateCostTracking` - Tạo bản ghi chi phí mới
- `GetCostTracking` - Lấy thông tin chi phí theo ID
- `UpdateCostTracking` - Cập nhật thông tin chi phí
- `DeleteCostTracking` - Xóa bản ghi chi phí
- `ListCostTrackings` - Liệt kê chi phí với bộ lọc và phân trang
- `GetCostTrackingsByPlantingCycle` - Lấy chi phí theo chu kỳ trồng trọt

## 🧪 Testing

Chạy client test để tương tác với service:

```bash
make client
```

Điều này sẽ khởi động một client tương tác nơi bạn có thể test tất cả các endpoint gRPC.

## 📝 Cấu hình

Ứng dụng sử dụng Viper để quản lý cấu hình. Các tùy chọn cấu hình chính:

- `node_env`: Môi trường (development, production)
- `url_db`: Chuỗi kết nối PostgreSQL
- `name_service`: Tên service cho discovery
- `port_grpc`: Cổng gRPC server
- `host_grpc`: Host gRPC server
- `interval_check`: Khoảng thời gian kiểm tra sức khỏe
- `timeout_check`: Timeout kiểm tra sức khỏe

## 🚀 Triển khai

1. **Build ứng dụng**:
   ```bash
   make build
   ```

2. **Thiết lập cơ sở dữ liệu production**:
   ```bash
   make create-db
   make up
   ```

3. **Chạy service**:
   ```bash
   ./bin/app
   ```

## 🤝 Đóng góp

1. Fork repository
2. Tạo feature branch
3. Thực hiện thay đổi
4. Thêm tests nếu cần thiết
5. Submit pull request

## 📄 Giấy phép

Dự án này được cấp phép theo MIT License.

## 🆘 Hỗ trợ

Để được hỗ trợ và đặt câu hỏi, vui lòng tạo issue trong repository.

---

**Lưu ý**: Service này được thiết kế để quản lý chi phí trồng trọt và theo dõi tài chính trong hệ thống nông nghiệp, tuân theo các nguyên tắc kiến trúc microservice để có thể mở rộng và bảo trì dễ dàng.
