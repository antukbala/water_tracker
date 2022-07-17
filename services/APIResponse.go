package service

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var Response = make(map[string]APIResponse)

func init() {

	Response["catch_error"] = APIResponse{4000, "Error from catch block"}
	Response["token_not_allowed"] = APIResponse{403, "Token not allowed"}
	Response["qr_token_invalid"] = APIResponse{4003, "Qr Token invalid"}
	Response["text_not_found"] = APIResponse{4000, "Text not found"}
	Response["image_not_valid"] = APIResponse{4000, "Image not valid.Please upload correct image"}
	Response["request_not_valid"] = APIResponse{4000, "Request not found.Please request first"}
	Response["field_missing"] = APIResponse{4001, "Field missing"}
	Response["success"] = APIResponse{1000, "success"}
	Response["SQL_Error"] = APIResponse{4001, "Syntax error"}

	Response["process_failled"] = APIResponse{4002, "Process failled"}
	Response["invalid_token"] = APIResponse{4002, "Invalid Token"}
}
func GetCode(message string) APIResponse {
	return Response[message]
}
