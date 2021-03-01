package green

type ClinetInfo struct {
	SdkVersion  string `json:"sdkVersion"`
	CfgVersion  string `json:"cfgVersion"`
	UserType    string `json:"userType"`
	UserId      string `json:"userId"`
	UserNick    string `json:"userNick"`
	Avatar      string `json:"avatar"`
	Imei        string `json:"imei"`
	Imsi        string `json:"imsi"`
	Umid        string `json:"umid"`
	Ip          string `json:"ip"`
	Os          string `json:"os"`
	Channel     string `json:"channel"`
	HostAppName string `json:"hostAppName"`
	HostPackage string `json:"hostPackage"`
	HostVersion string `json:"hostVersion"`
}

type Task struct {
	DataId string `json:"dataId"`
	Url    string `json:"url"`
}

type BizData struct {
	BizType string   `json:"bizType"`
	Scenes  []string `json:"scenes"`
	Tasks   []Task   `json:"tasks"`
}

type ScanResponse struct {
	Code      int32        `json:"code"`      // HTTP状态码
	Msg       string       `json:"msg"`       // 信息
	Data      []ScanResult `json:"data"`      // 检测结果数据
	requestId string       `json:"requestId"` // 请求id
}

type ScanResult struct {
	Code      int32      `json:"code"`      // 错误码，和HTTP状态码一致
	Msg       string     `json:"msg"`       // 请求信息的响应信息
	DataId    string     `json:"dataId"`    // 检测对象对应的数据ID
	TaskId    string     `json:"taskId"`    // 检测任务的ID
	Url       string     `json:"url"`       // 检测对象的URL
	StoredUrl string     `json:"storedUrl"` // OSS存储空间对应的http_url地址
	Extras    HitLibInfo `json:"extras"`    // 额外附加信息
	Results   []Result   `json:"results"`   // 返回结果。调用成功时（code=200），返回结果中包含一个或多个元素。每个元素是个结构体
}

// 额外附加信息
type HitLibInfo struct {
	Context string `json:"context"` // 文字命中的自定义文本内容
	LibCode string `json:"libCode"` // 文字命中的自定义文本内容对应的库code
	LibName string `json:"libName"` // 文字命中的自定义文本内容对应的库名称
}

// 扫描结果
type Result struct {
	Scene           string        `json:"scene"`           // 图片检测场景
	Label           string        `json:"label"`           // 检测结果的分类。不同检测场景的结果分类不同.
	Sublabel        string        `json:"sublabel"`        // 如果检测场景包含智能鉴黄（porn）和暴恐涉政（terrorism），则该字段可以返回检测结果的细分类标签
	Suggestion      string        `json:"suggestion"`      // 建议您执行的后续操作 pass：结果正常，无需进行其余操作; review：结果不确定，需要进行人工审核; block：结果违规，建议直接删除或者限制公开
	Rate            float32       `json:"rate"`            // 置信度分数，取值范围：0（表示置信度最低）~100（表示置信度最高）。 如果suggestion为pass，则置信度越高，表示内容正常的可能性越高；如果suggestion为review或block，则置信度越高，表示内容违规的可能性越高
	Frames          []Frame       `json:"frames"`          // 如果待检测图片因为过长被截断，该参数返回截断后的每一帧图像的临时访问地址, 具体结构描述请参见frame
	HintWordsInfo   HintWordsInfo `json:"hintWordsInfo"`   // 图片中含有广告或文字违规信息时，返回图片中广告文字命中的风险关键词信息。具体结构描述请参见hintWordsInfo
	QrcodeData      []string      `json:"qrcodeData"`      // 图片中含有二维码时，返回图片中所有二维码包含的文本信息
	ProgramCodeData string        `json:"programCodeData"` // 图片中含有小程序码时，返回小程序码的位置信息，具体结构描述请参见programCodeData
	LogoData        string        `json:"logoData"`        // 图片中含有logo时，返回识别出来的logo信息，具体结构描述请参见logoData。
	SfaceData       string        `json:"sfaceData"`       // 图片中包含暴恐识涉政内容时，返回识别出来的暴恐涉政信息，具体结构描述请参见sfaceData
	OcrData         string        `json:"ocrData"`         // 识别到的图片中的完整文字信息
}

type Frame struct {
	Rate float32 `json:"rate"` // 置信度分数，取值范围：0~100，置信度越高表示检测结果的可信度越高。建议您不要在业务中使用该分数。
	Url  string  `json:"url"`  // 被截断的图片的临时访问URL，地址有效期是5分钟
}

type ProgramCodeData struct {
	X float32 `json:"x"` // 以图片左上角为坐标原点，小程序码区域左上角到y轴距离，单位：像素
	Y float32 `json:"y"` // 以图片左上角为坐标原点，小程序码区域左上角到x轴距离，单位：像素。
	W float32 `json:"w"` // 小程序码区域宽度，单位：像素。
	H float32 `json:"h"` // 小程序码区域高度，单位：像素。
}

type LogoData struct {
	Type string  `json:"type"` // 识别出的logo类型，取值为TV （台标）
	Name string  `json:"name"` //识别出的logo名称。
	X    float32 `json:"x"`    // 以图片左上角为坐标原点，logo区域左上角到y轴距离，单位：像素。
	Y    float32 `json:"y"`    // 以图片左上角为坐标原点，logo区域左上角到x轴距离，单位：像素。
	W    float32 `json:"w"`    // logo区域宽度，单位：像素。
	H    float32 `json:"h"`    // logo区域高度，单位：像素。
}

type SfaceData struct {
	X     float32 `json:"x"`     // 以图片左上角为坐标原点，人脸区域左上角到y轴距离，单位：像素。
	Y     float32 `json:"y"`     // 以图片左上角为坐标原点，人脸区域左上角到x轴距离，单位：像素。
	W     float32 `json:"w"`     // 人脸区域宽度，单位：像素。
	H     float32 `json:"h"`     // 人脸区域高度，单位：像素。
	Faces []Face  `json:"faces"` // 识别出的人脸信息，
}

type HintWordsInfo struct {
	Context string `json:"context"` // 文字命中的风险关键词内容。
}

type Face struct {
	Id   string  `json:"id"`   // 人脸id
	Name string  `json:"name"` // 字符串类型，相似人物的名称。
	Rate float32 `json:"rate"` // 浮点数类型，置信度分数，取值范围：0（表示置信度最低）~100（表示置信度最高）。置信度越高表示人物识别结果的可信度越高。
}
