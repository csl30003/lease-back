package dto

//const productId = ref(0) // 商品id
//const productName = ref('') // 商品名称
//const productPrice = ref(0) // 商品价格
//const productImage = ref('') // 商品图片
//const productQuantity = ref(0) // 商品数量
//const useDays = ref(0) // 租赁天数
//const delivery = ref(0) // 配送方式
//const deliveryValue = ref('') // 配送方式
//const actualPayment = ref(0) // 实付款
//const paymentType = ref(0) // 支付方式id，默认0支付宝
//const paymentTypeValue = ref('支付宝支付') // 支付方式
//const userId = ref(0) // 用户id
//const hisId = ref(0) // 对方id
//
//const myAddressId = ref(0) // 我方地址id
//const myAddressName = ref('') // 我方名称
//const myAddressPhone = ref('') // 我方电话
//const myAddressProvince = ref('') // 我方省
//const myAddressCity = ref('') // 我方市
//const myAddressDistrict = ref('') // 我方区
//const myAddressDetail = ref('') // 我方详细地址
//const hisAddressId = ref(0) // 对方地址id
//const hisAddressName = ref('') // 对方名称
//const hisAddressPhone = ref('') // 对方电话
//const hisAddressProvince = ref('') // 对方省
//const hisAddressCity = ref('') // 对方市
//const hisAddressDistrict = ref('') // 对方区
//const hisAddressDetail = ref('') // 对方详细地址

type GetOrderResp struct {
	ID                  int     `json:"id"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
	Identifier          string  `json:"identifier"`
	Status              int     `json:"status"`
	PayTime             string  `json:"pay_time"`
	HisDeliveryTime     string  `json:"his_delivery_time"`
	MyReceiveTime       string  `json:"my_receive_time"`
	ReturnTime          string  `json:"return_time"`
	HisReceiveTime      string  `json:"his_receive_time"`
	InspectCompleteTime string  `json:"inspect_complete_time"`
	AllSolveTime        string  `json:"all_solve_time"`
	CompleteTime        string  `json:"complete_time"`
	ProductPrice        float64 `json:"product_price"`
	UseDays             int     `json:"use_days"`
	ProductQuantity     int     `json:"product_quantity"`
	Freight             float64 `json:"freight"`
	ActualPayment       float64 `json:"actual_payment"`
	PaymentType         int     `json:"payment_type"`
	UserID              int     `json:"user_id"`
	HisID               int     `json:"his_id"`
	MyAddressID         int     `json:"my_address_id"`
	MyAddressName       string  `json:"my_address_name"`
	MyAddressPhone      string  `json:"my_address_phone"`
	MyAddressProvince   string  `json:"my_address_province"`
	MyAddressCity       string  `json:"my_address_city"`
	MyAddressDistrict   string  `json:"my_address_district"`
	MyAddressDetail     string  `json:"my_address_detail"`
	HisAddressID        int     `json:"his_address_id"`
	HisAddressName      string  `json:"his_address_name"`
	HisAddressPhone     string  `json:"his_address_phone"`
	HisAddressProvince  string  `json:"his_address_province"`
	HisAddressCity      string  `json:"his_address_city"`
	HisAddressDistrict  string  `json:"his_address_district"`
	HisAddressDetail    string  `json:"his_address_detail"`
	ProductID           int     `json:"product_id"`
	ProductName         string  `json:"product_name"`
	ProductImage        string  `json:"product_image"`
}
