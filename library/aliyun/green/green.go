package green

import (
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var AliGreen = new(aliGreen)

type aliGreen struct{}

// 图片检测场景
type Scene int8

const (
	Porn      = iota // 图片智能鉴黄
	Terrorism        // 图片暴恐涉政
	Ad               // 图文违规
	Qrcode           // 图片二维码
	Live             // 图片不良场景
	Logo             // 图片logo
)

// 图片检测结果分类，参考 https://help.aliyun.com/document_detail/70292.html?spm=a2c4g.11186623.6.628.599d4cac2IY3ba
const (
	RESULTLABEL_NORMAL = "normal" // 检测结果分类正常值
	RESULTLABEL_PORN   = "porn"   // 检测结果分类色情
	RESULTLABEL_SFACE  = "sface"  // 检测结果敏感人物
)

// 扫描图片内容安全
func (sc *aliGreen) ScanImg(tasks []Task, bizType string, scenes []string) (*ScanResponse, error) {
	if bizType == "" {
		return nil, gerror.New("参数bizType不能为空")
	}
	if scenes == nil || len(scenes) == 0 {
		return nil, gerror.New("参数scenes不能为空")
	}
	var client IAliYunClient = DefaultClient
	bizData := BizData{
		BizType: bizType, //"default",
		Scenes:  scenes,  // []string{"porn"},
		Tasks:   tasks,
	}
	scanResult := client.GetScanResult(bizData)
	return scanResult, nil
}

func (sc *aliGreen) AnalysisScanResult(scanResult ScanResponse) []string {
	var ret []string
	if scanResult.Code >= 300 {
		ret = append(ret, "您上传的图片有误，请更换后重试")
	} else {
		arr := garray.NewFrom(gconv.Interfaces(scanResult.Data), true)
		arr.Iterator(func(k int, v interface{}) bool {
			result := new(ScanResult)
			if err := gconv.Struct(v, &result); err != nil {
				return false
			} else {
				if result.Code != 200 {
					// 检测结果
					ret = append(ret, "有不良图片，请更换")
				} else {
					ret = analysisResult(result.Results)
				}
				return true
			}
		})
	}
	return ret
}

// 分析task对应的结果
func analysisResult(results []Result) []string {
	var ret []string
	resultsArr := garray.NewFrom(gconv.Interfaces(results), true)
	resultsArr.Iterator(func(rk int, rv interface{}) bool {
		rs := new(Result)
		if err := gconv.Struct(rv, rs); err != nil {
			return false
		} else {
			switch rs.Scene {
			case "porn":
				{
					// 图片涉黄
					if rs.Label != RESULTLABEL_PORN {
						ret = append(ret, "图片涉黄")
					}
					break
				}
			case "terrorism":
				{
					// 图片涉恐
					teer := garray.NewStrArrayFrom(g.SliceStr{
						"bloody",
						"explosion",
						"outfit",
						"logo",
						"weapon",
						"politics",
						"violence",
						"crowd",
						"parade",
						"carcrash",
						"flag",
					}).Contains(rs.Label)
					if teer {
						ret = append(ret, "图片涉恐")
					}
					break
				}
			case "sface":
				{
					// 敏感人物
					if rs.Label == RESULTLABEL_SFACE {
						ret = append(ret, "图片含敏感人物")
					}
					break
				}
			}
			return true
		}
	})
	return ret
}

//// 敏感人脸识别
//func (sc *aliGreen) ScanFace(tasks []Task) (*ScanResponse) {
//	var client IAliYunClient = DefaultClient
//	bizData := BizData{
//		BizType: "default",
//		Scenes:  []string{"sface"},
//		Tasks: tasks,
//	}
//	scanResult := client.GetScanResult(bizData)
//	return scanResult
//}
