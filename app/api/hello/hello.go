package hello

import (
    "gf-demo/app/service"
    "gf-demo/app/service/define"
    "gf-demo/library/oss"
    "gf-demo/library/response"
    "gf-demo/library/wechat"
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/net/ghttp"
    "github.com/gogf/gf/os/gfile"
)

var Hello = new(helloApi)

type helloApi struct {}


// @summary 分页获取用户
// @tags    用户服务
// @produce json
// @Security ApiKeyAuth
// @Param data body define.UserQuery true "用户分页查询"
// @router  /hello/users [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *helloApi) Users(r *ghttp.Request)  {
    var (
        q *define.UserQuery
    )
    if err := r.Parse(&q); err != nil {
        response.JsonExit(r, -1, "参数错误")
    }
    if list, err := service.User.GetUserList(q); err != nil {
        response.JsonExit(r, -1, "查询异常")
    } else {
        response.SuccExit(r, "ok", list)
    }
}

// @summary     测试接口
// @tags        用户服务
// @produce     json
// @router      /hello/test [POST]
// @success     200 {object} response.JsonResponse "执行结果"
func (a *helloApi) Test(r *ghttp.Request)  {
    if result, err := wechat.MiniCode2Session("123123213"); err == nil {
        response.JsonExit(r, 0, result.OpenID)
    }
    response.JsonExit(r,0, "test")
}

// @summary     获取当前用户信息
// @tags        用户服务
// @produce     json
// @Security    ApiKeyAuth
// @router      /hello/profile [GET]
// @success     200 {object} define.ProfileResult "执行结果"
func (a *helloApi) Profile(r *ghttp.Request) {
    if profile, err := service.User.Profile(r.Context()); err != nil {
        response.Fail(r, -1, err.Error())
    } else {
        response.Succ(r, "ok", profile)
    }
}

// @summary     上传文件
// @tags        用户服务
// @produce     json
// @Security    ApiKeyAuth
// @Param       file formData file true "文件"
// @router      /hello/upload [POST]
// @success     200 {object} response.JsonResponse "执行结果"
func (a helloApi) Upload(r *ghttp.Request) {
    upfile := r.GetUploadFile("file")
    if fileName, err := upfile.Save(gfile.TempDir()); err != nil {
        response.FailExit(r, -1, "上传失败")
    } else {
        tempFile := gfile.Join(gfile.TempDir(), fileName)
        _, err := oss.FileUpload(&oss.MinioFileObject{
            ObjectName: upfile.Filename,
            FilePath: tempFile,
            Bucket: oss.MinioBucket{
                BucketName: "img",
                Location: "cn-north-1",
            },
            Config: oss.MinioConfig{
                EndPoint: g.Cfg().GetString("minio.endpoint"),
                AccessKeyID: g.Cfg().GetString("minio.accessKeyID"),
                SecretAccessKey: g.Cfg().GetString("minio.secretAccessKey"),
                UseSSL: g.Cfg().GetBool("minio.useSSL"),
            },
        })
        if err != nil {
            response.FailExit(r, -1, "上传失败")
        } else {
            go gfile.Remove(tempFile)
            response.Succ(r,"上传成功")
        }
    }
}

// @summary     查看文件
// @tags        用户服务
// @produce     json
// @Security    ApiKeyAuth
// @Param       fileName query string true "文件名"
// @router      /hello/viewfile [GET]
// @success     200 {object} response.JsonResponse "执行结果"
func (a *helloApi) ViewFile(r *ghttp.Request) {
    //data := r.GetFormMapStrStr()
    //fileName := data["fileName"]
    fileName := r.GetQueryString("fileName")
    miniFileObject := &oss.MinioFileObject{
        ObjectName: fileName,
        Bucket: oss.MinioBucket{
            BucketName: "img",
            Location: "cn-north-1",
        },
        Config: oss.MinioConfig{
            EndPoint: g.Cfg().GetString("minio.endpoint"),
            AccessKeyID: g.Cfg().GetString("minio.accessKeyID"),
            SecretAccessKey: g.Cfg().GetString("minio.secretAccessKey"),
            UseSSL: g.Cfg().GetBool("minio.useSSL"),
        },
    }
    object, err := oss.GetFile(miniFileObject)
    if err != nil {
        response.FailExit(r, -1, "文件不存在")
    } else {
        r.Response.Write(object)
    }
}