-- 14. BẢNG THEO DÕI CHI PHÍ
CREATE TABLE cost_tracking (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()), -- Mã định danh duy nhất cho mỗi chi phí (UUID)
    
    planting_cycle_id VARCHAR(36), -- Tham chiếu đến chu kỳ trồng trọt liên quan đến khoản chi phí
    
    cost_category VARCHAR(100), -- Loại chi phí chính: seed, fertilizer, pesticide, labor, utilities, equipment, packaging, transportation
    
    cost_type VARCHAR(100), -- Hình thức chi phí: fixed, variable, one_time, recurring
    
    item_name VARCHAR(255), -- Tên vật tư, sản phẩm hoặc dịch vụ
    description TEXT, -- Mô tả chi tiết về vật tư hoặc dịch vụ
    
    quantity DECIMAL(12,4), -- Số lượng mua hoặc sử dụng
    unit VARCHAR(20), -- Đơn vị đo lường (kg, lít, cái, giờ công...)
    
    unit_cost DECIMAL(12,2), -- Chi phí cho một đơn vị
    total_cost DECIMAL(12,2), -- Tổng chi phí = quantity * unit_cost - discount + tax
    
    currency VARCHAR(10) DEFAULT 'VND', -- Loại tiền tệ, mặc định là VND
    
    purchase_date DATE, -- Ngày phát sinh chi phí hoặc ngày mua
    
    supplier VARCHAR(255), -- Tên nhà cung cấp
    supplier_contact VARCHAR(255), -- Thông tin liên hệ của nhà cung cấp
    
    invoice_number VARCHAR(100), -- Số hóa đơn/chứng từ
    
    payment_method VARCHAR(50), -- Phương thức thanh toán: tiền mặt (casher), chuyển khoản (bank), thẻ...
    payment_status VARCHAR(50), -- Trạng thái thanh toán: pending, paid, overdue, cancelled
    payment_due_date DATE, -- Ngày đến hạn thanh toán (nếu có công nợ)
    
    tax_amount DECIMAL(12,2), -- Thuế áp dụng cho chi phí
    discount_amount DECIMAL(12,2), -- Số tiền chiết khấu/giảm giá
    
    warranty_period INTEGER, -- Thời gian bảo hành (tính bằng ngày)
    
    notes TEXT, -- Ghi chú bổ sung
    
    receipt_image VARCHAR(36), -- Tham chiếu đến bảng media lưu hình ảnh biên lai
    
    created_by VARCHAR(36), -- Người tạo bản ghi (tham chiếu đến bảng users)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời điểm tạo bản ghi
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Thời điểm cập nhật gần nhất
    
    INDEX idx_cost_tracking_cycle_category (planting_cycle_id, cost_category), -- Index tìm kiếm theo chu kỳ và loại chi phí
    INDEX idx_cost_tracking_date (purchase_date), -- Index tìm kiếm theo ngày phát sinh
    INDEX idx_cost_tracking_supplier (supplier(100)) -- Index tìm kiếm nhanh theo nhà cung cấp
);
