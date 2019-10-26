package dao

import "strings"

const captchaTemplate = `
<html>
    <body>
        <div>尊敬的用户，您好！</div>
        <h2>
		欢迎注册杂货铺社区
        </h2>
        </br>
        <hr>
        <div style="color: #99ccff;font-size: 26px;text-align: center;"><strong>你的验证码是：</strong></div>
        <pre style="overflow-x:scroll"">
{{CAPTCHA}}
        </pre>
    </body>
</html>
`

func getCaptchaTemplate(code string) string {
	return strings.ReplaceAll(captchaTemplate, "{{CAPTCHA}}", code)
}
