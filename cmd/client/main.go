package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_cost_tracking "github.com/anhvanhoa/sf-proto/gen/cost_tracking/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var serverAddress string

func init() {
	viper.SetConfigFile("dev.config.yml")
	viper.ReadInConfig()
	serverAddress = fmt.Sprintf("%s:%s", viper.GetString("host_grpc"), viper.GetString("port_grpc"))
}

func inputPaging(reader *bufio.Reader) (int32, int32) {
	fmt.Print("Nhập trang (mặc định 1): ")
	offsetStr, _ := reader.ReadString('\n')
	offsetStr = cleanInput(offsetStr)
	offset := int32(1)
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = int32(o)
		}
	}

	fmt.Print("Nhập số bản ghi mỗi trang (mặc định 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = cleanInput(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = int32(l)
		}
	}

	return offset, limit
}

type CostTrackingServiceClient struct {
	costTrackingClient proto_cost_tracking.CostTrackingServiceClient
	conn               *grpc.ClientConn
}

func NewCostTrackingServiceClient(address string) (*CostTrackingServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &CostTrackingServiceClient{
		costTrackingClient: proto_cost_tracking.NewCostTrackingServiceClient(conn),
		conn:               conn,
	}, nil
}

func (c *CostTrackingServiceClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// --- Helper để làm sạch input ---
func cleanInput(s string) string {
	return strings.ToValidUTF8(strings.TrimSpace(s), "")
}

// ================== Cost Tracking Service Tests ==================

func (c *CostTrackingServiceClient) TestCreateCostTracking() {
	fmt.Println("\n=== Kiểm thử Tạo bản ghi Chi phí ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID chu kỳ trồng: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	fmt.Print("Nhập danh mục chi phí (seed/fertilizer/pesticide/labor/utilities/equipment/packaging/transportation): ")
	costCategory, _ := reader.ReadString('\n')
	costCategory = cleanInput(costCategory)

	fmt.Print("Nhập loại chi phí (fixed/variable/one_time/recurring): ")
	costType, _ := reader.ReadString('\n')
	costType = cleanInput(costType)

	fmt.Print("Nhập tên sản phẩm: ")
	itemName, _ := reader.ReadString('\n')
	itemName = cleanInput(itemName)

	fmt.Print("Nhập mô tả: ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	fmt.Print("Nhập số lượng: ")
	quantityStr, _ := reader.ReadString('\n')
	quantityStr = cleanInput(quantityStr)
	quantity := float64(1.0)
	if quantityStr != "" {
		if q, err := strconv.ParseFloat(quantityStr, 64); err == nil {
			quantity = q
		}
	}

	fmt.Print("Nhập đơn vị: ")
	unit, _ := reader.ReadString('\n')
	unit = cleanInput(unit)

	fmt.Print("Nhập giá đơn vị: ")
	unitCostStr, _ := reader.ReadString('\n')
	unitCostStr = cleanInput(unitCostStr)
	unitCost := float64(0.0)
	if unitCostStr != "" {
		if c, err := strconv.ParseFloat(unitCostStr, 64); err == nil {
			unitCost = c
		}
	}

	fmt.Print("Nhập tổng chi phí: ")
	totalCostStr, _ := reader.ReadString('\n')
	totalCostStr = cleanInput(totalCostStr)
	totalCost := float64(0.0)
	if totalCostStr != "" {
		if c, err := strconv.ParseFloat(totalCostStr, 64); err == nil {
			totalCost = c
		}
	}

	fmt.Print("Nhập tiền tệ (VD: VND, USD): ")
	currency, _ := reader.ReadString('\n')
	currency = cleanInput(currency)

	fmt.Print("Nhập ngày mua (YYYY-MM-DD): ")
	purchaseDateStr, _ := reader.ReadString('\n')
	purchaseDateStr = cleanInput(purchaseDateStr)
	var purchaseDate *timestamppb.Timestamp
	if purchaseDateStr != "" {
		if t, err := time.Parse("2006-01-02", purchaseDateStr); err == nil {
			purchaseDate = timestamppb.New(t)
		}
	}

	fmt.Print("Nhập nhà cung cấp: ")
	supplier, _ := reader.ReadString('\n')
	supplier = cleanInput(supplier)

	fmt.Print("Nhập liên hệ nhà cung cấp: ")
	supplierContact, _ := reader.ReadString('\n')
	supplierContact = cleanInput(supplierContact)

	fmt.Print("Nhập số hóa đơn: ")
	invoiceNumber, _ := reader.ReadString('\n')
	invoiceNumber = cleanInput(invoiceNumber)

	fmt.Print("Nhập phương thức thanh toán (cash/bank/card): ")
	paymentMethod, _ := reader.ReadString('\n')
	paymentMethod = cleanInput(paymentMethod)

	fmt.Print("Nhập trạng thái thanh toán (pending/paid/overdue/cancelled): ")
	paymentStatus, _ := reader.ReadString('\n')
	paymentStatus = cleanInput(paymentStatus)

	fmt.Print("Nhập ngày đến hạn thanh toán (YYYY-MM-DD): ")
	paymentDueDateStr, _ := reader.ReadString('\n')
	paymentDueDateStr = cleanInput(paymentDueDateStr)
	var paymentDueDate *timestamppb.Timestamp
	if paymentDueDateStr != "" {
		if t, err := time.Parse("2006-01-02", paymentDueDateStr); err == nil {
			paymentDueDate = timestamppb.New(t)
		}
	}

	fmt.Print("Nhập số tiền thuế: ")
	taxAmountStr, _ := reader.ReadString('\n')
	taxAmountStr = cleanInput(taxAmountStr)
	taxAmount := float64(0.0)
	if taxAmountStr != "" {
		if t, err := strconv.ParseFloat(taxAmountStr, 64); err == nil {
			taxAmount = t
		}
	}

	fmt.Print("Nhập số tiền chiết khấu: ")
	discountAmountStr, _ := reader.ReadString('\n')
	discountAmountStr = cleanInput(discountAmountStr)
	discountAmount := float64(0.0)
	if discountAmountStr != "" {
		if d, err := strconv.ParseFloat(discountAmountStr, 64); err == nil {
			discountAmount = d
		}
	}

	fmt.Print("Nhập thời gian bảo hành (ngày): ")
	warrantyPeriodStr, _ := reader.ReadString('\n')
	warrantyPeriodStr = cleanInput(warrantyPeriodStr)
	warrantyPeriod := int32(0)
	if warrantyPeriodStr != "" {
		if w, err := strconv.Atoi(warrantyPeriodStr); err == nil {
			warrantyPeriod = int32(w)
		}
	}

	fmt.Print("Nhập ghi chú: ")
	notes, _ := reader.ReadString('\n')
	notes = cleanInput(notes)

	fmt.Print("Nhập hình ảnh biên lai: ")
	receiptImage, _ := reader.ReadString('\n')
	receiptImage = cleanInput(receiptImage)

	fmt.Print("Nhập người tạo: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.costTrackingClient.CreateCostTracking(ctx, &proto_cost_tracking.CreateCostTrackingRequest{
		PlantingCycleId: plantingCycleId,
		CostCategory:    costCategory,
		CostType:        costType,
		ItemName:        itemName,
		Description:     description,
		Quantity:        quantity,
		Unit:            unit,
		UnitCost:        unitCost,
		TotalCost:       totalCost,
		Currency:        currency,
		PurchaseDate:    purchaseDate,
		Supplier:        supplier,
		SupplierContact: supplierContact,
		InvoiceNumber:   invoiceNumber,
		PaymentMethod:   paymentMethod,
		PaymentStatus:   paymentStatus,
		PaymentDueDate:  paymentDueDate,
		TaxAmount:       taxAmount,
		DiscountAmount:  discountAmount,
		WarrantyPeriod:  warrantyPeriod,
		Notes:           notes,
		ReceiptImage:    receiptImage,
		CreatedBy:       createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateCostTracking: %v\n", err)
		return
	}

	fmt.Printf("Kết quả tạo bản ghi chi phí:\n")
	if resp.CostTracking != nil {
		fmt.Printf("ID: %s\n", resp.CostTracking.Id)
		fmt.Printf("Planting Cycle ID: %s\n", resp.CostTracking.PlantingCycleId)
		fmt.Printf("Cost Category: %s\n", resp.CostTracking.CostCategory)
		fmt.Printf("Cost Type: %s\n", resp.CostTracking.CostType)
		fmt.Printf("Item Name: %s\n", resp.CostTracking.ItemName)
		fmt.Printf("Total Cost: %.2f %s\n", resp.CostTracking.TotalCost, resp.CostTracking.Currency)
	}
}

func (c *CostTrackingServiceClient) TestGetCostTracking() {
	fmt.Println("\n=== Kiểm thử Lấy bản ghi Chi phí ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID bản ghi chi phí: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.costTrackingClient.GetCostTracking(ctx, &proto_cost_tracking.GetCostTrackingRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetCostTracking: %v\n", err)
		return
	}

	fmt.Printf("Kết quả lấy bản ghi chi phí:\n")
	if resp.CostTracking != nil {
		fmt.Printf("ID: %s\n", resp.CostTracking.Id)
		fmt.Printf("Planting Cycle ID: %s\n", resp.CostTracking.PlantingCycleId)
		fmt.Printf("Cost Category: %s\n", resp.CostTracking.CostCategory)
		fmt.Printf("Cost Type: %s\n", resp.CostTracking.CostType)
		fmt.Printf("Item Name: %s\n", resp.CostTracking.ItemName)
		fmt.Printf("Description: %s\n", resp.CostTracking.Description)
		fmt.Printf("Quantity: %.2f %s\n", resp.CostTracking.Quantity, resp.CostTracking.Unit)
		fmt.Printf("Unit Cost: %.2f %s\n", resp.CostTracking.UnitCost, resp.CostTracking.Currency)
		fmt.Printf("Total Cost: %.2f %s\n", resp.CostTracking.TotalCost, resp.CostTracking.Currency)
		fmt.Printf("Supplier: %s\n", resp.CostTracking.Supplier)
		fmt.Printf("Payment Status: %s\n", resp.CostTracking.PaymentStatus)
		fmt.Printf("Tax Amount: %.2f %s\n", resp.CostTracking.TaxAmount, resp.CostTracking.Currency)
		fmt.Printf("Discount Amount: %.2f %s\n", resp.CostTracking.DiscountAmount, resp.CostTracking.Currency)
		fmt.Printf("Warranty Period: %d days\n", resp.CostTracking.WarrantyPeriod)
		fmt.Printf("Notes: %s\n", resp.CostTracking.Notes)
		fmt.Printf("Created By: %s\n", resp.CostTracking.CreatedBy)
	}
}

func (c *CostTrackingServiceClient) TestListCostTrackings() {
	fmt.Println("\n=== Kiểm thử Liệt kê Bản ghi Chi phí ===")

	reader := bufio.NewReader(os.Stdin)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.costTrackingClient.ListCostTrackings(ctx, &proto_cost_tracking.ListCostTrackingsRequest{
		Pagination: &proto_common.PaginationRequest{
			Page:      offset,
			PageSize:  limit,
			SortBy:    "created_at",
			SortOrder: "desc",
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListCostTrackings: %v\n", err)
		return
	}

	fmt.Printf("Kết quả liệt kê bản ghi chi phí:\n")
	fmt.Printf("Tổng số: %d\n", resp.Pagination.Total)
	fmt.Printf("Danh sách bản ghi chi phí:\n")
	for i, record := range resp.CostTrackings {
		fmt.Printf("  [%d] ID: %s, Chu kỳ trồng: %s, Danh mục: %s, Tổng chi phí: %.2f %s\n",
			i+1, record.Id, record.PlantingCycleId, record.CostCategory, record.TotalCost, record.Currency)
	}
}

func (c *CostTrackingServiceClient) TestGetCostTrackingsByPlantingCycle() {
	fmt.Println("\n=== Kiểm thử Lấy Bản ghi Chi phí theo Chu kỳ trồng ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID chu kỳ trồng: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.costTrackingClient.GetCostTrackingsByPlantingCycle(ctx, &proto_cost_tracking.GetCostTrackingsByPlantingCycleRequest{
		PlantingCycleId: plantingCycleId,
		Pagination: &proto_common.PaginationRequest{
			Page:      offset,
			PageSize:  limit,
			SortBy:    "purchase_date",
			SortOrder: "desc",
		},
	})
	if err != nil {
		fmt.Printf("Error calling GetCostTrackingsByPlantingCycle: %v\n", err)
		return
	}

	fmt.Printf("Kết quả lấy bản ghi chi phí theo chu kỳ trồng:\n")
	fmt.Printf("Tổng số: %d\n", resp.Pagination.Total)
	fmt.Printf("Danh sách bản ghi chi phí:\n")
	for i, record := range resp.CostTrackings {
		fmt.Printf("  [%d] ID: %s, Danh mục: %s, Tổng chi phí: %.2f %s, Trạng thái: %s\n",
			i+1, record.Id, record.CostCategory, record.TotalCost, record.Currency, record.PaymentStatus)
	}
}

func (c *CostTrackingServiceClient) TestUpdateCostTracking() {
	fmt.Println("\n=== Kiểm thử Cập nhật Bản ghi Chi phí ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID bản ghi chi phí cần cập nhật: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Nhập ID chu kỳ trồng: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	fmt.Print("Nhập danh mục chi phí (seed/fertilizer/pesticide/labor/utilities/equipment/packaging/transportation): ")
	costCategory, _ := reader.ReadString('\n')
	costCategory = cleanInput(costCategory)

	fmt.Print("Nhập loại chi phí (fixed/variable/one_time/recurring): ")
	costType, _ := reader.ReadString('\n')
	costType = cleanInput(costType)

	fmt.Print("Nhập tên sản phẩm: ")
	itemName, _ := reader.ReadString('\n')
	itemName = cleanInput(itemName)

	fmt.Print("Nhập mô tả: ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	fmt.Print("Nhập số lượng: ")
	quantityStr, _ := reader.ReadString('\n')
	quantityStr = cleanInput(quantityStr)
	quantity := float64(1.0)
	if quantityStr != "" {
		if q, err := strconv.ParseFloat(quantityStr, 64); err == nil {
			quantity = q
		}
	}

	fmt.Print("Nhập đơn vị: ")
	unit, _ := reader.ReadString('\n')
	unit = cleanInput(unit)

	fmt.Print("Nhập giá đơn vị: ")
	unitCostStr, _ := reader.ReadString('\n')
	unitCostStr = cleanInput(unitCostStr)
	unitCost := float64(0.0)
	if unitCostStr != "" {
		if c, err := strconv.ParseFloat(unitCostStr, 64); err == nil {
			unitCost = c
		}
	}

	fmt.Print("Nhập tổng chi phí: ")
	totalCostStr, _ := reader.ReadString('\n')
	totalCostStr = cleanInput(totalCostStr)
	totalCost := float64(0.0)
	if totalCostStr != "" {
		if c, err := strconv.ParseFloat(totalCostStr, 64); err == nil {
			totalCost = c
		}
	}

	fmt.Print("Nhập tiền tệ (VD: VND, USD): ")
	currency, _ := reader.ReadString('\n')
	currency = cleanInput(currency)

	fmt.Print("Nhập ngày mua (YYYY-MM-DD): ")
	purchaseDateStr, _ := reader.ReadString('\n')
	purchaseDateStr = cleanInput(purchaseDateStr)
	var purchaseDate *timestamppb.Timestamp
	if purchaseDateStr != "" {
		if t, err := time.Parse("2006-01-02", purchaseDateStr); err == nil {
			purchaseDate = timestamppb.New(t)
		}
	}

	fmt.Print("Nhập nhà cung cấp: ")
	supplier, _ := reader.ReadString('\n')
	supplier = cleanInput(supplier)

	fmt.Print("Nhập liên hệ nhà cung cấp: ")
	supplierContact, _ := reader.ReadString('\n')
	supplierContact = cleanInput(supplierContact)

	fmt.Print("Nhập số hóa đơn: ")
	invoiceNumber, _ := reader.ReadString('\n')
	invoiceNumber = cleanInput(invoiceNumber)

	fmt.Print("Nhập phương thức thanh toán (cash/bank/card): ")
	paymentMethod, _ := reader.ReadString('\n')
	paymentMethod = cleanInput(paymentMethod)

	fmt.Print("Nhập trạng thái thanh toán (pending/paid/overdue/cancelled): ")
	paymentStatus, _ := reader.ReadString('\n')
	paymentStatus = cleanInput(paymentStatus)

	fmt.Print("Nhập ngày đến hạn thanh toán (YYYY-MM-DD): ")
	paymentDueDateStr, _ := reader.ReadString('\n')
	paymentDueDateStr = cleanInput(paymentDueDateStr)
	var paymentDueDate *timestamppb.Timestamp
	if paymentDueDateStr != "" {
		if t, err := time.Parse("2006-01-02", paymentDueDateStr); err == nil {
			paymentDueDate = timestamppb.New(t)
		}
	}

	fmt.Print("Nhập số tiền thuế: ")
	taxAmountStr, _ := reader.ReadString('\n')
	taxAmountStr = cleanInput(taxAmountStr)
	taxAmount := float64(0.0)
	if taxAmountStr != "" {
		if t, err := strconv.ParseFloat(taxAmountStr, 64); err == nil {
			taxAmount = t
		}
	}

	fmt.Print("Nhập số tiền chiết khấu: ")
	discountAmountStr, _ := reader.ReadString('\n')
	discountAmountStr = cleanInput(discountAmountStr)
	discountAmount := float64(0.0)
	if discountAmountStr != "" {
		if d, err := strconv.ParseFloat(discountAmountStr, 64); err == nil {
			discountAmount = d
		}
	}

	fmt.Print("Nhập thời gian bảo hành (ngày): ")
	warrantyPeriodStr, _ := reader.ReadString('\n')
	warrantyPeriodStr = cleanInput(warrantyPeriodStr)
	warrantyPeriod := int32(0)
	if warrantyPeriodStr != "" {
		if w, err := strconv.Atoi(warrantyPeriodStr); err == nil {
			warrantyPeriod = int32(w)
		}
	}

	fmt.Print("Nhập ghi chú: ")
	notes, _ := reader.ReadString('\n')
	notes = cleanInput(notes)

	fmt.Print("Nhập hình ảnh biên lai: ")
	receiptImage, _ := reader.ReadString('\n')
	receiptImage = cleanInput(receiptImage)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.costTrackingClient.UpdateCostTracking(ctx, &proto_cost_tracking.UpdateCostTrackingRequest{
		Id:              id,
		PlantingCycleId: plantingCycleId,
		CostCategory:    costCategory,
		CostType:        costType,
		ItemName:        itemName,
		Description:     description,
		Quantity:        quantity,
		Unit:            unit,
		UnitCost:        unitCost,
		TotalCost:       totalCost,
		Currency:        currency,
		PurchaseDate:    purchaseDate,
		Supplier:        supplier,
		SupplierContact: supplierContact,
		InvoiceNumber:   invoiceNumber,
		PaymentMethod:   paymentMethod,
		PaymentStatus:   paymentStatus,
		PaymentDueDate:  paymentDueDate,
		TaxAmount:       taxAmount,
		DiscountAmount:  discountAmount,
		WarrantyPeriod:  warrantyPeriod,
		Notes:           notes,
		ReceiptImage:    receiptImage,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateCostTracking: %v\n", err)
		return
	}

	fmt.Printf("Kết quả cập nhật bản ghi chi phí:\n")
	if resp.CostTracking != nil {
		fmt.Printf("ID: %s\n", resp.CostTracking.Id)
		fmt.Printf("Planting Cycle ID: %s\n", resp.CostTracking.PlantingCycleId)
		fmt.Printf("Cost Category: %s\n", resp.CostTracking.CostCategory)
		fmt.Printf("Cost Type: %s\n", resp.CostTracking.CostType)
		fmt.Printf("Item Name: %s\n", resp.CostTracking.ItemName)
		fmt.Printf("Total Cost: %.2f %s\n", resp.CostTracking.TotalCost, resp.CostTracking.Currency)
	}
}

func (c *CostTrackingServiceClient) TestDeleteCostTracking() {
	fmt.Println("\n=== Kiểm thử Xóa Bản ghi Chi phí ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID bản ghi chi phí cần xóa: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.costTrackingClient.DeleteCostTracking(ctx, &proto_cost_tracking.DeleteCostTrackingRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteCostTracking: %v\n", err)
		return
	}

	fmt.Printf("Kết quả xóa bản ghi chi phí: %s\n", resp.Message)
}

// ================== Menu Functions ==================

func printMainMenu() {
	fmt.Println("\n=== Ứng dụng kiểm thử gRPC Cost Tracking Service ===")
	fmt.Println("1. Dịch vụ Theo dõi Chi phí")
	fmt.Println("0. Thoát")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func printCostTrackingMenu() {
	fmt.Println("\n=== Dịch vụ Theo dõi Chi phí ===")
	fmt.Println("1. Tạo bản ghi chi phí")
	fmt.Println("2. Lấy bản ghi chi phí")
	fmt.Println("3. Liệt kê bản ghi chi phí")
	fmt.Println("4. Cập nhật bản ghi chi phí")
	fmt.Println("5. Xóa bản ghi chi phí")
	fmt.Println("6. Lấy bản ghi theo chu kỳ trồng")
	fmt.Println("0. Quay lại menu chính")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func main() {
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Đang kết nối tới máy chủ gRPC tại %s...\n", address)
	client, err := NewCostTrackingServiceClient(address)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer client.Close()

	fmt.Println("Kết nối thành công!")

	reader := bufio.NewReader(os.Stdin)

	for {
		printMainMenu()
		choice, _ := reader.ReadString('\n')
		choice = cleanInput(choice)

		switch choice {
		case "1":
			// Dịch vụ Theo dõi Chi phí
			for {
				printCostTrackingMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateCostTracking()
				case "2":
					client.TestGetCostTracking()
				case "3":
					client.TestListCostTrackings()
				case "4":
					client.TestUpdateCostTracking()
				case "5":
					client.TestDeleteCostTracking()
				case "6":
					client.TestGetCostTrackingsByPlantingCycle()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "0":
			fmt.Println("Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ. Vui lòng thử lại.")
		}
	}
}
